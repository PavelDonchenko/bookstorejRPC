package service

import (
	pb "github.com/PavelDonchenko/bookstoreCRUD/gen/proto"
	model "github.com/PavelDonchenko/bookstoreCRUD/server/models"
)

type UService interface {
	GetOne(id uint32) (*pb.GetUserResponse, error)
	GetAll() (*pb.GetAllAUserResponse, error)
	Create(a model.User) (*pb.CreateUserResponse, error)
	Update(id uint32) (*pb.UpdateUserResponse, error)
	Delete(id uint32) (*pb.DeleteUserResponse, error)
}
