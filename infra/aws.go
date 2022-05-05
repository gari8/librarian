package infra

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	conf "github.com/gari8/librarian/config"
)

type AWS struct {
	conf.Config
}

func NewAWS(c conf.Config) *AWS {
	return &AWS{c}
}

func (a AWS) Setup() (aws.Config, error) {
	var cfg aws.Config
	var err error
	fmt.Printf("::%+v\n", a.Config.AWS)
	if true {
		cred := credentials.NewStaticCredentialsProvider(a.Config.AWS.AccessKeyId, a.Config.AWS.SecretKey, "")
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					PartitionID: "aws",
					URL:         "http://localhost:9000",
				}, nil
			})),
			config.WithCredentialsProvider(cred))
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO())
	}
	return cfg, err
}
