package engine

import (
	"log"
)
type SimpleEngine struct {} //包装一下

func (e SimpleEngine) Run(seeds ...Request){//...为可变参数列表
	var requests []Request
	for _, r := range seeds{
		requests = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult ,err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests,parseResult.Requests...)//parserResult.Requests...语法表示将slice parserResult.Requests的所有值加到slice requests里去，如果不这样的话就只能
		//requests = append(requests,parserResult.Requests[0],parserResult.Requests[1],parserResult.Requests[2])这样来写

		for _,item := range parseResult.Items{
			log.Printf("Got item %v",item)
		}


	}
}


