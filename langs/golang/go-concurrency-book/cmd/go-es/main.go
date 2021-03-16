package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

const (
	host = "http://localhost:9200"
)

func main() {

	var err error
	client, err = elastic.NewClient(elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("info::", info)
	fmt.Println("code::", code)

	version, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Println("version::", version)
}
