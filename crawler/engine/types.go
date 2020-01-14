package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url string
	ParserFunc ParserFunc
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

func NilParser([]byte) ParseResult {
	return ParseResult{}
}





























