package handler

import (
	"context"
	"fmt"
	"jaz.com/uala-api-movies/client"
)


func HandleRequest(ctx context.Context, input Input) {
	fmt.Println("------------------- all movies --------------")
	client.GetAllMoviesSinceYearWithRating(input.Since, input.MinRating)
}
