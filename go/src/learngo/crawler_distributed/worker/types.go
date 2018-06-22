package worker

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
	"learngo/crawler/zhenai/parser"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	FunctionName string //函数的名字
	Args interface{} //该函数的参数
}



//定义RPC调用函数，下面这个函数虽然满足RPC对函数的要求，但是engine.Request里面有接口Parser，Parser不能在网络上传递，
//所以需要定义一个和engine.Request差不多的结构，让此结构能在网络上传输，达到调用RPC服务的目的
//func (CrawlService) Process(req engine.Request,result *engine.ParserResult) error{}

type Request struct {
	Url string
	Parser SerializedParser
}

func SerializeRequest(r engine.Request) Request{
	name, args := r.Parser.Serialize()
	return Request{
		Url:r.Url,
		Parser: SerializedParser {
			FunctionName:name,
			Args:args,
		},
	}
}

func DeserializeRequest(r Request) (engine.Request, error){
	parser , err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{},err
	}else {
		return engine.Request{
			Url: r.Url,
			Parser: parser,
		}, nil
	}

}

type ParseResult struct {
	Items []engine.Item
	Requests []Request
}

func SerializeResult(r engine.ParserResult) ParseResult{
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests{
		result.Requests = append(result.Requests,SerializeRequest(req))
	}

	return result
}

func DeserializeResult(r ParseResult) (engine.ParserResult,error){
	result := engine.ParserResult{
		Items: r.Items,
	}
	for _, req := range r.Requests{
		engineReq , err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v",err)
			continue
		}
		result.Requests = append(result.Requests,engineReq)
	}

	return result, nil
}

//注意这里的go语言编程技巧，engine.Parser是一个接口，后面返回的都是结构体，所以在返回时就为engine.Parser接口进行赋值，engine.Parser就能使用结构体的参数与方法
func deserializeParser(p SerializedParser) (engine.Parser,error){
	//第一种方法，将所有函数名与其对应ParserFunc注册到一张全局的map里，map[string]ParserFunc，然后进行匹配待用
	//第二种switch case
	switch p.FunctionName {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList,config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity,config.ParseCity), nil
	case config.ParseProfile:
		if userName , ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		}else {
			return nil , fmt.Errorf("invalid arg: %v",p.Args)
		}

	case config.NilParser:
		return engine.NilParserFunc{}, nil
	default:
		return nil, errors.New("unkonow parser name")
	}

}
