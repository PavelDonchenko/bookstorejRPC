package service

import (
	pb "github.com/PavelDonchenko/bookstorejRPC/server/gen/proto"
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
)

type UService interface {
	GetOne(id uint32) (*pb.GetUserResponse, error)
	GetAll() (*pb.GetAllAUserResponse, error)
	Create(u model.User) (*pb.CreateUserResponse, error)
	Update(u model.User) (*pb.UpdateUserResponse, error)
	Delete(id uint32) (*pb.DeleteUserResponse, error)
}

type BService interface {
	GetOne(id uint32) (*pb.GetBookResponse, error)
	GetAll() (*pb.GetAllBooksResponse, error)
	Create(u model.Book) (*pb.CreateBookResponse, error)
	Update(u model.Book) (*pb.UpdateBookResponse, error)
	Delete(id uint32) (*pb.DeleteBookResponse, error)
}
