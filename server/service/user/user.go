package service

import (
	pb "github.com/PavelDonchenko/bookstorejRPC/gen/proto"
	model "github.com/PavelDonchenko/bookstorejRPC/models"
	repository "github.com/PavelDonchenko/bookstorejRPC/repository/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) GetOne(id uint32) (*pb.GetUserResponse, error) {
	user, err := us.userRepo.GetOne(id)

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
func (us *UserService) GetAll() (*pb.GetAllAUserResponse, error) {
	users, err := us.userRepo.GetAll()

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

	return &pb.GetAllAUserResponse{User: items}, err
}
func (us *UserService) Create(u model.User) (*pb.CreateUserResponse, error) {
	user, err := us.userRepo.Create(u)

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
func (us *UserService) Update(u model.User) (*pb.UpdateUserResponse, error) {
	user, err := us.userRepo.Update(u.ID)

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
func (us *UserService) Delete(id uint32) (*pb.DeleteUserResponse, error) {
	result, err := us.userRepo.Delete(id)

	return &pb.DeleteUserResponse{Success: result}, err
}
