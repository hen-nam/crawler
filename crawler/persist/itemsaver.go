package persist

import (
	"context"
	"crawler/crawler/engine"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
)

// ItemSaver 项目存储器
func ItemSaver() (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	in := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-in
			log.Printf("Item Saver: got item #%d: %+v", itemCount, item)
			itemCount++

			err := Save(client, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %+v: %v", item, err)
			}
		}
	}()

	return in, nil
}

// Save 存储项目
func Save(client *elastic.Client, item engine.Item) error {
	if item.Index == "" {
		return errors.New("must supply Index")
	}

	service := client.Index().Index(item.Index).BodyJson(item)
	if item.Id != "" {
		service.Id(item.Id)
	}
	_, err := service.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
