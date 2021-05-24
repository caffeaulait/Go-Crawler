package client

import (
	"go_crawler/crawler/engine"
	"go_crawler/crawler_distributed/config"
	"go_crawler/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
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
			//Call rpc to save Item
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver Error: saving item :%v: %v",item, err)
			}
		}
	}()
	return out, nil
}
