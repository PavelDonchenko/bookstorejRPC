package grpcHandler

import (
	"context"
	"time"

	pb "github.com/PavelDonchenko/bookstorejRPC/server/gen/proto"
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
	service "github.com/PavelDonchenko/bookstorejRPC/server/service/user"
)

type UserHandler struct {
	us *service.UserService
	pb.UnimplementedUserServer
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{us: userService}
}

func (uh *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return uh.us.GetUser(req.GetId())
}

func (uh *UserHandler) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	return uh.us.GetAllUsers(req.GetPage())
}

func (uh *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := model.User{
		ID:        req.GetUser().GetId(),
		Nickname:  req.GetUser().GetNickname(),
		Email:     req.GetUser().GetEmail(),
		Password:  req.GetUser().GetPassword(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return uh.us.CreateUser(user)
}

func (uh *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := model.User{
		ID:        req.GetUser().GetId(),
		Nickname:  req.GetUser().GetNickname(),
		Email:     req.GetUser().GetEmail(),
		Password:  req.GetUser().GetPassword(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return uh.us.UpdateUser(user)
}

func (uh *UserHandler) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return uh.us.DeleteUser(req.GetId())
}
