package handler

import (
	"context"
	"fmt"
	"jaz.com/uala-api-movies/utils/client"
)

func HandleRequest(ctx context.Context, input Input) {
	client, err := client.NewClient()
	if err != nil {
		fmt.Println("Client Error: ", err)
		return
	}
	moviesRepository := NewRepository(client)
	handler := NewHandler(moviesRepository)
	handler.Handle(input)
}

type Handler struct{
	moviesRepository Repository
}

func NewHandler(repository * Repository) * Handler{
	this := new(Handler)
	this.moviesRepository = *repository
	return this
}


func (h Handler) Handle(input Input) {
	fmt.Println("------------------- all movies --------------")
	client, _ := client.NewClient()
	moviesRepository := NewRepository(client)
	moviesRepository.GetAllMoviesSinceYearWithRating(input.Since, input.MinRating)
}
