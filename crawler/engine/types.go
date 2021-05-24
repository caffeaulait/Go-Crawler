package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
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
		name: name,
	}
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
	Id string
	Url string
	Type string
	Payload interface{}
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type Scheduler interface {
	ReadyNotifier
	Submit(r Request)
	GetWorker() chan Request
	Run()
}

type NilParser struct {

}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}
