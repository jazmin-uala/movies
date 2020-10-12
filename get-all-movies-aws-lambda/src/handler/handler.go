package handler

import (
	"context"
	"fmt"
	"jaz.com/uala-api-movies/src/client"
)


func HandleRequest(ctx context.Context, input Input) {
	fmt.Println("------------------- all movies --------------")
	client, _ := client.NewLocalClient()
	moviesRepository := NewRepository(client)
	moviesRepository.GetAllMoviesSinceYearWithRating(input.Since, input.MinRating)
}
