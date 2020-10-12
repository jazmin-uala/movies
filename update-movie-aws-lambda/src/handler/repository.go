package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"jaz.com/uala-api-movies/utils/domain"
	"jaz.com/uala-api-movies/utils/repository"
	"strconv"
)

type Repository struct {
	movies *repository.MoviesRepository
}

func NewRepository(client *dynamodb.DynamoDB) * Repository {
	this := new(Repository)
	this.movies = repository.NewMoviesRepository(client)
	return this
}

func (repository Repository) updateMovieRating(movie domain.Item) error{
	updatedMovie := domain.Item{}

	year:= strconv.Itoa(movie.Year)
	fmt.Println("updating '" + movie.Title + "' (" + year+ ")")

	newRating := fmt.Sprintf("%.1f",movie.Rating)
	fmt.Println("new rating: ",newRating)

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(repository.movies.TableName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String(newRating),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(year),
			},
			"Title": {
				S: aws.String(movie.Title),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rating = :r"),
	}

	result, err := repository.movies.DynamoClient.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Successfully updated '" + movie.Title + "' (" + year + ") rating to " + fmt.Sprintf("%.1f",movie.Rating))

	err = dynamodbattribute.UnmarshalMap(result.Attributes, &updatedMovie)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return err
}
