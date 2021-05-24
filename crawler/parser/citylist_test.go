package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents, "")
	const resultSize = 470
	expectedUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng"}
	//expectedCities := []string{"阿坝", "阿克苏", "阿拉善盟"}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but have %d requests", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items, but have %d items", resultSize, len(result.Items))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
	//for i, city := range expectedCities{
	//	if result.Items[i].(string) != city {
	//		t.Errorf("expected city #%d: %s; but was %s", i, city, result.Items[i].(string))
	//	}
//}
}