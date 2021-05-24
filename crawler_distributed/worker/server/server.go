package main

import (
	"flag"
	"fmt"
	"go_crawler/crawler_distributed/rpcsupport"
	"go_crawler/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		&worker.CrawlService{}))
}