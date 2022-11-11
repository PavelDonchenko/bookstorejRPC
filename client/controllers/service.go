package controllers

import (
	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
)

type UserApiController interface {
	GetAllUsers(page uint32) (*pb.GetAllUsersResponse, error)
	GetUser(id uint32) (*pb.GetUserResponse, error)
	CreateUser(u *pb.UserItem) (*pb.CreateUserResponse, error)
	UpdateUser(u *pb.UserItem) (*pb.UpdateUserResponse, error)
	DeleteUser(id uint32) (*pb.DeleteUserResponse, error)
}

type BookApiController interface {
	GetAllBooks(page uint32) (*pb.GetAllBooksResponse, error)
	GetBook(id uint32) (*pb.GetBookResponse, error)
	CreateBook(u *pb.BookItem) (*pb.CreateBookResponse, error)
	UpdateBook(u *pb.BookItem) (*pb.UpdateBookResponse, error)
	DeleteBook(id uint32) (*pb.DeleteBookResponse, error)
}

type BookHistoryApiControllers interface {
	GetOneBookHistory(id uint64) (*pb.GetOneHistoryResponse, error)
	InsertBookHistory(bh *pb.BookHistoryItem) (*pb.InsertHistoryResponse, error)
	DeleteBookHistory(id uint64) (*pb.DeleteHistoryResponse, error)
}
