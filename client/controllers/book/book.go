package controllers

import (
	"context"

	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
)

type BaseBookHandler struct {
	grpcClient pb.BookClient
}

func NewBaseBookHandler(client pb.BookClient) *BaseBookHandler {
	return &BaseBookHandler{grpcClient: client}
}

func (bh *BaseBookHandler) GetAllBooks(page uint32) (*pb.GetAllBooksResponse, error) {
	return bh.grpcClient.GetAllBooks(context.Background(), &pb.GetAllBooksRequest{Page: page})
}

func (bh *BaseBookHandler) GetBook(id uint32) (*pb.GetBookResponse, error) {
	return bh.grpcClient.GetBook(context.Background(), &pb.GetBookRequest{Id: id})
}

func (bh *BaseBookHandler) CreateBook(u *pb.BookItem) (*pb.CreateBookResponse, error) {
	return bh.grpcClient.CreateBook(context.Background(), &pb.CreateBookRequest{Book: u})
}

func (bh *BaseBookHandler) UpdateBook(u *pb.BookItem) (*pb.UpdateBookResponse, error) {
	return bh.grpcClient.UpdateBook(context.Background(), &pb.UpdateBookRequest{Book: u})
}

func (bh *BaseBookHandler) DeleteBook(id uint32) (*pb.DeleteBookResponse, error) {
	return bh.grpcClient.DeleteBook(context.Background(), &pb.DeleteBookRequest{Id: id})
}
