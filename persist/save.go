package persist

import (
	pb "Michael-Min/octopus/proto"
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
)

func Save(
	client *elastic.Client, index string,
	item *pb.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	return err
}

func FromJsonObj(o interface{}) (pb.Profile, error) {
	var profile pb.Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
