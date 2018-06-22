package persist

import (
	"learngo/crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"learngo/crawler/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error{
	err := persist.Save(s.Client,item,s.Index)
	log.Printf("Item %v saved.",item)
	if err == nil {
		*result = "ok"
	}else {
		log.Printf("Error saving Item %v : %v.",item,err)
	}
	return err
}