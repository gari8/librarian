package infra

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gari8/librarian/config"
	"io"
)

type S3Client struct {
	Uploader *manager.Uploader
	Bucket   string
}

type IAws interface {
	Setup() (aws.Config, error)
}

func NewS3Client(ctx context.Context, ia IAws) (*S3Client, error) {
	conf, err := config.ReadConfig(ctx)
	if err != nil {
		return nil, err
	}
	cfg, err := ia.Setup()
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = true
	})
	return &S3Client{
		Uploader: manager.NewUploader(client),
		Bucket:   conf.AWS.Bucket,
	}, nil
}

func (sc S3Client) UploadFile(filename string, f io.Reader) (string, error) {
	result, err := sc.Uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(sc.Bucket),
		Key:         aws.String(filename),
		ACL:         types.ObjectCannedACLPublicRead,
		Body:        f,
		ContentType: aws.String(""),
	})
	if err != nil {
		return "", err
	}
	return result.Location, nil
}
