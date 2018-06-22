package engine

type ParserFunc func(contents []byte,url string) ParserResult

type Parser interface {
	Parse(contents []byte,url string) ParserResult
	Serialize() (name string, args interface{}) //name为函数的名字，args是该函数的参数，这些都需要进行序列化
}

type Request struct {
	Url string
	//ParserFunc ParserFunc  //调用RPC server的函数，由于ParserFunc函数不能在网络上进行传输，所以需要定义一个接口，里面包含对函数进行序列化的函数
	Parser Parser //调用RPC，需要将此函数转成序列化的结果
}

type ParserResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url 	 string//用于存储URL，可以有实际用途的数据
	Type     string//elasticsearch的type(表名)
	Id 		 string//elasticsearch的ID，用于数据层面的去重
	Payload  interface{}
}
type NilParserFunc struct {

}

func (NilParserFunc) Parse(_ []byte, _ string) ParserResult {
	return ParserResult{}
}

func (NilParserFunc) Serialize() (name string, args interface{}) {
	return "NilParser",nil
}

type FuncParser struct {
	parser ParserFunc
	name string
}

func (f *FuncParser) Parse(contents []byte, url string) ParserResult {
	return f.parser(contents,url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

//因为func (f *FuncParser) Parse、func (f *FuncParser) Serialize()是指针接收者，所以*FuncParser应该是指针
func NewFuncParser(p ParserFunc,name string) *FuncParser {
	return &FuncParser{
		parser:p,
		name:name,
	}
}
