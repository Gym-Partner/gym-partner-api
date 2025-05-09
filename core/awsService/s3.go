package awsService

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	*s3.S3
}

func NewAWSS3(sess *AWSService) *S3Service {
	s3Session := s3.New(sess.Session)
	return &S3Service{s3Session}
}
