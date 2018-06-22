package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"learngo/crawler/engine"
	"errors"
)

func ItemSaver(index string) (chan engine.Item,error) {
	//http.Post()
	//使用elastic client包来进行elasticsearch操作
	client,err := elastic.NewClient(
		elastic.SetURL("http://192.168.78.131:9200"),
		//mist turn off sniff in docker,因为sniff是用于跑在本机的elasticsearch进行维护的，elasticsearch跑在docker上面，访问不了，不能sniff，所以关闭这个功能
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil,err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for{
			item := <- out
			log.Printf("Item Saver: got item #%d: %v",itemCount,item)
			itemCount ++

			err := Save(client,item,index)
			if err != nil {
				log.Printf("Item Saver ERROR; saving item %v: %v\n",item,err)
			}

		}
	}()
	return out,nil
}

func Save(client *elastic.Client,item engine.Item,index string) (err error){
	if item.Type == "" {
		return errors.New("must spply Type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).//这里不加.Id，系统会自动分配一个ID
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err =indexService.Do(context.Background())//函数参数是接口类型，传入值可以是实现了该接口的所有类型的变量，只要该类型实现了该接口，就能传入

	if err != nil {
		return err
	}
	//fmt.Printf("%+v",resp)//%+v会将结构体中的字段名打出来
	return nil

}
