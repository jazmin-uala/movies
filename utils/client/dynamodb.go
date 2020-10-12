package client

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

const (
	DefaultRegion = "us-east-1"
)

var (
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
)

type Client struct {};

func NewLocalClient() (*dynamodb.DynamoDB, error) {
	conf := aws.Config{Region: aws.String(DefaultRegion)}
	session, err := getLocalSession()
	if err != nil {
		return nil, err
	}
	svc := dynamodb.New(session,  &conf)
	return svc, nil
}

func NewClientWithConfig(config aws.Config) (*dynamodb.DynamoDB, error) {
	session, err := getLocalSession()
	if err != nil {
		return nil, err
	}
	svc := dynamodb.New(session,&config)
	return svc, nil
}

func NewClient() (*dynamodb.DynamoDB, error) {
	conf := aws.Config{Region: aws.String(DefaultRegion)}
	session, err := getLambdaSession()
	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(session,  &conf)
	return svc, nil
}

func getLambdaSession() (*session.Session, error) {
	fmt.Println("running in dev environment")
	session := session.Must(session.NewSession())
	return session, nil
}

func  getLocalSession() (*session.Session, error) {
	fmt.Println("running in local environment")
	session, err := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Profile: "uala-arg-playground-dev",

		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String(DefaultRegion),
		},

		// Force enable Shared Config support
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}



