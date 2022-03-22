package client

import (
	"crawler/crawler/engine"
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/rpcsupport"
	"log"
)

// ItemSaver 项目存储器
func ItemSaver(address string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(address)
	if err != nil {
		return nil, err
	}

	in := make(chan engine.Item)

	go func() {
		itemCount := 0
		result := ""

		for {
			item := <-in
			log.Printf("Item Saver: got item #%d: %+v", itemCount, item)
			itemCount++

			err = client.Call(config.ItemSaverMethod, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item %+v: %v", item, err)
			}
		}
	}()

	return in, nil
}
