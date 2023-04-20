package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"path/filepath"
	"postgraduate-pm-backend/utils/helper"
	"postgraduate-pm-backend/utils/zookeeper"
)

type MinioConfig struct {
	Endpoint  string `json:"endPoint"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

var config *MinioConfig

var client *minio.Client

func MinioInit() error {
	var err error

	config = &MinioConfig{}
	err = zookeeper.GetUtilsConfig("/minio", config)
	if err != nil {
		return err
	}

	client, err = minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.AccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}
	return nil
}

func createBucket(client *minio.Client, bucketName string) error {
	// 检查存储桶是否已经存在
	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		return err
	}

	// 如果存储桶不存在，则创建一个新的存储桶
	if !exists {
		err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
		fmt.Printf("Successfully created %s bucket\n", bucketName)
	} else {
		fmt.Printf("Bucket %s already exists\n", bucketName)
	}

	return nil
}

func UploadFile(bucketName string, fileHeader *multipart.FileHeader) (string, error) {
	// 打开文件
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 获取文件的大小
	fileSize := fileHeader.Size

	objectName := fmt.Sprintf("%s%s", helper.RandomString(32), filepath.Ext(fileHeader.Filename))

	// 上传文件
	_, err = client.PutObject(context.Background(), bucketName, objectName, file, fileSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return "", err
	}
	// fmt.Printf("Successfully uploaded %s to %s/%s\n", fileHeader.Filename, bucketName, objectName)
	url := fmt.Sprintf("http://%s/%s/%s", config.Endpoint, bucketName, objectName)
	return url, nil
}

func GetMinioClient() *minio.Client {
	return client
}
