package parser

import (
	"learngo/crawler/engine"
	"regexp"
	"learngo/crawler_distributed/config"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte,_ string) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}
	for _,m := range matches{
		//result.Items = append(result.Items,"City "+ string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url : string(m[1]),
			//注意Parser是接口类型，engine.NewFuncParser返回一个结构体，所以这里是接口赋值
			Parser: engine.NewFuncParser(ParseCity,config.ParseCity),
		})
		//fmt.Printf("City:%s   URL:%s\n",m[2],m[1])
	}
	return result
}