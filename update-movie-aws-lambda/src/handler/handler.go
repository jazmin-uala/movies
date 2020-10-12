package handler

import (
	"context"
	"fmt"
	"jaz.com/uala-api-movies/src/client"
	"jaz.com/uala-api-movies/src/domain"
	"strconv"
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

func (h Handler) Handle(input Input){
	year, err := strconv.Atoi(input.Year)
	if err != nil {
		fmt.Print("Error in year: ",year)
		fmt.Println(err)
		return
	}
	rating, err := strconv.ParseFloat(input.Rating, 32)
	if err != nil {
		fmt.Print("Error in ratio: ",rating)
		fmt.Println(err)
		return
	}
	movie:= domain.Item{}
	movie.Plot = input.Plot
	movie.Rating = float32(rating)
	movie.Year = year
	movie.Title = input.Title
	h.moviesRepository.updateMovieRating(movie)
}



