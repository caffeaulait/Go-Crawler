package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go_crawler/crawler_distributed/config"
	"go_crawler/crawler_distributed/persist"
	"go_crawler/crawler_distributed/rpcsupport"
	"log"
)


var port = flag.Int("port", 0, "the port for me to listen")

func main()  {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index: index,
	})
}
