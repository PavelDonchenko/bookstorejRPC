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

func (bh *BookHandler) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	return bh.us.GetBook(req.GetId())
}

func (bh *BookHandler) GetAllBooks(ctx context.Context, req *pb.GetAllBooksRequest) (*pb.GetAllBooksResponse, error) {
	return bh.us.GetAllBooks(req.GetPage())
}

func (bh *BookHandler) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := model.Book{
		ID:         req.GetBook().GetId(),
		Name:       req.GetBook().GetName(),
		BookAuthor: req.GetBook().GetBookAuthor(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return bh.us.CreateBook(book)
}

func (bh *BookHandler) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	book := model.Book{
		ID:         req.GetBook().GetId(),
		Name:       req.GetBook().GetName(),
		BookAuthor: req.GetBook().GetBookAuthor(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return bh.us.UpdateBook(book)
}

func (bh *BookHandler) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	return bh.us.DeleteBook(req.GetId())
}
