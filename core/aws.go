package core

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
    "github.com/spf13/viper"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
    "strings"
)

type AWS struct {
	Cognito *CognitoService
	Log *Log
}

type CognitoService struct {
	CognitoProvider *cognitoidentityprovider.CognitoIdentityProvider
	CognitoAppIdClient string
	Log *Log
}

func NewAWS(log *Log) *AWS {
	return &AWS{
		Log: log,
	}
}

func (a *AWS) NewCognito() (*CognitoService, *Error) {
	config := &aws.Config{
		Region: aws.String(viper.GetString("AWS_REGION")),
	}
	sess, err := session.NewSession(config)
	if err != nil {
		return nil, NewError(500, "Error create new session to AWS Cognito", err)
	}

	client := cognitoidentityprovider.New(sess)

	return &CognitoService{
		CognitoProvider: client,
		CognitoAppIdClient: viper.GetString("APP_CLIENT_ID"),
		Log: a.Log,
	}, nil
}

func (cs *CognitoService) SignUp(user model.User) *Error {
	userCognito := &cognitoidentityprovider.SignUpInput{
		ClientId:aws.String(cs.CognitoAppIdClient),
		Username: aws.String(user.Id),
		Password: aws.String(strings.TrimSpace(user.Password)),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name: aws.String("name"),
				Value: aws.String(user.UserName),
			},
			{
				Name: aws.String("email"),
				Value: aws.String(user.Email),
			},
		},
	}

	_, err := cs.CognitoProvider.SignUp(userCognito)
	if err != nil {
		cs.Log.Error("Erro to create user in AWS Cognito")
		return NewError(500, "Error to create user in AWS Cognito", err)
	}

	return nil
}

func (cs *CognitoService) SignIn(user model.User) (string, *Error) {
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": user.Id,
			"PASSWORD": user.Password,
		}),
		ClientId: aws.String(cs.CognitoAppIdClient),
	}

	result, err := cs.CognitoProvider.InitiateAuth(authInput)
	if err != nil {
		cs.Log.Error("Failed to auth user to AWS Cognito")
		return "", NewError(500, "Failed to auth user to AWS Cognito", err)
	}

	return *result.AuthenticationResult.AccessToken, nil
}

func (cs *CognitoService) GetUserByToken(token string) (*cognitoidentityprovider.GetUserOutput, *Error) {
	input := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(token),
	}

	result, err := cs.CognitoProvider.GetUser(input)
	if err != nil {
		cs.Log.Error("Failed to recover the User by his token")
		return nil, NewError(500, "Failed to recover the User by his token", err)
	}

	return result, nil
}

func (cs *CognitoService) DeleteUser(token string) *Error {
	input := &cognitoidentityprovider.DeleteUserInput{
		AccessToken: aws.String(token),
	}

	_, err := cs.CognitoProvider.DeleteUser(input)
	if err != nil {
		cs.Log.Error("Failed to delete User")
		return NewError(500, "Failed to delete User", err)
	}

	return nil
}