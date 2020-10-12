package repository

import "github.com/aws/aws-sdk-go/service/dynamodb"

const (
	tableName = "Movies"
)

type MoviesRepository struct {
	TableName string
	DynamoClient *dynamodb.DynamoDB
}

func NewMoviesRepository(client *dynamodb.DynamoDB) *MoviesRepository {
	repository := new(MoviesRepository)
	repository.TableName = tableName
	repository.DynamoClient = client
	return repository
}


