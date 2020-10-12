package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"jaz.com/uala-api-movies/update-movie-aws-lambda/src/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
