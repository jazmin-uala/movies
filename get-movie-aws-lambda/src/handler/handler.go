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

func (h Handler) Handle(input Input){

	movieName := input.MovieName
	movieYear := input.MovieYear

	fmt.Println("------------------- looking for --------------")

	fmt.Println(movieName)
	fmt.Println(movieYear)

	client, _ := client.NewClient()
	moviesRepository := NewRepository(client)


	movie, err := moviesRepository.GetMovie(movieName, movieYear)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("------------------- Found movie --------------")
	fmt.Println(" Title:   ", movie.Title)
	fmt.Println(" Plot:    ", movie.Plot)
	fmt.Println(" Raiting: ", movie.Rating)
	fmt.Println(" Year:    ", movie.Year)
}
