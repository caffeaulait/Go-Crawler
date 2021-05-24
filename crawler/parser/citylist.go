package parser

import (
	"go_crawler/crawler/config"
	"go_crawler/crawler/engine"
	"regexp"
)

const cityListReg = `<a href="(.*www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	reg := regexp.MustCompile(cityListReg)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches{
		url := string(m[1])
		//city := "City " + string(m[2])
		//result.Items = append(result.Items, city)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity)})
		//fmt.Printf("City: %s, URL: %s\n ", city, url)
	}
	return result
}
