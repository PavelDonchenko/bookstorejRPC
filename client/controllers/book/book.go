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

func (bh *BaseBookHandler) GetAll() (*pb.GetAllBooksResponse, error) {
	return bh.grpcClient.GetAll(context.Background(), &pb.GetAllBooksRequest{})
}
func (bh *BaseBookHandler) Get(id uint32) (*pb.GetBookResponse, error) {
	return bh.grpcClient.GetOne(context.Background(), &pb.GetBookRequest{Id: id})
}
func (bh *BaseBookHandler) Create(u *pb.BookItem) (*pb.CreateBookResponse, error) {
	return bh.grpcClient.Create(context.Background(), &pb.CreateBookRequest{Book: u})
}
func (bh *BaseBookHandler) Update(u *pb.BookItem) (*pb.UpdateBookResponse, error) {
	return bh.grpcClient.Update(context.Background(), &pb.UpdateBookRequest{Book: u})
}
func (bh *BaseBookHandler) Delete(id uint32) (*pb.DeleteBookResponse, error) {
	return bh.grpcClient.Delete(context.Background(), &pb.DeleteBookRequest{Id: id})
}
