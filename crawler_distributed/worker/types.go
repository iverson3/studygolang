package worker

import (
	"errors"
	"log"
	"studygolang/crawler/engine"
	"studygolang/crawler/zhenai/parser"
	"studygolang/crawler_distributed/config"
)

// 序列化ParserFunc 以便其能在网络上传输
// 序列化的结果：  {"ParseCityList", nil}   {"ProfileParser", {username, sex}}
type SerializedParser struct {
	Name string         // 函数名
	Args interface{}    // 函数参数
}

type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}


// 将engine.Request 序列化成 Request 以便其能在网络上传递
// engine.Request结构中包含函数 无法在网络上传递
// Request 则通过SerializedParser序列化ParserFunc这个解析函数 使其能在网络上传递
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url:    r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

// 将engine.ParseResult 序列化成 ParseResult 以便其能在网络上传递
// engine.ParseResult中包含了很多engine.Request 而engine.Request结构中包含了函数 无法在网络上传递
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items:    r.Items,
	}
	// 循环调用SerializeRequest 将所有的engine.Request 序列化成 Request
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

// 反序列化
// 将Request 反序列化成 engine.Request
func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

// 反序列化
// 将ParseResult 反序列化成 engine.ParseResult
func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	// 循环调用DeserializeRequest 将所有的Request 反序列化成 engine.Request
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

// 反序列化Parser
// 将序列化的parser字符串 如 {"ParseCityList", nil}   {"ProfileParser", {username, sex}}
// 反序列化成可以调用的Parser函数
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		// panic: interface conversion: interface {} is map[string]interface {}, not parser.ProfileParser
		//args := p.Args.(parser.ProfileParser)

		// p.Args is map[Sex:男士 UserName:宇墨]
		args := p.Args.(map[string]interface{})
		return parser.NewProfileParser(args["UserName"].(string), args["Sex"].(string)), nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
























