package domain

import "gorm.io/gorm"

type Review struct {
	BookId      uint
	UserId      uint
	Rating      uint
	ReviewTitle string
	ReviewText  string
	gorm.Model
}

type ReviewResponse struct {
	BookAuthor   string `json:"book_author"`
	BookTitle    string `json:"book_title"`
	ReviewAuthor string `json:"review_author"`
	Rating       uint   `json:"rating"`
	ReviewTitle  string `json:"review_title"`
	ReviewText   string `json:"review_text"`
}
