package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Region string
	Bucket string
}

func LoadS3Config() S3Config {
	env := LoadEnv()

	return S3Config{
		Region: env.AWSRegion,
		Bucket: env.AWSBucket,
	}
}

func InitS3Client(cfg *S3Config) *s3.Client {
	// Automatically use the EC2 instance role for credentials
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.Region))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}
	return s3.NewFromConfig(awsCfg)
}
