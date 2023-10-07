package s3Client

import (
	"context"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	S3   *S3Client
	once sync.Once
)

type S3Client struct {
	S3         *s3.Client
	BucketName string
}

func GetS3() *S3Client {
	once.Do(InitS3)
	return S3
}

func InitS3() {
	accessKey, err1 := os.LookupEnv("AWS_ACCESS_KEY_ID")
	secretKey, err2 := os.LookupEnv("AWS_SECRET_KEY_ID")
	bucketName, err3 := os.LookupEnv("AWS_BUCKET_NAME")

	if !err1 {
		panic("S3_ACCESS_KEY_ID not found")
	}
	if !err2 {
		panic("S3_SECRET_ACCESS_KEY not found")
	}
	if !err3 {
		panic("S3_BUCKET_NAME not found")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")))
	if err != nil {
		panic("Failed to load AWS config")
	}

	s3Client := s3.NewFromConfig(cfg)
	S3 = &S3Client{s3Client, bucketName}
}
