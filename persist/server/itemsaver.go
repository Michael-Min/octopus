package main

import (
	"Michael-Min/octopus/config"
	"Michael-Min/octopus/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	go func() {
		log.Fatal(serveRpc(
			fmt.Sprintf(":%d", *port),
			config.ElasticIndex))
	}()

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		_, err := res.Write([]byte("pong"))
		if err != nil {
			log.Fatal("write err--->", err)
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("open http err--->", err)
	}
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticHost),
		elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("[Error]: dail to ES wrong: %+v",err)
		return err
	}

	return rpcsupport.ServeRpc(host,
		&rpcsupport.RPCService{
			Client: client,
			Index:  index,
		})

}
