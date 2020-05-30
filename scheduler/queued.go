package scheduler

import (
	"fmt"
	"Michael-Min/octopus/engine"
	"log"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	fmt.Printf("Submit: request,url:%v\n",r.Url)
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(
	w chan engine.Request) {
	fmt.Println("Submit: chan request")
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 &&
				len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			log.Printf("Scheduler: requestQ size:%d, workerQ size:%d\n",len(requestQ),len(workerQ))
			select {
			case r := <-s.requestChan:
				log.Printf("Scheduler: get request")
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				log.Printf("Scheduler: get chan request")
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				log.Printf("Scheduler: filling chan request")
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
