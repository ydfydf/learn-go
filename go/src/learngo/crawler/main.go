package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler/persist"
)

//由于珍爱网是gbk编码，而go语言是utf-8，获取的数据是乱码，进行编码转换，需要下载两个非go语言官方标准库，它们是golang.org/x下的库
//使用命令：gopm get -g -v golang.org/x/text


//由于很多网站的编码格式不统一，也有可能不是GBK，所以utf8Reader := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())这行代码的局限性就很大了，
//为了适应各种编码到我们的go语言系统中进行UTF-8转换，需要检测网站的编码，下载包：gopm get -g -v golang.org/x/net/html

const seedurl = "http://www.zhenai.com/zhenghun"

func main() {
	var seed engine.Request
	seed = engine.Request{
		Url : seedurl,
		ParserFunc : parser.ParseCityList,
	}
	itemChan , err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	//engine.SimpleEngine.Run(seed)
	//engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//}.Run(seed)
	// 因为engine.ConcurrentEngine是指针接收者，所以不能直接使用，需先定义变量

	//重构后Scheduler: &scheduler.QueuedScheduler{}和Scheduler: &scheduler.SimpleScheduler{}都能直接运行，这就是重构的目的
	//因为SimpleScheduler和QueuedScheduler都实现了接口Scheduler里的函数，根据go语言中的duck typing概念，使用者不用去管实现者是怎么实现的，直接使用就行了
	e := engine.ConcurrentEngine{//这里的e的数据类型是接口，此为接口变量赋值
		Scheduler: &scheduler.QueuedScheduler{},//Scheduler: &scheduler.SimpleScheduler{},也是接口变量赋值，确定使用哪个实现的方法
		WorkerCount: 100,
		ItemChan: itemChan, //persist.ItemSaver(),这里会先调用persist.ItemSaver()来生成一个chan interface{}
		RequestProcessor: engine.Worker,
	}
	e.Run(seed)//接口变量e调用engine.ConcurrentEngine为其实现的方法
}
