package engine

import "studygolang/crawler_distributed/config"

type ParserFunc func(contents []byte, url string) ParseResult 

// Parser接口的实现者有： NilParser  FuncParser  ProfileParser
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})    // 序列化ParserFunc
}

type Request struct {
	Url string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items []Item
}

type Item struct {
	Type string    // elastic对应的type配置项
	Id string
	Url string
	Payload interface{}
}

type NilParser struct {}

func (n NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}


type FuncParser struct {
	parser ParserFunc
	name string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}





























