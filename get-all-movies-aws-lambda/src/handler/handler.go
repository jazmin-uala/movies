package handler

import (
	"context"
	"errors"
	"fmt"
	"jaz.com/uala-api-movies/utils/client"
)

func HandleRequest(ctx context.Context, input Input)  (Output, error){
	client, err := client.NewClient()
	if err != nil {
		errors.New(fmt.Sprintf("ClientError, %v", err))
	}
	moviesRepository := NewRepository(client)
	handler := NewHandler(moviesRepository)
	return handler.Handle(input)
}

type Handler struct{
	moviesRepository Repository
}

func NewHandler(repository * Repository) * Handler{
	this := new(Handler)
	this.moviesRepository = *repository
	return this
}


func (h Handler) Handle(input Input)  (Output, error) {
	fmt.Println("------------------- all movies --------------")
	result, _ := h.moviesRepository.GetAllMoviesSinceYearWithRating(input.Since, input.MinRating)
	movies := make([]Movie, 0)
	for _, item := range result {
		movies = append(movies, Movie{ Title: item.Title, Year: fmt.Sprintf("%d", item.Year), Plot: item.Plot, Rating: fmt.Sprintf("%.1f",item.Rating)})
	}
	return Output{Movies: movies},nil
}
