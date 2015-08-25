package main

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

func main() {

	bucketName := "xxxxx"

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

	err = bucket.Put(oldPath, []byte("goamz example"), "text/plain", s3.BucketOwnerFull)
	if err != nil {
		panic(err.Error())
	}

	// copy object
	err2 := bucket.Copy(oldPath, newPath, s3.BucketOwnerFull)
	if err2 != nil {
		panic(err2.Error())
	}
}
