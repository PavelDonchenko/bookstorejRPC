package repository

import model "github.com/PavelDonchenko/bookstoreCRUD/server/models"

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetOne(id uint32) (model.User, error)
	Create(a model.User) (model.User, error)
	Update(id uint32) (model.User, error)
	Delete(id uint32) (bool, error)
}
