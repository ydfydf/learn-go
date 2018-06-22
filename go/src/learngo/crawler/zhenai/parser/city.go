package parser

import (
	"learngo/crawler/engine"
	"regexp"
	"learngo/crawler_distributed/config"
)

var ProfileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`)
var cityUrlRe = regexp.MustCompile(`"(http://www.zhenai.com/zhenghun/[0-9[a-z][A-Z]/[\d]+)"`)

func ParseCity(contents []byte,_ string) engine.ParserResult{
	//re := regexp.MustCompile(cityRe)
	match := ProfileRe.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}
	for _,m := range match{
		url := string(m[1])
		name := string(m[2])
		//result.Items = append(result.Items,"User " + name)
		result.Requests = append(result.Requests,engine.Request{
			Url: url,
			//这里使用函数式编程的一个常用技巧，通过匿名函数改写函数，使函数对外接口保持一致
			//ParserFunc: func(contents []byte) engine.ParserResult {
				//return ParseProfile(contents,url,name)
			Parser:NewProfileParser(name),
		})
	}
	match = cityUrlRe.FindAllSubmatch(contents,-1)
	for _,m := range match{
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(ParseCity,config.ParseCity),
		})
	}


	return result
}
