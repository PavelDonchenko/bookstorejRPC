package model

import "time"

type BookHistory struct {
	Id         uint64    `json:"id"`
	BookAuthor string    `json:"book_author"`
	BookName   string    `json:"book_name"`
	LastUser   string    `json:"last_user"`
	BookTaken  time.Time `json:"bookTaken"`
	BookReturn time.Time `json:"bookReturn"`
}
