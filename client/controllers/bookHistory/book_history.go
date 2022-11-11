package controllers

import (
	"context"

	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
)

type BaseBookHistoryHandler struct {
	grpcClient pb.BookHistoryClient
}

func NewBaseBookHistoryHandler(grpcClient pb.BookHistoryClient) *BaseBookHistoryHandler {
	return &BaseBookHistoryHandler{grpcClient: grpcClient}
}

func (bhh *BaseBookHistoryHandler) GetOneBookHistory(id uint64) (*pb.GetOneHistoryResponse, error) {
	return bhh.grpcClient.GetOne(context.Background(), &pb.GetOneHistoryRequest{Id: id})
}
func (bhh *BaseBookHistoryHandler) InsertBookHistory(bh *pb.BookHistoryItem) (*pb.InsertHistoryResponse, error) {
	return bhh.grpcClient.Insert(context.Background(), &pb.InsertHistoryRequest{BookHistory: bh})
}
func (bhh *BaseBookHistoryHandler) DeleteBookHistory(id uint64) (*pb.DeleteHistoryResponse, error) {
	return bhh.grpcClient.Delete(context.Background(), &pb.DeleteHistoryRequest{Id: id})
}
