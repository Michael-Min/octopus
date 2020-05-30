package client

import (
	"context"
	"errors"
	"Michael-Min/octopus/discovery"
	pb "Michael-Min/octopus/proto"
	"Michael-Min/octopus/rpcsupport"
	"log"
	"time"

	"Michael-Min/octopus/engine"
	"Michael-Min/octopus/worker"
)

func CreateProcessor(
	clientChan chan pb.ReptilesClient) engine.Processor {
	return func(
		req engine.Request) (
		engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)
		c := <-clientChan
		// Call RPC to send work
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		sResult, err := c.Process(ctx, &pb.ProcessRequest{
			Url: sReq.Url, SerializedParser: sReq.SerializedParser})
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(*sResult),
			nil
	}
}

func ActClientPool(d *discovery.CrawlerDiscovery) chan pb.ReptilesClient {
	out := make(chan pb.ReptilesClient)
	go func() {
		clients := map[string]pb.ReptilesClient{}

		for {
			if host, ok := d.GetRandomIterm("worker"); ok {
				if client, ok := clients[host]; ok {
					log.Printf("[Rpc,Connect]: Old clinet to %s\n", host)
					out <- client
					continue
				} else {
					client, err := rpcsupport.NewClient(host)
					if err == nil {
						log.Printf("[Rpc,Connect]: new client to %s\n", host)
						clients[host] = client
						out <- client
						continue
					} else {
						log.Printf("[Error]: connecting to %s: %v\n", host, err)
					}
				}
			} else {
				log.Printf("[Warnning]: host is empty!\n")
			}
			time.Sleep(5 * time.Second)
		}
	}()
	return out
}

func CreateClientPool(
	hosts []string) (chan pb.ReptilesClient, error) {
	var clients []pb.ReptilesClient
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf(
				"Error connecting to %s: %v",
				h, err)
		}
	}

	if len(clients) == 0 {
		return nil, errors.New(
			"no connections available")
	}
	out := make(chan pb.ReptilesClient)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out, nil
}
