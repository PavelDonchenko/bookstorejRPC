package service

import (
	pb "github.com/PavelDonchenko/bookstorejRPC/server/gen/proto"
	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
	repository2 "github.com/PavelDonchenko/bookstorejRPC/server/repository"
	repository "github.com/PavelDonchenko/bookstorejRPC/server/repository/user"
	"github.com/PavelDonchenko/bookstorejRPC/server/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	userRepo repository2.UserRepository
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) GetUser(id uint32) (*pb.GetUserResponse, error) {
	user, err := us.userRepo.GetUser(id)

	item := pb.UserItem{
		Id:       uint32(user.ID),
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: user.Password,
		CrateAt:  timestamppb.New(user.CreatedAt),
		UpdateAt: timestamppb.New(user.UpdatedAt),
	}
	return &pb.GetUserResponse{User: &item}, err
}

func (us *UserService) GetAllUsers(page uint32) (*pb.GetAllUsersResponse, error) {

	offset, limit := utils.Pagination(page)

	users, err := us.userRepo.GetAllUsers(offset, limit)

	items := []*pb.UserItem{}

	for _, v := range users {
		u := &pb.UserItem{
			Id:       uint32(v.ID),
			Nickname: v.Nickname,
			Email:    v.Email,
			Password: v.Password,
			CrateAt:  timestamppb.New(v.CreatedAt),
			UpdateAt: timestamppb.New(v.UpdatedAt),
		}
		items = append(items, u)
	}

	return &pb.GetAllUsersResponse{User: items}, err
}

func (us *UserService) CreateUser(u model.User) (*pb.CreateUserResponse, error) {
	user, err := us.userRepo.CreateUser(u)

	item := pb.UserItem{
		Id:       uint32(user.ID),
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: user.Password,
		CrateAt:  timestamppb.New(user.CreatedAt),
		UpdateAt: timestamppb.New(user.UpdatedAt),
	}

	return &pb.CreateUserResponse{User: &item}, err

}

func (us *UserService) UpdateUser(u model.User) (*pb.UpdateUserResponse, error) {
	user, err := us.userRepo.UpdateUser(u)

	item := pb.UserItem{
		Id:       uint32(user.ID),
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: user.Password,
		CrateAt:  timestamppb.New(user.CreatedAt),
		UpdateAt: timestamppb.New(user.UpdatedAt),
	}

	return &pb.UpdateUserResponse{User: &item}, err
}

func (us *UserService) DeleteUser(id uint32) (*pb.DeleteUserResponse, error) {
	result, err := us.userRepo.DeleteUser(id)

	return &pb.DeleteUserResponse{Success: result}, err
}
