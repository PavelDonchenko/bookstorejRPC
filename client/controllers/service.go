package controllers

import pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"

type UserApiController interface {
	GetAll() (*pb.GetAllAUserResponse, error)
	Get(id uint32) (*pb.GetUserResponse, error)
	Create(u *pb.UserItem) (*pb.CreateUserResponse, error)
	Update(u *pb.UserItem) (*pb.UpdateUserResponse, error)
	Delete(id uint32) (*pb.DeleteUserResponse, error)
}
