package client

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"jaz.com/uala-api-movies/domain"
	"os"
	"strconv"
)

const (
	DefaultRegion = "us-east-1"
	tableName = "Movies"
)

var (
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
)

func GetMovie(movieName string, movieYear string) (domain.Item, error) {
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


func GetAllMoviesSinceYearWithRating(year int, minRating  float32) () {
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
		TableName:                 aws.String(tableName),
	}

	svc, err := getClient()

	fmt.Println("Make the DynamoDB Query API call")
	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
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

func UpdateMovieRating(movie domain.Item){
	// Update item in table Movies


	year:= strconv.Itoa(movie.Year)
	fmt.Println("updating '" + movie.Title + "' (" + year+ ")")

	newRating := fmt.Sprintf("%.1f",movie.Rating)
	fmt.Println("new rating: ",newRating)

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
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

	client, err := getClient()
	_, err = client.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated '" + movie.Title + "' (" + year + ") rating to " + fmt.Sprintf("%.1f",movie.Rating))
}


func getClient() (*dynamodb.DynamoDB, error) {

	conf := aws.Config{Region: aws.String(DefaultRegion)}
	session, err := getSession()
	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(session,  &conf)
	return svc, nil

}

func getSession() (*session.Session, error) {

	if ENVIRONMENT == "dev" {
		fmt.Println("running in dev environment")
		session := session.Must(session.NewSession())
		return session, nil
	}

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



