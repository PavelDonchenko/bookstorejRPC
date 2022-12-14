package service

import (
	pb "github.com/PavelDonchenko/bookstorejRPC/server/gen/proto"
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
)

type UserService interface {
	GetUser(id uint32) (*pb.GetUserResponse, error)
	GetAllUsers(page uint32) (*pb.GetAllUsersResponse, error)
	CreateUser(u model.User) (*pb.CreateUserResponse, error)
	UpdateUser(u model.User) (*pb.UpdateUserResponse, error)
	DeleteUser(id uint32) (*pb.DeleteUserResponse, error)
}

type BookService interface {
	GetBook(id uint32) (*pb.GetBookResponse, error)
	GetAllBooks(page uint32) (*pb.GetAllBooksResponse, error)
	CreateBook(u model.Book) (*pb.CreateBookResponse, error)
	UpdateBook(u model.Book) (*pb.UpdateBookResponse, error)
	DeleteBook(id uint32) (*pb.DeleteBookResponse, error)
}
