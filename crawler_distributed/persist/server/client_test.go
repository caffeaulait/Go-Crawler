package main

import (
	"go_crawler/crawler/engine"
	"go_crawler/crawler/model"
	"go_crawler/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/7561952810363423875",
		Type: "zhenai",
		Id: "7561952810363423875",
		Payload: model.Profile{
			Age:        39,
			Height:     162,
			Weight:     50,
			Income:     "5001-7000元",
			Gender:     "女",
			Name:       "安静的雪",
			Sign:       "双子座",
			Occupation: "运营",
			Marriage:   "单身",
			House:      "已购房",
			Origin:     "浙江杭州",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)

	if err != nil || result !="ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
