package main

import (
	"Michael-Min/octopus/discovery"
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	dis := discovery.NewCrawlerDiscover()
	_ = dis.Discovery.Register(dis.Hosts, "worker", "http://127.0.0.1:9020", discovery.EtcdServiceInfo{Info: "1"})

	testwatch()
	time.Sleep(4 * time.Second)
	dis.Discovery.UpdateInfo("worker", "http://127.0.0.1:9020", discovery.EtcdServiceInfo{Info: "2"})
	time.Sleep(4 * time.Second)
	_ = dis.Discovery.Register(dis.Hosts, "worker", "http://127.0.0.1:9021", discovery.EtcdServiceInfo{Info: "1"})
	time.Sleep(4 * time.Second)
	dis.Discovery.Stop("worker", "http://127.0.0.1:9020")
	time.Sleep(4 * time.Second)
	dis.Discovery.Stop("worker", "http://127.0.0.1:9021")
	time.Sleep(4 * time.Second)
}

func testwatch() {
	dis := discovery.NewCrawlerDiscover()
	dis.Discovery.Watch(dis.Hosts, "worker")
	go func() {
		for {
			fmt.Println("[Test]: fetch all node....")
			if nodes, ok := dis.Discovery.GetServiceInfoAllNode("worker"); ok {
				bytes, _ := json.Marshal(nodes)
				if len(nodes) > 0 {
					fmt.Println(string(bytes))
					for _, node := range nodes {
						service, key, _ := discovery.SplitServiceNameKey(node.Key)
						fmt.Printf("[Test,Node]: name:%s, key:%s\n", service, key)
					}
				}
			}
			time.Sleep(2 * time.Second)
		}
	}()
}
