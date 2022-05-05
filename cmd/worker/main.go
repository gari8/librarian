package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gari8/librarian/config"
	"github.com/gari8/librarian/infra"
	"log"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest() (string, error) {
	_ = setup()
	return "", nil
}

type edinet interface {
}

type s3Client interface {
}

type redisClient interface {
}

type set struct {
	edinet
	s3Client
	redisClient
}

func setup() *set {
	c, err := config.Load()
	if err != nil {
		log.Fatalln("cannot load config")
	}
	h := infra.NewHttp()
	e := infra.NewEdiNet(h)
	a := infra.NewAWS(c)
	s, err := infra.NewS3Client(c.SetConfig(context.Background()), a)
	if err != nil {
		log.Fatalln("cannot create s3 client")
	}
	r := infra.NewRedisClient()
	return &set{
		edinet:      e,
		s3Client:    s,
		redisClient: r,
	}
}
