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
	fmt.Printf("info::%+v\n", info)
	fmt.Println("code::", code)

	version, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Println("version::", version)

	rsp, err := client.Index().Index("info").Id("1").BodyString("hello world").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id, rsp.Index, rsp.Type)

	get, err := client.Get().Index("info").Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get.Found {
		fmt.Println(get.Id, get.Version)
	}

	del, err := client.Delete().Index("info").Id("1").Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(del.Version)
}
