package persist

import (
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"go_crawler/crawler/engine"
	"go_crawler/crawler/model"
	"golang.org/x/net/context"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	const index = "dating_test"
	err = Save(expected, client, index)
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index(index).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	profile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = profile
	//verify
	if actual != expected {
		t.Errorf("got %v, but expected %v", actual, expected)
	}
}
