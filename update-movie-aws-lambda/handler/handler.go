package handler

import (
	"context"
	"fmt"
	"jaz.com/uala-api-movies/client"
	"jaz.com/uala-api-movies/domain"
	"strconv"
)


func HandleRequest(ctx context.Context, input Input) {
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

	client.UpdateMovieRating(movie)
}
