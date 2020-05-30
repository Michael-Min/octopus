package rpcsupport

import (
	"Michael-Min/octopus/engine"
	"Michael-Min/octopus/persist"
	pb "Michael-Min/octopus/proto"
	t "Michael-Min/octopus/worker"
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)


type RPCService struct {
	Client *elastic.Client
	Index  string
}


func (s *RPCService) Process(
	ctx context.Context, req *pb.ProcessRequest)(*pb.ProcessResult,error){
	engineReq, err := t.DeserializeRequest(req)
	if err != nil {
		return nil, err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		fmt.Printf("==engineResult: error:%s \n",err)
		return nil, err
	}
	fmt.Printf("==engineResult: request %d,item %d \n",len(engineResult.Requests),len(engineResult.Items))
	var result = t.SerializeResult(engineResult)

	return &result, nil
}


func (s *RPCService) SaveItem(
	ctx context.Context, item *pb.SaveItemRequest) (*pb.SaveItemResult, error) {
	err := persist.Save(s.Client, s.Index, item.Item)
	log.Printf("Item %v saved.", item.Item)
	if err != nil {
		log.Printf("Error saving item %v: %v",
			item, err)
	}
	return &pb.SaveItemResult{}, err
}
