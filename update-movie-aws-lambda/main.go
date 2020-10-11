package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"jaz.com/uala-api-movies/update-movie-aws-lambda/handler"
)

func main() {
	/*input := handler.Input{}
	input.Title = "La sirenita"
	input.Year = "1994"
	input.Rating = "9.4"
	input.Plot = "Una joven sirena se enamora de un humano"
	handler.HandleRequest(nil, input)*/
	lambda.Start(handler.HandleRequest)
}
