package elasticsearch

import (
	"chs/config"
	"chs/modules/elasticsearch/mappings"
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var client *elastic.Client

func InitEs() {
	if client != nil {
		log.Println("es has already init")
		return
	}
	host := config.Conf.Get("es.host")
	if host == "" {
		log.Fatalf("es init fail [0]: config host never gotten")
	}
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host.(string)))
	if err != nil {
		log.Fatalf("es init fail [1]: %v", err)
	}
	exists, err := client.IndexExists("replies").Do(context.Background())
	if err != nil {
		log.Fatalf("es init fail [2]: %v", err)
	}
	if !exists {
		createIndex()
	}
}

func createIndex() bool {
	//TODO CREATE MULTI-INDEX
	response, err := client.CreateIndex("replies").Body(mappings.GetMapping()).Do(context.Background())
	if err != nil {
		log.Fatalf("es create index fail [1]: %v", err)
	}
	log.Fatalf("es create index [2]: %v", response.Acknowledged)
	return response.Acknowledged
}

func DeleteIndex(index ...string) bool {
	response, err := client.DeleteIndex(index...).Do(context.Background())
	if err != nil {
		log.Printf("es delete index failed, err: %v\n", err)
	}
	log.Fatalf("es create delete index: %v", response.Acknowledged)
	return response.Acknowledged
}
