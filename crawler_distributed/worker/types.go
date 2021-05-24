package worker

import (
	"fmt"
	"github.com/pkg/errors"
	"go_crawler/crawler/config"
	"go_crawler/crawler/engine"
	"go_crawler/crawler/parser"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Items 	 []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func DeSerializeRequest(r Request) (engine.Request, error) {
	parser, err := DeSerializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{Url: r.Url, Parser: parser}, nil
}

func DeSerializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if name, ok := p.Args.(string); ok {
			return parser.NewProfileParser(name), nil
		} else {
			return nil, fmt.Errorf("invalid args: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeSerializeResult(r ParseResult) (engine.ParseResult, error) {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeSerializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result, nil
}

