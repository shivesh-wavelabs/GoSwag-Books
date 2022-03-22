package models

// swagger:model Book
type Book struct {
	// ID of the book
	// in: int
	ID int `json:"id" example:"1"`
	// Title of the book
	// in: string
	Title string `json:"title" example:"Staggering Forward"`
	// Author of the book
	// in: string
	Author string `json:"author" example:"Mr. Karnad"`
	// Year of the book
	// in: string
	Year string `json:"year" example:"2010"`
}
