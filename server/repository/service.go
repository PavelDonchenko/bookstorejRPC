package repository

import (
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetOne(id uint32) (*model.User, error)
	Create(u model.User) (model.User, error)
	Update(u model.User) (model.User, error)
	Delete(id uint32) (bool, error)
}

type BookRepository interface {
	GetAll() ([]model.Book, error)
	GetOne(id uint32) (model.Book, error)
	Create(u model.Book) (model.Book, error)
	Update(u model.Book) (model.Book, error)
	Delete(id uint32) (bool, error)
}
