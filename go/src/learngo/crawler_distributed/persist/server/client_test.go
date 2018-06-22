package main

import (
	"testing"
	"learngo/crawler_distributed/rpcsupport"
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"time"
	"learngo/crawler_distributed/config"
)

func TestItemSaver(t *testing.T){
	const host =  ":3333"
	go ServeRpc(host,"test1")

	time.Sleep(time.Second)

	rpcclient , err := rpcsupport.NewClient(host)
	if err != nil{
		panic(err)
	}
	var result string
	expected := engine.Item{
		Url: "http://album.zhenai.com/u/110409917",
		Type: "zhenai",
		Id: "110409917",
		Payload:model.Profile{
			"静静等待",
			"女",
			40,
			160,
			0,
			"3001-5000元",
			"未婚",
			"高中及以下",
			"监",
			"广西百色",
			"双鱼座",
			"未",
			"未购车",
		},
	}
	err = rpcclient.Call(config.ItemSaverRpc,expected,&result)
	if err != nil{
		panic(err)
	}
	if result == "ok" {
		fmt.Println("Test Passed")
	}
}
