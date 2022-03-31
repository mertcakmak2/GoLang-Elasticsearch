package config

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

var Client *elastic.Client

func CreateElasticConnection() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, _, pingErr := client.Ping("http://elasticsearch:9200").Do(context.TODO())
	if pingErr != nil {
		fmt.Println(pingErr.Error())
	}
	indexExists, err := client.IndexExists("user").Do(context.TODO())
	if err != nil {
		fmt.Println(pingErr.Error())
	}
	if indexExists {
		fmt.Println("exist index name")
	}
	res, err := client.CreateIndex("user").Do(context.TODO())
	if err != nil {
		fmt.Println(pingErr.Error())
	}
	if !res.Acknowledged {
		fmt.Println("Acknowledged error")
	}
	return client, err
}
