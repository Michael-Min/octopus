package client

import (
	pb "Michael-Min/octopus/proto"
	"Michael-Min/octopus/rabbitmq"
	"Michael-Min/octopus/rpcsupport"
	"context"
	"encoding/json"
	"log"
	"time"
)

func ItemSaver(
	host string) (chan pb.Item, error) {
	c, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan pb.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %s", itemCount, item.Car)
			itemCount++

			// Call RPC to save item
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			_, err := c.SaveItem(ctx, &pb.SaveItemRequest{Item: &item})
			cancel()
			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v",
					item, err)
			}
		}
	}()

	return out, nil
}

func ItemSaverAsync(host string) error {
	c, err := rpcsupport.NewClient(host)
	if err != nil {
		return err
	}

	mqSimple := rabbitmq.NewRabbitMQSimple("xcar")
	deliverChan,err := mqSimple.ConsumeSimple()
	if err!=nil{
		log.Println(err)
	}
	go func() {
		itemCount := 0
		var item pb.Item
		for {
			if mqSimple.Conn.IsClosed() {
				log.Println("RabbitMQ连接断开，重新连接...")
				mqSimple.DoConnect()
				deliverChan,err = mqSimple.ConsumeSimple()
				if err!=nil{
					log.Println(err)
				}
			}
			deliver:= <- deliverChan
			log.Println(deliver.Body)
			if deliver.Body == nil{
				continue
			}
			errJson:= json.Unmarshal(deliver.Body,&item)
			if errJson != nil{
				log.Println(errJson)
				panic(errJson)
			}
			log.Printf("Item Saver: got item "+
				"#%d: %s", itemCount, item.Car)
			itemCount++

			// Call RPC to save item
			ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
			_, err := c.SaveItem(ctx, &pb.SaveItemRequest{Item: &item})
			cancel()

			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v",
					item, err)
			}
		}
	}()

	return nil
}
