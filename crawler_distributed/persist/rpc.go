package persist

import (
	"github.com/olivere/elastic/v7"
	"go_crawler/crawler/engine"
	"go_crawler/crawler/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error{
	err := persist.Save(item, s.Client, s.Index)
	if err == nil {
		log.Printf("Item %v saved", item)
		*result = "ok"
	} else {
		log.Printf("Error saving Item %v: %v", item, err)
	}
	return err
}
