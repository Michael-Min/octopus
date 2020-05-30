package main

import (
	"flag"
	"fmt"
	"Michael-Min/octopus/discovery"
	"Michael-Min/octopus/rpcsupport"
	"log"
	"net/http"

	"Michael-Min/octopus/fetcher"
)

var (
	port = flag.Int("port", 9000,
		"the port for me to listen on")
	httpPort = flag.Int("httpPort", 0,
		"the port for me to listen on")
)

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	dis := discovery.NewCrawlerDiscover()
	go func() {
		if err := dis.Discovery.Register(dis.Hosts, "worker", fmt.Sprintf("127.0.0.1:%d", *port), discovery.EtcdServiceInfo{Info: "1"}); err != nil {
			log.Fatal(err)
		}
		if err := rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), &rpcsupport.RPCService{}); err != nil {
			dis.Discovery.Stop("worker", fmt.Sprintf("127.0.0.1:%d", *port))
			log.Fatal(err)
		}
	}()

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		_, err := res.Write([]byte("pong"))
		if err != nil {
			log.Fatal("write err--->", err)
		}
	})
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil); err != nil {
		log.Fatal("open http err--->", err)
	}
}
