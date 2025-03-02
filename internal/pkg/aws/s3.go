package aws

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
)

type S3Client struct {
	client *s3.Client
}

func CreateS3Client() *S3Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println("++++++there is a issue with s3 client++++++\n", err)
	}

	client := s3.NewFromConfig(cfg)
	return &S3Client{
		client: client,
	}
}

func (s *S3Client) UploadToS3(ctx context.Context, bucketName string, objectKey string, obj io.Reader) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   obj,
	})
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "EntityTooLarge" {
			log.Printf("Error while uploading object to %s. The object is too large.\n"+
				"To upload objects larger than 5GB, use the S3 console (160GB max)\n"+
				"or the multipart upload API (5TB max).", bucketName)
		} else {
			log.Printf("Couldn't upload file to %v:%v. Here's why: %v\n",
				bucketName, objectKey, err)
		}
	} else {
		err = s3.NewObjectExistsWaiter(s.client).Wait(
			ctx, &s3.HeadObjectInput{Bucket: aws.String(bucketName), Key: aws.String(objectKey)}, time.Minute)
		if err != nil {
			log.Printf("Failed attempt to wait for object %s to exist.\n", objectKey)
		}
	}
	return err
}

func (s *S3Client) PresignGetS3Url(ctx context.Context, bucketName string, objectKey string) (string, error) {
	presignClient := s3.NewPresignClient(s.client)
	url, err := presignClient.PresignGetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
		}, s3.WithPresignExpires(time.Minute*15))
	if err != nil {
		return "", err
	}
	return url.URL, nil
}

func (s *S3Client) DeleteS3Item(ctx context.Context, bucketName string, objectKey string) (bool, error) {
	deleted := false
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}
	_, err := s.client.DeleteObject(ctx, input)
	if err != nil {
		return false, err
	} else {
		deleted = true
	}
	return deleted, nil
}
