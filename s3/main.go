package main

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/k0kubun/pp"
)

func main() {

	bucketName := "xxxx"

	// Befor export env AWS_ACCESS_KEY, AWS_SECRET_ACCESS_KEY
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err.Error())
	}

	// open Bucket
	s := s3.New(auth, aws.APNortheast)
	bucket := s.Bucket(bucketName)

	oldPath := "/test/sample.txt"
	newPath :=  "/test/sample2.txt"

	// オブジェクト作成
	err = bucket.Put(oldPath, []byte("goamz example2"), "text/plain", s3.BucketOwnerFull)
	if err != nil {
		panic(err.Error())
	}

	// オブジェクトコピー
	err2 := bucket.Copy(oldPath, newPath, s3.BucketOwnerFull)
	if err2 != nil {
		panic(err2.Error())
	}

	// バケット内のすべてのキー取得
	list, _ := bucket.GetBucketContents()
	pp.Println(list)

	// プレフィックスを付けて取得
	listWithPrefix, _ := bucket.List("test/sample", "/", "", 10)
	pp.Println(listWithPrefix)
}
