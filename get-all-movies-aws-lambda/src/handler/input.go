package handler

type Input struct {
	Since int `json:"since"`
	MinRating float32 `json:"min_rating"`
}

