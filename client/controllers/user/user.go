package controllers

import (
	"context"

	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
)

type BaseHandler struct {
	grpcClient pb.UserClient
}

func NewBaseHandler(client pb.UserClient) *BaseHandler {
	return &BaseHandler{grpcClient: client}
}

func (bh *BaseHandler) GetAll() (*pb.GetAllAUserResponse, error) {
	return bh.grpcClient.GetAll(context.Background(), &pb.GetAllAUserRequest{})
}
func (bh *BaseHandler) Get(id uint32) (*pb.GetUserResponse, error) {
	return bh.grpcClient.GetOne(context.Background(), &pb.GetUserRequest{Id: id})
}
func (bh *BaseHandler) Create(u *pb.UserItem) (*pb.CreateUserResponse, error) {
	return bh.grpcClient.Create(context.Background(), &pb.CreateUserRequest{User: u})
}
func (bh *BaseHandler) Update(u *pb.UserItem) (*pb.UpdateUserResponse, error) {
	return bh.grpcClient.Update(context.Background(), &pb.UpdateUserRequest{User: u})
}
func (bh *BaseHandler) Delete(id uint32) (*pb.DeleteUserResponse, error) {
	return bh.grpcClient.Delete(context.Background(), &pb.DeleteUserRequest{Id: id})
}
