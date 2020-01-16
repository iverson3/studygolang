package persist

import (
	"gopkg.in/olivere/elastic.v6"
	"log"
	"studygolang/crawler/engine"
	"studygolang/crawler/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	log.Printf("Item saved: %v", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}



























