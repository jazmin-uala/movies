package handler

type Movie struct {
	Title string `json:"Title"`
	Year string `json:"Year"`
	Plot string `json:"Plot"`
	Rating string `json:"Rating"`

}
type Output struct {
	Movies []Movie `json:"Movies"`
}
