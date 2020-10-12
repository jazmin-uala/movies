package test

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"io/ioutil"
	"jaz.com/uala-api-movies/update-movie-aws-lambda/src/handler"
	"jaz.com/uala-api-movies/utils/client"
	"jaz.com/uala-api-movies/utils/domain"
	"jaz.com/uala-api-movies/utils/testutils"
	"os"
	"testing"
)

var localClient *dynamodb.DynamoDB

func TestMain(m *testing.M) {
	container, _ := testutils.CreateDynamoContainer()
	defer container.Pool.Purge(container.Resource)

	dynamoDockerConfig := aws.Config{
		Endpoint:    aws.String(container.Endpoint),
		Credentials: credentials.NewStaticCredentials("local", "local", "local"),
	}
	localClient, _ = client.NewClientWithConfig(dynamoDockerConfig)

	code := m.Run()
	os.Exit(code)
}





func TestUpdateMovieRecord(test *testing.T){
	it := testutils.Before(test, localClient, getMoviesFromFile())
	movieBefore :=  testutils.GetMovie(it,"La sirenita","1994")
	fmt.Printf("movie.Rating before testutils = %f; want 5.0 \n", movieBefore.Rating)

	//Given
	input := handler.Input{}
	input.Title = "La sirenita"
	input.Year = "1994"
	input.Rating = "8.0"
	moviesRepository := handler.NewRepository(localClient)
	updateHandler := handler.NewHandler(moviesRepository)

	//When
	updateHandler.Handle(input)

	//Then
	movie:= testutils.GetMovie(it,"La sirenita","1994")
	if movie.Rating != 8.0{
		test.Errorf("movie.Rating = %f; want 8.0", movie.Rating)
	}
	testutils.After(it)
}

func getMoviesFromFile() [] domain.Item {
	raw, err := ioutil.ReadFile("./data/movies.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var items [] domain.Item
	json.Unmarshal(raw, &items)
	return items
}






