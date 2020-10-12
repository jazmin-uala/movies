package handler

import (
	"context"
	"errors"
	"fmt"
	"jaz.com/uala-api-movies/utils/client"
)

func HandleRequest(ctx context.Context, input Input)   (Output, error) {
	client, err := client.NewClient()
	if err != nil {
		errors.New("Error at creating Client")
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


func (h Handler) Handle(input Input) (Output, error)  {

	movieName := input.MovieName
	movieYear := input.MovieYear

	fmt.Print("------------------- looking for: ")
	fmt.Print(movieName)
	fmt.Println(movieYear)

	movie, err := h.moviesRepository.GetMovie(movieName, movieYear)
	if err != nil {
		errors.New(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("------------------- Found movie return answer --------------")
	fmt.Println("Title: ", movie.Title)
	fmt.Println("Plot: ", movie.Plot)
	fmt.Println("Year: ", movie.Year)
	fmt.Println("Rating:", movie.Rating)

	return Output{Title: movie.Title, Plot: movie.Plot, Year: fmt.Sprintf("%d", movie.Year), Rating: fmt.Sprintf("%.1f",movie.Rating)}, nil

}
