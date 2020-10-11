package handler

import (
	"context"
	"fmt"
)


func HandleRequest(ctx context.Context, input Input) {


	fmt.Println("------------------- Input --------------")
	fmt.Println(" Title:   ", input.Title)
	fmt.Println(" Plot:    ", input.Plot)
	fmt.Println(" Raiting: ", input.Rating)
	fmt.Println(" Year:    ", input.Year)
}
