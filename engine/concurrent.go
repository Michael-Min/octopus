package engine

import (
	"Michael-Min/octopus/bloom"
	"Michael-Min/octopus/config"
	"Michael-Min/octopus/gredis"
	pb "Michael-Min/octopus/proto"
	"Michael-Min/octopus/rabbitmq"
	"encoding/json"
	"log"
	"time"
)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan pb.Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

var (
	rateLimiter = time.Tick(
		10 * time.Second / config.Qps)
)

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult,1)
	gredis.Setup()
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(),
			out, e.Scheduler)
	}

	for _, r := range seeds {
		// 去重
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	mqSimple := rabbitmq.NewRabbitMQSimple("xcar")
	for {
		log.Println("Loop: fetch result <==>")
		result := <-out
		for _, item := range result.Items {
			go func(i pb.Item) {
				//use MQ
				bytes, _ := json.Marshal(i)
				if mqSimple.Conn.IsClosed() {
					log.Println("RabbitMQ连接断开，重新连接...")
					mqSimple.DoConnect()
				}
				err:=mqSimple.PublishSimple(string(bytes))
				if err != nil{
					log.Println(err)
				}
				//e.ItemChan <- i
				log.Printf("Iterm:%#v\n", i)
			}(*item)
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(
	in chan Request,
	out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			log.Println("Process wait request...")
			request := <-in
			log.Printf("Process load request...")
			<-rateLimiter
			log.Printf("Process fetch request start...")
			result, err := e.RequestProcessor(
				request)
			if err != nil {
				log.Printf("Error: Process return==> %v\n", err)
				continue
			}
			log.Printf("Process finish...")
			// 这里是获取的结果
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

//func isDuplicate(url string) bool {
//	if visitedUrls[url] {
//		return true
//	}
//
//	visitedUrls[url] = true
//	return false
//}

func isDuplicate(url string) bool {
	b, err := bloom.NewBloomFilter().IsContains(url)
	if err != nil {
		log.Println("IsContains failed:", err.Error())
		return false
	}
	if b == 1 {
		log.Printf("Existed url: %s\n", url)
		//debug时注释掉
		//return true
	}
	err = bloom.NewBloomFilter().Insert(url)
	if err != nil {
		log.Println("Insert failed:%s", err.Error())
		return false
	}
	return false
}
