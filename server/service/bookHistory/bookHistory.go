package service

import (
	"context"
	"time"

	pb "github.com/PavelDonchenko/bookstorejRPC/server/gen/proto"
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
	repository2 "github.com/PavelDonchenko/bookstorejRPC/server/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookHistoryService struct {
	bookHistoryRepo repository2.BookHistoryRepository
}

func NewBookHistoryService(bookHistoryRepo repository2.BookHistoryRepository) *BookHistoryService {
	return &BookHistoryService{
		bookHistoryRepo: bookHistoryRepo}
}

func (bhs *BookHistoryService) GetOneBookHistory(ctx context.Context, id uint64) (*pb.GetOneHistoryResponse, error) {
	bh, err := bhs.bookHistoryRepo.GetOneBookHistory(ctx, id)
	if err != nil {
		return nil, err
	}

	doc := pb.BookHistoryItem{
		Id:         bh.Id,
		BookAuthor: bh.BookAuthor,
		BookName:   bh.BookName,
		LastUser:   bh.LastUser,
		BookTaken:  timestamppb.New(bh.BookTaken),
		BookReturn: timestamppb.New(bh.BookReturn),
	}

	return &pb.GetOneHistoryResponse{BookHistory: &doc}, nil
}
func (bhs *BookHistoryService) InsertBookHistory(ctx context.Context, bh model.BookHistory) (*pb.InsertHistoryResponse, error) {
	bt := time.Now().UTC()
	br := time.Now().UTC().AddDate(1, 2, 3)

	doc := model.BookHistory{
		Id:         bh.Id,
		BookAuthor: bh.BookAuthor,
		BookName:   bh.BookName,
		LastUser:   bh.LastUser,
		BookTaken:  bt,
		BookReturn: br,
	}

	if err := bhs.bookHistoryRepo.InsertBookHistory(ctx, doc); err != nil {
		return nil, err
	}

	return &pb.InsertHistoryResponse{Success: true}, nil
}
func (bhs *BookHistoryService) DeleteBookHistory(ctx context.Context, id uint64) (*pb.DeleteHistoryResponse, error) {
	result, err := bhs.bookHistoryRepo.DeleteBookHistory(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteHistoryResponse{Success: result}, nil
}
