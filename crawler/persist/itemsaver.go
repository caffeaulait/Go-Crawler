package persist

import (
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"go_crawler/crawler/engine"
	"golang.org/x/net/context"
	"log"
)

/*
run elasticsearch in docker:
	docker pull docker.elastic.co/elasticsearch/elasticsearch:7.4.2
	docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.4.2
*/

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			itemCount++
			log.Printf("Got Profile #%d: %v",itemCount, item)
			err := Save(item, client, index)
			if err != nil {
				log.Printf("Item Saver Error: saving item :%v: %v",item, err)
			}
		}
	}()
	return out, nil
}

func Save(item engine.Item, client *elastic.Client, index string) (err error){ //must turn off sniff in docker
	if item.Type == "" {
		return  errors.New("must have type")
	}
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return  nil
}

