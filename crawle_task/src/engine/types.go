package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// 一个 空 的 NilParser
func NilParser([] byte) ParseResult  {
	return ParseResult{}
}
