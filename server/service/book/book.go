package service

import (
	pb "github.com/PavelDonchenko/bookstorejRPC/server/gen/proto"
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
	repository2 "github.com/PavelDonchenko/bookstorejRPC/server/repository"
	repository "github.com/PavelDonchenko/bookstorejRPC/server/repository/book"
	"github.com/PavelDonchenko/bookstorejRPC/server/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookService struct {
	bookRepo repository2.BookRepository
}

func NewBookService(bookRepo *repository.BookRepo) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

func (us *BookService) GetBook(id uint32) (*pb.GetBookResponse, error) {
	book, err := us.bookRepo.GetBook(id)

	item := pb.BookItem{
		Id:         uint32(book.ID),
		Name:       book.Name,
		BookAuthor: book.BookAuthor,
		CrateAt:    timestamppb.New(book.CreatedAt),
		UpdateAt:   timestamppb.New(book.UpdatedAt),
	}
	return &pb.GetBookResponse{Book: &item}, err
}

func (us *BookService) GetAllBooks(page uint32) (*pb.GetAllBooksResponse, error) {

	offset, limit := utils.Pagination(page)

	books, err := us.bookRepo.GetAllBooks(offset, limit)

	items := []*pb.BookItem{}

	for _, v := range books {
		u := &pb.BookItem{
			Id:         uint32(v.ID),
			Name:       v.Name,
			BookAuthor: v.BookAuthor,
			CrateAt:    timestamppb.New(v.CreatedAt),
			UpdateAt:   timestamppb.New(v.UpdatedAt),
		}
		items = append(items, u)
	}

	return &pb.GetAllBooksResponse{Book: items}, err
}

func (us *BookService) CreateBook(u model.Book) (*pb.CreateBookResponse, error) {
	book, err := us.bookRepo.CreateBook(u)

	item := pb.BookItem{
		Id:         uint32(book.ID),
		Name:       book.Name,
		BookAuthor: book.BookAuthor,
		CrateAt:    timestamppb.New(book.CreatedAt),
		UpdateAt:   timestamppb.New(book.UpdatedAt),
	}

	return &pb.CreateBookResponse{Book: &item}, err
}

func (us *BookService) UpdateBook(u model.Book) (*pb.UpdateBookResponse, error) {
	book, err := us.bookRepo.UpdateBook(u)

	item := pb.BookItem{
		Id:         uint32(book.ID),
		Name:       book.Name,
		BookAuthor: book.BookAuthor,
		CrateAt:    timestamppb.New(book.CreatedAt),
		UpdateAt:   timestamppb.New(book.UpdatedAt),
	}

	return &pb.UpdateBookResponse{Book: &item}, err
}

func (us *BookService) DeleteBook(id uint32) (*pb.DeleteBookResponse, error) {
	result, err := us.bookRepo.DeleteBook(id)

	return &pb.DeleteBookResponse{Success: result}, err
}
