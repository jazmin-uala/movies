package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"jaz.com/uala-api-movies/src/domain"
	"jaz.com/uala-api-movies/src/repository"
	"os"
)

type Repository struct {
	movies *repository.MoviesRepository
}

func NewRepository(client *dynamodb.DynamoDB) * Repository {
	this := new(Repository)
	this.movies = repository.NewMoviesRepository(client)
	return this
}

func (repository Repository)  GetAllMoviesSinceYearWithRating(year int, minRating  float32) () {
	// Create the Expression to fill the input struct with.
	// Get all movies in that year; we'll pull out those with a higher rating later
	filter := expression.Name("Year").GreaterThan(expression.Value(2000))

	// Get back the title, year, and rating
	proj := expression.NamesList(expression.Name("Title"), expression.Name("Year"), expression.Name("Rating"))

	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(repository.movies.TableName),
	}


	fmt.Println("Make the DynamoDB Query API call")
	// Make the DynamoDB Query API call
	result, err := repository.movies.DynamoClient.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}


	numItems := 0
	for _, i := range result.Items {
		item := domain.Item{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if item.Rating > minRating {
			numItems++
			fmt.Println("Title: ", item.Title)
			fmt.Println("Rating:", item.Rating)
			fmt.Println()
		}
	}

	fmt.Println("Found", numItems, "movie(s) with a rating above", minRating, "since", year)
}

