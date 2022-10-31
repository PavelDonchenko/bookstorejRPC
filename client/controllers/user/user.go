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

func (bh *BaseUserHandler) GetAllUsers(page uint32) (*pb.GetAllUsersResponse, error) {
	return bh.grpcClient.GetAllUsers(context.Background(), &pb.GetAllUsersRequest{Page: page})
}

func (bh *BaseUserHandler) GetUser(id uint32) (*pb.GetUserResponse, error) {
	return bh.grpcClient.GetUser(context.Background(), &pb.GetUserRequest{Id: id})
}

func (bh *BaseUserHandler) CreateUser(u *pb.UserItem) (*pb.CreateUserResponse, error) {
	return bh.grpcClient.CreateUser(context.Background(), &pb.CreateUserRequest{User: u})
}

func (bh *BaseUserHandler) UpdateUser(u *pb.UserItem) (*pb.UpdateUserResponse, error) {
	return bh.grpcClient.UpdateUser(context.Background(), &pb.UpdateUserRequest{User: u})
}

func (bh *BaseUserHandler) DeleteUser(id uint32) (*pb.DeleteUserResponse, error) {
	return bh.grpcClient.DeleteUser(context.Background(), &pb.DeleteUserRequest{Id: id})
}
