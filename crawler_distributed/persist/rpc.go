package persist

import (
	"crawler/crawler/engine"
	"crawler/crawler/persist"
	"github.com/olivere/elastic/v7"
	"log"
)

// ItemSaverService 项目存储服务
type ItemSaverService struct {
	Client *elastic.Client
}

// Save 存储项目
func (service *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(service.Client, item)
	log.Printf("Item saved: %+v", item)
	if err == nil {
		*result = "ok"
	} else {
		*result = "error"
		log.Printf("Error saving item %+v: %v", item, err)
	}
	return err
}
