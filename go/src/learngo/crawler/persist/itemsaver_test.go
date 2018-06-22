package persist

import (
	"testing"
	"learngo/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
	"learngo/crawler/engine"
)

func TestItemSaver(t *testing.T) {
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


	//TODO: Start up elastic search in docker firstly and the test can run
	client ,err := elastic.NewClient(
		elastic.SetURL("http://192.168.78.130:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	const index = "dating_test"

	err = Save(client,expected,index)
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	//t.Logf("%+v",resp)
	t.Logf("%s",resp.Source)

	var actual engine.Item
	//json反序列化
	err = json.Unmarshal(*resp.Source,&actual)//因为Source *json.RawMessage，所以这里的*resp.Source必须加*号
	if err != nil {
		panic(err)
	}

	actualProfile , _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v",actual,expected)
	}
}