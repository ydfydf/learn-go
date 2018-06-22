package view

import (
	"testing"
	"os"
	"learngo/crawler/frontend/model"
	"learngo/crawler/engine"
	model2 "learngo/crawler/model"
	"learngo/crawler/frontend/view"
)

func TestSearchResultView_Render(t *testing.T){
	page := model.SearchResult{}

	item := engine.Item{
		Url:"http://album.zhenai.com/u/110409917",
		Type:"zhenai",
		Id:"110409917",
		Payload:model2.Profile{
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
	for i := 0 ; i < 10 ;i ++ {
		page.Items = append(page.Items, item)
	}

	file, err := os.Create("template_test.html")
	if err != nil {
		panic(err)
	}
	template := view.CreateSearchResultView("template.html")

	err = template.Render(file,page)

	if err != nil {
		panic(err)
	}
}

