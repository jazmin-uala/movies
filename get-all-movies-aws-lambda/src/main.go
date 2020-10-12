package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"jaz.com/uala-api-movies/get-all-movies-aws-lambda/src/handler"
)

func main() {
	/*input := handler.Input{}
	input.MinRating = 1.0
	input.Since = 1990
	jsonResult,_:= handler.HandleRequest(nil, input)
	fmt.Print(jsonResult)*/
	lambda.Start(handler.HandleRequest)
}
