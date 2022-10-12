package grpcHandler

import (
	"context"
	"time"

	pb "github.com/PavelDonchenko/bookstorejRPC/server/gen/proto"
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
	service "github.com/PavelDonchenko/bookstorejRPC/server/service/book"
)

type BookHandler struct {
	us *service.BookService
	pb.UnimplementedBookServer
}

func NewBookHandler(bookService *service.BookService) *BookHandler {
	return &BookHandler{us: bookService}
}

func (bh *BookHandler) GetOne(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	return bh.us.GetOne(req.GetId())
}

func (bh *BookHandler) GetAll(ctx context.Context, req *pb.GetAllBooksRequest) (*pb.GetAllBooksResponse, error) {
	return bh.us.GetAll()
}

func (bh *BookHandler) Create(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := model.Book{
		ID:         req.GetBook().GetId(),
		Name:       req.GetBook().GetName(),
		BookAuthor: req.GetBook().GetBookAuthor(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return bh.us.Create(book)
}

func (bh *BookHandler) Update(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	book := model.Book{
		ID:         req.GetBook().GetId(),
		Name:       req.GetBook().GetName(),
		BookAuthor: req.GetBook().GetBookAuthor(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return bh.us.Update(book)
}

func (bh *BookHandler) Delete(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	return bh.us.Delete(req.GetId())
}
