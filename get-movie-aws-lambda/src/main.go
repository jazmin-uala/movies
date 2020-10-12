package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"jaz.com/uala-api-movies/get-movie-aws-lambda/src/handler"
)

func main() {
	/*input := handler.Input{}
	input.MovieName = "The Big New Movie"
	input.MovieYear = "2016"
	x,_:= handler.HandleRequest(nil, input)
	fmt.Print(x)*/
	lambda.Start(handler.HandleRequest)
}


