package repository

import (
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
)

type UserRepository interface {
	GetAllUsers(offset int, limit int) ([]model.User, error)
	GetUser(id uint32) (*model.User, error)
	CreateUser(u model.User) (model.User, error)
	UpdateUser(u model.User) (*model.User, error)
	DeleteUser(id uint32) (bool, error)
}

type BookRepository interface {
	GetAllBooks(offset int, limit int) ([]model.Book, error)
	GetBook(id uint32) (*model.Book, error)
	CreateBook(u model.Book) (model.Book, error)
	UpdateBook(u model.Book) (*model.Book, error)
	DeleteBook(id uint32) (bool, error)
}
