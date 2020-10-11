package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"uala.com/uala-api-users/client"
)

type Input struct {
	 MovieName string `json:"MovieName"`
	 MovieYear string `json:"MovieYear"`
}


func HandleRequest(ctx context.Context, input Input) (string, error)  {

	movieName := input.MovieName
	movieYear := input.MovieYear

	fmt.Println("------------------- looking for --------------")

	fmt.Println(movieName)
	fmt.Println(movieYear)

	user, err := client.GetUser(movieName, movieYear)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("------------------- Found user --------------")
	fmt.Println(" Title:   ", user.Title)
	fmt.Println(" Plot:    ", user.Plot)
	fmt.Println(" Raiting: ", user.Rating)
	fmt.Println(" Year:    ", user.Year)
	return "",nil
}

func main() {
	lambda.Start(HandleRequest)
}


