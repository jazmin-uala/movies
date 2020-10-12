package handler

import (
	"context"
	"errors"
	"fmt"
	"jaz.com/uala-api-movies/utils/client"
	"jaz.com/uala-api-movies/utils/domain"
	"strconv"
)

func HandleRequest(ctx context.Context, input Input)  {
	client, err := client.NewLocalClient()
	if err != nil {
		fmt.Println( err)
		errors.New("Client Error")
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
	year, err := strconv.Atoi(input.Year)
	if err != nil {
		errors.New(fmt.Sprintf("Error in year: ", year))
	}
	rating, err := strconv.ParseFloat(input.Rating, 32)
	if err != nil {
		errors.New(fmt.Sprintf("Error in ratio: ",rating))
	}
	movie:= domain.Item{}
	movie.Rating = float32(rating)
	movie.Year = year
	movie.Title = input.Title
	err = h.moviesRepository.updateMovieRating(movie)
	if err != nil {
		errors.New(fmt.Sprintf("Error updating: ",err))
	}
}



