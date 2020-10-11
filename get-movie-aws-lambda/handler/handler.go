package handler

import (
	"context"
	"fmt"
	"jaz.com/uala-api-movies/client"
)


func HandleRequest(ctx context.Context, input Input) {

	movieName := input.MovieName
	movieYear := input.MovieYear

	fmt.Println("------------------- looking for --------------")

	fmt.Println(movieName)
	fmt.Println(movieYear)

	movie, err := client.GetMovie(movieName, movieYear)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("------------------- Found movie --------------")
	fmt.Println(" Title:   ", movie.Title)
	fmt.Println(" Plot:    ", movie.Plot)
	fmt.Println(" Raiting: ", movie.Rating)
	fmt.Println(" Year:    ", movie.Year)
}
