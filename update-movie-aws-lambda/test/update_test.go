package test

import (
	"fmt"
	"jaz.com/uala-api-movies/src/client"
	"jaz.com/uala-api-movies/update-movie-aws-lambda/src/handler"
	"testing"
)

func TestUpdateMovieRecord(t *testing.T){
	input := handler.Input{}
	input.Title = "La sirenita"
	input.Year = "1994"
	input.Rating = "8.0"
	input.Plot = "Una joven sirena se enamora de un humano"

	client, err := client.NewLocalClient()
	if err != nil {
		fmt.Println("Client Error: ", err)
		return
	}
	moviesRepository := handler.NewRepository(client)
	updateHandler := handler.NewHandler(moviesRepository)
	updateHandler.Handle(input)

}
