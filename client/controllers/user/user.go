package controllers

import (
	"context"

	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
)

type BaseUserHandler struct {
	grpcClient pb.UserClient
}

func NewBaseUserHandler(client pb.UserClient) *BaseUserHandler {
	return &BaseUserHandler{grpcClient: client}
}

func (bh *BaseUserHandler) GetAll() (*pb.GetAllAUserResponse, error) {
	return bh.grpcClient.GetAll(context.Background(), &pb.GetAllAUserRequest{})
}
func (bh *BaseUserHandler) Get(id uint32) (*pb.GetUserResponse, error) {
	return bh.grpcClient.GetOne(context.Background(), &pb.GetUserRequest{Id: id})
}
func (bh *BaseUserHandler) Create(u *pb.UserItem) (*pb.CreateUserResponse, error) {
	return bh.grpcClient.Create(context.Background(), &pb.CreateUserRequest{User: u})
}
func (bh *BaseUserHandler) Update(u *pb.UserItem) (*pb.UpdateUserResponse, error) {
	return bh.grpcClient.Update(context.Background(), &pb.UpdateUserRequest{User: u})
}
func (bh *BaseUserHandler) Delete(id uint32) (*pb.DeleteUserResponse, error) {
	return bh.grpcClient.Delete(context.Background(), &pb.DeleteUserRequest{Id: id})
}
