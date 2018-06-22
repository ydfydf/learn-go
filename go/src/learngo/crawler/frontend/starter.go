package main

import (
	"net/http"
	"learngo/crawler/frontend/controller"
	"fmt"
	"learngo/crawler_distributed/config"
)

func main() {
	http.Handle("/",http.FileServer(http.Dir("crawler/frontend/view")))//使用http.FileServer来提供静态内容，css，js，图片，首页等
	http.Handle("/search",controller.CreateSearchResultHandle("crawler/frontend/view/template1.html"))
	err := http.ListenAndServe(fmt.Sprintf(":%d",config.FrontPort),nil)
	if err != nil {
		panic(err)
	}
}
