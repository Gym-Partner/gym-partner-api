package core

import (
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/spf13/viper"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type ICognito interface {
	SignUp(user model.User) *Error
	SignIn(user model.User) (string, *Error)
	GetUserByToken(token string) (*string, *Error)
	DeleteUser(token string) *Error
}

type Cognito struct {
	CognitoProvider    *cognitoidentityprovider.CognitoIdentityProvider
	CognitoAppIdClient string
	Log                *Log
}

func NewCognito(log *Log) *Cognito {
	config := &aws.Config{
		Region: aws.String(viper.GetString("AWS_REGION")),
	}
	sess, err := session.NewSession(config)
	if err != nil {
		log.Error(ErrAWSCognitoCreateSession)
		return nil
	}

	client := cognitoidentityprovider.New(sess)

	return &Cognito{
		CognitoProvider:    client,
		CognitoAppIdClient: viper.GetString("APP_CLIENT_ID"),
		Log:                log,
	}
}

func (c *Cognito) SignUp(user model.User) *Error {
	userCognito := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(c.CognitoAppIdClient),
		Username: aws.String(user.Id),
		Password: aws.String(strings.TrimSpace(user.Password)),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("name"),
				Value: aws.String(user.UserName),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
		},
	}

	_, err := c.CognitoProvider.SignUp(userCognito)
	if err != nil {
		c.Log.Error(ErrAWSCognitoCreateUser)
		return NewError(http.StatusBadRequest, ErrAWSCognitoCreateUser, err)
	}

	return nil
}

func (c *Cognito) SignIn(user model.User) (string, *Error) {
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": user.Id,
			"PASSWORD": user.Password,
		}),
		ClientId: aws.String(c.CognitoAppIdClient),
	}

	result, err := c.CognitoProvider.InitiateAuth(authInput)
	if err != nil {
		c.Log.Error(ErrAWSCognitoAuthUser)
		return "", NewError(http.StatusUnauthorized, ErrAWSCognitoAuthUser, err)
	}

	return *result.AuthenticationResult.AccessToken, nil
}

func (c *Cognito) GetUserByToken(token string) (*string, *Error) {
	input := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(token),
	}

	result, err := c.CognitoProvider.GetUser(input)
	if err != nil {
		c.Log.Error(ErrAWSCognitoGetUserByToken)
		return nil, NewError(http.StatusBadRequest, ErrAWSCognitoGetUserByToken, err)
	}

	return result.Username, nil
}

func (c *Cognito) DeleteUser(token string) *Error {
	input := &cognitoidentityprovider.DeleteUserInput{
		AccessToken: aws.String(token),
	}

	_, err := c.CognitoProvider.DeleteUser(input)
	if err != nil {
		c.Log.Error(ErrAWSCognitoDeleteUser)
		return NewError(http.StatusBadRequest, ErrAWSCognitoDeleteUser, err)
	}

	return nil
}
