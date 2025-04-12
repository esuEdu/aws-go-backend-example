package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Region          string
	Bucket          string
	AccessKeyID     string
	SecretAccessKey string
}

func LoadS3Config() S3Config {
	env := LoadEnv()

	return S3Config{
		Region:          env.AWSRegion,
		Bucket:          env.AWSBucket,
		AccessKeyID:     env.AccessKey,
		SecretAccessKey: env.SecretAccessKey,
	}
}

func InitS3Client(cfg *S3Config) *s3.Client {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKeyID, cfg.SecretAccessKey, "",
		)),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}
	return s3.NewFromConfig(awsCfg)
}
