package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/viper"
)

type AWSSession struct {
	*session.Session
}

func NewAWSSession() *AWSSession {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(viper.GetString("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(
				viper.GetString("AWS_ACCESS_KEY_ID"),
				viper.GetString("AWS_SECRET_ACCESS_KEY"),
				"",
			),
		},
	))

	return &AWSSession{sess}
}
