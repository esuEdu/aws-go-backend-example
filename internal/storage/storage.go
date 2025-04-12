package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

type S3Uploader struct {
	client *s3.Client
	Bucket string
}

func NewS3Uploader(client *s3.Client, bucket string) S3Uploader {
	return S3Uploader{
		client: client,
		Bucket: bucket,
	}
}

func (u *S3Uploader) UploadFile(file *multipart.FileHeader) (string, string, error) {
	ctx := context.Background()

	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	fileID := uuid.New().String() + ext

	_, err = u.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(u.Bucket),
		Key:           aws.String(fileID),
		Body:          src,
		ContentLength: &file.Size,
		ContentType:   aws.String(file.Header.Get("Content-Type")),
		ACL:           types.ObjectCannedACLPrivate,
	})

	if err != nil {
		return "", "", fmt.Errorf("failed to upload to S3: %w", err)
	}

	// Generate a presigned URL for GET
	presignClient := s3.NewPresignClient(u.client)

	presignParams := &s3.GetObjectInput{
		Bucket: aws.String(u.Bucket),
		Key:    aws.String(fileID),
	}

	presignedURL, err := presignClient.PresignGetObject(ctx, presignParams,
		func(opts *s3.PresignOptions) {
			opts.Expires = 15 * time.Minute
		},
	)

	if err != nil {
		return "", "", fmt.Errorf("failed to create presigned URL: %w", err)
	}

	return fileID, presignedURL.URL, nil

}
