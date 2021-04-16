package main

import (
	"github.com/minio/minio-go/v6"
	"log"
)

func main() {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln("new err", err)
	}
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		log.Fatalln("list err", err)
	}

	for _, b := range buckets {
		log.Println("bucket name is ", b.Name)
	}

}