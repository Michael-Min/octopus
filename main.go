package main

import (
	"Michael-Min/octopus/config"
	"Michael-Min/octopus/discovery"
	"Michael-Min/octopus/engine"
	"Michael-Min/octopus/fetcher"
	"Michael-Min/octopus/gredis"
	persistClient "Michael-Min/octopus/persist/client"
	pb "Michael-Min/octopus/proto"
	"Michael-Min/octopus/scheduler"
	worker "Michael-Min/octopus/worker/client"
	"Michael-Min/octopus/xcar/parser"
)

func main() {
	var (
		itemChan chan pb.Item
		err      error
	)
	gredis.Setup()
	fetcher.SetVerboseLogging()
	itemChan, err = persistClient.ItemSaver(":9010")
	if err != nil {
		panic(err)
	}

	dis := discovery.NewCrawlerDiscover()
	dis.Discovery.Watch(dis.Hosts, "worker")
	pool := worker.ActClientPool(dis)

	processor := worker.CreateProcessor(pool)
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      2,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: "http://newcar.xcar.com.cn/car/",
		Parser: engine.NewFuncParser(
			parser.ParseCarList,
			config.ParseCarList),
	})
}
