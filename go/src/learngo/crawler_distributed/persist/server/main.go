package main

import (
	"gopkg.in/olivere/elastic.v5"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/persist"
	"log"
	"learngo/crawler_distributed/config"
	"fmt"
	"flag"
)
var port = flag.Int("port",0,"input the elastic search server port that you will listen,Usage:--port=")
func main(){
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a elastic search server port")
		return
	}
	//err := ServeRpc(":1234","dating_profile")
	//if err != nil {
	//	panic(err)
	//}

	//除了上面调用一个server的方法，还有用一种偷懒式的方式log	.Fetal
	log.Fatal(ServeRpc(fmt.Sprintf(":%d",*port),config.ElasticIndex))
}

func ServeRpc(host string,index string) error {
	elasticClient, err := elastic.NewClient(
		elastic.SetURL(config.ElaticUrl),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,&persist.ItemSaverService{
		Client:elasticClient,
		Index:index,
	})

}
