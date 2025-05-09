package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSS3 struct {
	*s3.S3
}

func NewAWSS3(sess *session.Session) *AWSS3 {
	s3Session := s3.New(sess)
	return &AWSS3{s3Session}
}
