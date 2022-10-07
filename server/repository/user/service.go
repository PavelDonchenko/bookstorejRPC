package repository

import (
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetOne(id uint32) (model.User, error)
	Create(u model.User) (model.User, error)
	Update(u model.User) (model.User, error)
	Delete(id uint32) (bool, error)
}