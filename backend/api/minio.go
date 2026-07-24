package api

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"omciAnalyzer/global"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// GeneratePresignedURL generates a presigned URL for the given object key
func GeneratePresignedURL(objectKey string, expiration time.Duration) (string, error) {
	// Initialize MinIO client
	client, err := minio.New(global.AppConf.MinIO.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			global.AppConf.MinIO.AccessKeyID,
			global.AppConf.MinIO.SecretAccessKey,
			"",
		),
		Secure: global.AppConf.MinIO.Secure,
	})
	if err != nil {
		return "", err
	}

	// Generate presigned URL
	presignedURL, err := client.PresignedGetObject(
		context.Background(),
		global.AppConf.MinIO.Bucket,
		objectKey,
		expiration,
		nil,
	)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}

// DownloadFileFromMinio downloads a file from MinIO to a local path
func DownloadFileFromMinio(objectKey string, localPath string) error {
	// Initialize MinIO client
	client, err := minio.New(global.AppConf.MinIO.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			global.AppConf.MinIO.AccessKeyID,
			global.AppConf.MinIO.SecretAccessKey,
			"",
		),
		Secure: global.AppConf.MinIO.Secure,
	})
	if err != nil {
		return err
	}

	// Get object from MinIO
	object, err := client.GetObject(
		context.Background(),
		global.AppConf.MinIO.Bucket,
		objectKey,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return err
	}
	defer object.Close()

	// Create directory if it doesn't exist
	dir := filepath.Dir(localPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Create local file
	localFile, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer localFile.Close()

	// Copy content from MinIO to local file
	if _, err := io.Copy(localFile, object); err != nil {
		return err
	}

	return nil
}
