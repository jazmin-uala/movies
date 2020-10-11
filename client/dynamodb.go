package client

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"uala.com/uala-api-users/domain"
)

const (
	DefaultRegion = "us-east-1"
	tableName = "Movies"
)

func GetUser(movieName string, movieYear string) (domain.Item, error) {
	user := domain.Item{}

	svc, err := getClient()

	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(movieYear),
			},
			"Title": {
				S: aws.String(movieName),
			},
		},
	})


	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}

	if result.Item == nil {
		msg := "Could not find '" + movieName + "'"
		return user,errors.New(msg)
	}


	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}


	return user, nil
}


func getClient() (*dynamodb.DynamoDB, error) {


	conf := aws.Config{Region: aws.String(DefaultRegion)}
	session := session.Must(session.NewSession())


	/*session, err := session.NewSessionWithOptions(session.Options{
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
	}*/

	svc := dynamodb.New(session,  &conf)
	return svc, nil

}



