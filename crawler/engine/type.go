package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}
type ParseResult struct {
	Requests []Request
	Iterms []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

