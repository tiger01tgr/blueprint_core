package s3

import (
	"backend/config/s3Client"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadUserSubmissionVideo(fileName string, file *multipart.File) (string, error) {
	return uploadFile("user-submissions/video-"+fileName+".webm", "video/webm", *file)
}

func UploadResume(fileName string, file *multipart.File) (string, error) {
	return uploadFile("resumes/resume-"+fileName+".pdf", "application/pdf", *file)
}

func UploadCompanyLogo(fileName string, file *multipart.File) (string, error) {
	return uploadFile("company-logos/logo-"+fileName+".png", "image/png", *file)
}

func uploadFile(key string, contentType string, file multipart.File) (string, error) {
	s3Client := s3Client.GetS3()

	uploader := manager.NewUploader(s3Client.S3)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s3Client.BucketName),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Body:        file,
	})
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s3Client.BucketName, *result.Key)
	// The result contains information about the uploaded object, including the key
	return url, nil
}
