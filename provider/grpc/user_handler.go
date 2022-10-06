package grpcHandler

import (
	"context"
	"time"

	pb "github.com/PavelDonchenko/bookstoreCRUD/gen/proto"
	model "github.com/PavelDonchenko/bookstoreCRUD/server/models"
	service "github.com/PavelDonchenko/bookstoreCRUD/service/user"
)

type UserHandler struct {
	us *service.UserService
	pb.UnimplementedUserServer
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{us: userService}
}

func (uh *UserHandler) GetOne(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return uh.us.GetOne(req.GetId())
}

func (uh *UserHandler) GetAll(ctx context.Context, req *pb.GetAllAUserRequest) (*pb.GetAllAUserResponse, error) {
	return uh.us.GetAll()
}

func (uh *UserHandler) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := model.User{
		ID:        req.GetUser().GetId(),
		Nickname:  req.GetUser().GetNickname(),
		Email:     req.GetUser().GetEmail(),
		Password:  req.GetUser().GetPassword(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return uh.us.Create(user)
}

func (uh *UserHandler) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := model.User{
		ID:        req.GetUser().GetId(),
		Nickname:  req.GetUser().GetNickname(),
		Email:     req.GetUser().GetEmail(),
		Password:  req.GetUser().GetPassword(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return uh.us.Update(user)
}

func (uh *UserHandler) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return uh.us.Delete(req.GetId())
}
