package testutils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"jaz.com/uala-api-movies/utils/domain"
	"testing"
)

func Before(test *testing.T,  localClient *dynamodb.DynamoDB, items [] domain.Item) *IntegrationTest{
	it := NewIntegrationTest(test, localClient, "Movies" )
	it.createTable(table)
	insertMovies(it , items )
	return it
}

func After(it *IntegrationTest){
	it.deleteTable()
}

var table = dynamodb.CreateTableInput {
	AttributeDefinitions: []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("Title"),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String("Year"),
			AttributeType: aws.String("N"),
		},
	},
	KeySchema: []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("Year"),
			KeyType:       aws.String("HASH"),
		},
		{
			AttributeName: aws.String("Title"),
			KeyType:       aws.String("RANGE"),
		},
	},
	ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(10),
		WriteCapacityUnits: aws.Int64(10),
	},
	GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
		{
			IndexName:             aws.String("Year-index"),
			KeySchema:             []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("Year"),
					KeyType:       aws.String("HASH"),
				},
			},
			Projection:            &dynamodb.Projection{
				ProjectionType:   aws.String("ALL"),
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(10),
			},
		},
	},
	TableName: aws.String("Movies"),
}

func insertMovies(it *IntegrationTest, items [] domain.Item) {
	for _, item := range items {
		av, err := dynamodbattribute.MarshalMap(item)
		if err != nil {
			fmt.Println("Got error marshalling map:")
			fmt.Println(err.Error())
			it.test.Error(err)
		}

		// Create item in table Movies
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(it.table),
		}

		_, err = it.localClient.PutItem(input)
		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			it.test.Error(err)
		}
	}
}

func GetMovie(it *IntegrationTest, movieName string, movieYear string) domain.Item {
	movie := domain.Item{}
	result, _ := it.localClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Movies"),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(movieYear),
			},
			"Title": {
				S: aws.String(movieName),
			},
		},
	})
	dynamodbattribute.UnmarshalMap(result.Item, &movie)
	return movie
}
