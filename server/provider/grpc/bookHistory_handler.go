package grpcHandler

import (
	"context"
	"time"

	pb "github.com/PavelDonchenko/bookstorejRPC/server/gen/proto"
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
	service "github.com/PavelDonchenko/bookstorejRPC/server/service/bookHistory"
)

type BookHistoryHandler struct {
	bhs *service.BookHistoryService
	pb.UnimplementedBookHistoryServer
}

func NewBookHistoryHandler(bhs *service.BookHistoryService) *BookHistoryHandler {
	return &BookHistoryHandler{bhs: bhs}
}

func (bhh *BookHistoryHandler) Insert(ctx context.Context, req *pb.InsertHistoryRequest) (*pb.InsertHistoryResponse, error) {
	bh := model.BookHistory{
		Id:         req.GetBookHistory().GetId(),
		BookAuthor: req.GetBookHistory().GetBookAuthor(),
		BookName:   req.GetBookHistory().GetBookName(),
		LastUser:   req.GetBookHistory().GetLastUser(),
		BookTaken:  time.Now().UTC(),
		BookReturn: time.Now().UTC().AddDate(1, 2, 3),
	}
	return bhh.bhs.InsertBookHistory(ctx, bh)
}
func (bhh *BookHistoryHandler) GetOne(ctx context.Context, req *pb.GetOneHistoryRequest) (*pb.GetOneHistoryResponse, error) {
	return bhh.bhs.GetOneBookHistory(ctx, req.GetId())
}
func (bhh *BookHistoryHandler) Delete(ctx context.Context, req *pb.DeleteHistoryRequest) (*pb.DeleteHistoryResponse, error) {
	return bhh.bhs.DeleteBookHistory(ctx, req.GetId())
}
