package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	fmt.Println("--- AWS Cloud Automation Started ---")

	// Point to our Local Mock Cloud
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: "http://localhost:4566",
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	client := s3.NewFromConfig(cfg)
	bucketName := "career-comeback-bucket"

	// Create Bucket
	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
    
    if err != nil {
        fmt.Println("Status: Bucket exists or Mock Server not ready.")
    } else {
        fmt.Println("Success: Bucket Created!")
    }

	// Upload File
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("hello.txt"),
		Body:   strings.NewReader("Industry Level Automation Proof"),
	})

	if err == nil {
		fmt.Println("Success: File 'hello.txt' uploaded to Mock Cloud!")
	}
}

