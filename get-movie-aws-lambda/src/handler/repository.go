package handler

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"jaz.com/uala-api-movies/utils/domain"
	"jaz.com/uala-api-movies/utils/repository"
)

type Repository struct {
	movies *repository.MoviesRepository
}

func NewRepository(client *dynamodb.DynamoDB) * Repository {
	this := new(Repository)
	this.movies = repository.NewMoviesRepository(client)
	return this
}

func (repository Repository) GetMovie(movieName string, movieYear string) (domain.Item, error) {
	movie := domain.Item{}

	result, err := repository.movies.DynamoClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repository.movies.TableName),
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
		return movie, err
	}

	if result.Item == nil {
		msg := "Could not find '" + movieName + "'"
		return movie,errors.New(msg)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &movie)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return movie, nil
}