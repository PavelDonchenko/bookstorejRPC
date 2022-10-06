package service

import (
	pb "github.com/PavelDonchenko/bookstorejRPC/gen/proto"
	"github.com/PavelDonchenko/bookstorejRPC/models"
)

type UService interface {
	GetOne(id uint32) (*pb.GetUserResponse, error)
	GetAll() (*pb.GetAllAUserResponse, error)
	Create(a model.User) (*pb.CreateUserResponse, error)
	Update(a model.User) (*pb.UpdateUserResponse, error)
	Delete(id uint32) (*pb.DeleteUserResponse, error)
}
