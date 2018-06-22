package main

import (
	"learngo/poster"
	"fmt"
	_ "expvar"
	_ "go/types"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string{
	return r.Get(url)
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url,map[string]string{
		"name":"ydf",
		"course":"golang",
	})
}

type RetrieverPoster interface {//接口的组合
	Retriever
	Poster
	Connet()
}

func session(s RetrieverPoster) string{
	s.Post(url,map[string]string{
		"contents":"this is another fake imooc",
	})
	return s.Get(url)
}

func main() {
	var r RetrieverPoster
	r = &poster.Retriever{"it is a fake imooc"}
	fmt.Println(r.Get(url))
	fmt.Println(session(r))
	fmt.Println("===============================")

	var s Retriever
	s = &poster.Retriever{"it is a fake imooc"}
	fmt.Println(s.Get(url))
	//fmt.Println(session(s)) //因为s是Retriever接口，里面只有Get函数，所以不能调用Post函数
	fmt.Println("===============================")

	//因为poster.Retriever结构同时实现Get和Post函数，所以能使用
	p := &poster.Retriever{"it is a fake imooc"}
	fmt.Println(p.Get(url))
	fmt.Println(session(p))


}
