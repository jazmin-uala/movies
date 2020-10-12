package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"jaz.com/uala-api-movies/update-movie-aws-lambda/src/handler"
)

func main() {
	/*input := handler.Input{}
	input.Title = "The e"
	input.Year = "2016"
	input.Rating = "7.0"
	handler.HandleRequest(nil, input)*/
	lambda.Start(handler.HandleRequest)
}
