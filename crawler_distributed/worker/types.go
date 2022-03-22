package worker

import (
	"crawler/crawler/engine"
	"crawler/crawler/zhenai/parser"
	"crawler/crawler_distributed/config"
	"errors"
	"fmt"
	"log"
)

// SerializedParser 序列化解析器
type SerializedParser struct {
	Name string
	Args interface{}
}

// Request 请求
type Request struct {
	Url    string
	Parser SerializedParser
}

// ParseResult 解析结果
type ParseResult struct {
	Requests []Request
	Items    []engine.Item
}

// SerializeRequest 序列化请求
func SerializeRequest(request engine.Request) Request {
	name, args := request.Parser.Serialize()
	return Request{
		Url: request.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

// SerializeParseResult 序列化解析结果
func SerializeParseResult(result engine.ParseResult) ParseResult {
	res := ParseResult{
		Items: result.Items,
	}
	for _, request := range result.Requests {
		req := SerializeRequest(request)
		res.Requests = append(res.Requests, req)
	}
	return res
}

// DeserializeParser 反序列化解析器
func DeserializeParser(sp SerializedParser) (engine.Parser, error) {
	switch sp.Name {
	case config.CityListParser:
		p := engine.NewFunctionParser(parser.ParseCityList, config.CityListParser)
		return p, nil
	case config.CityParser:
		p := engine.NewFunctionParser(parser.ParseCity, config.CityParser)
		return p, nil
	case config.ProfileParser:
		name, ok := sp.Args.(string)
		if !ok {
			err := fmt.Errorf("invalid arg: %v", sp.Args)
			return nil, err
		}
		p := parser.NewProfileParser(name)
		return p, nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		err := errors.New("unknown parser name")
		return nil, err
	}
}

// DeserializeRequest 反序列化请求
func DeserializeRequest(request Request) (engine.Request, error) {
	parser, err := DeserializeParser(request.Parser)
	if err != nil {
		return engine.Request{}, err
	}

	req := engine.Request{
		Url:    request.Url,
		Parser: parser,
	}
	return req, nil
}

// DeserializeParseResult 反序列化解析结果
func DeserializeParseResult(result ParseResult) engine.ParseResult {
	res := engine.ParseResult{
		Items: result.Items,
	}
	for _, request := range result.Requests {
		req, err := DeserializeRequest(request)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}

		res.Requests = append(res.Requests, req)
	}
	return res
}
