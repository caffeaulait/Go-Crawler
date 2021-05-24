package parser

import (
	"go_crawler/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(.*album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
		`<a href="(.*www\.zhenai\.com/zhenghun/[^"]+)"[^>]*>([^<]+)</a>`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches{
		url := string(m[1])
		user := string(m[2])
		//result.Items = append(result.Items, user)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			Parser: NewProfileParser(user),
		})
	}
	////城市页面下一页
	//matches = cityUrlRe.FindAllSubmatch(contents, -1)
	//for _, m := range matches{
	//	//result.Items = append(result.Items, string(m[2]))
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url: string(m[1]),
	//		Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
	//	})
	//}
	return result
}