package engine

import "crawler/crawler_distributed/config"

// ParserFunction 解析器函数
type ParserFunction func(contents []byte, url string) (result ParseResult)

// Parser 解析器
type Parser interface {
	Parse(contents []byte, url string) (result ParseResult)
	Serialize() (name string, args interface{})
}

// Request 请求
type Request struct {
	Url    string
	Parser Parser
}

// Item 项目
type Item struct {
	Url     string
	Index   string
	Id      string
	Payload interface{}
}

// ParseResult 解析结果
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// NilParser 空解析器
type NilParser struct{}

// Parse 解析
func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

// Serialize 序列化
func (NilParser) Serialize() (string, interface{}) {
	return config.NilParser, nil
}

// FunctionParser 函数解析器
type FunctionParser struct {
	parser ParserFunction
	name   string
}

// Parse 解析
func (p *FunctionParser) Parse(contents []byte, url string) ParseResult {
	return p.parser(contents, url)
}

// Serialize 序列化
func (p *FunctionParser) Serialize() (string, interface{}) {
	return p.name, nil
}

// NewFunctionParser 创建函数解析器
func NewFunctionParser(parser ParserFunction, name string) *FunctionParser {
	return &FunctionParser{
		parser: parser,
		name:   name,
	}
}
