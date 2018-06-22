package client

import (
	"log"
	"learngo/crawler/engine"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/config"
)

func ItemSaver(host string) (chan engine.Item,error) {
	rpcClient, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil,err
	}
	out := make(chan engine.Item)
	for i := 0;i < 10 ; i++ {
		go func() {
			itemCount := 0
			for {
				item := <-out
				log.Printf("Item Saver: got item #%d: %v", itemCount, item)
				itemCount ++

				//Call RPC server to save item
				result := ""
				err = rpcClient.Call(config.ItemSaverRpc, item, &result)
				if err != nil {
					log.Printf("Item Saver ERROR; saving item %v: %v\n", item, err)
				}
			}
		}()
	}
	return out,nil
}
