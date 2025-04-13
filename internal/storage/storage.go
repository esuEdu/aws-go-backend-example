package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
)

type S3Uploader struct {
	Bucket string
	Region string
}

func NewS3Uploader(bucket string, region string) S3Uploader {
	return S3Uploader{
		Bucket: bucket,
		Region: region,
	}
}

func (u *S3Uploader) UploadFile(file *multipart.FileHeader) (string, string, error) {
	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	fileID := uuid.New().String() + ext

	uploadURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.Bucket, u.Region, fileID)

	req, err := http.NewRequest("PUT", uploadURL, src)
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", file.Header.Get("Content-Type"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("failed to upload file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return "", "", fmt.Errorf("upload failed: %s", string(body))
	}

	return fileID, uploadURL, nil
}
