//package controllers
//
//import (
//	"context"
//
//	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
//)
//
//type BaseUserHandler struct {
//	grpcClient pb.UserClient
//}
//
//type BaseBookHandler struct {
//	grpcClient pb.BookClient
//}
//
//func NewBaseHandler(clientUser pb.UserClient, clientBook pb.BookClient) (*BaseUserHandler, *BaseBookHandler) {
//	return &BaseUserHandler{grpcClient: clientUser}, &BaseBookHandler{grpcClient: clientBook}
//}
//
//func (bh *BaseBookHandler) GetAll() (*pb.GetAllBooksResponse, error) {
//	return bh.grpcClient.GetAll(context.Background(), &pb.GetAllBooksRequest{})
//}
//func (bh *BaseBookHandler) Get(id uint32) (*pb.GetBookResponse, error) {
//	return bh.grpcClient.GetOne(context.Background(), &pb.GetBookRequest{Id: id})
//}
//func (bh *BaseBookHandler) Create(u *pb.BookItem) (*pb.CreateBookResponse, error) {
//	return bh.grpcClient.Create(context.Background(), &pb.CreateBookRequest{Book: u})
//}
//func (bh *BaseBookHandler) Update(u *pb.BookItem) (*pb.UpdateBookResponse, error) {
//	return bh.grpcClient.Update(context.Background(), &pb.UpdateBookRequest{Book: u})
//}
//func (bh *BaseBookHandler) Delete(id uint32) (*pb.DeleteBookResponse, error) {
//	return bh.grpcClient.Delete(context.Background(), &pb.DeleteBookRequest{Id: id})
//}
//
//func (bh *BaseUserHandler) GetAll() (*pb.GetAllAUserResponse, error) {
//	return bh.grpcClient.GetAll(context.Background(), &pb.GetAllAUserRequest{})
//}
//func (bh *BaseUserHandler) Get(id uint32) (*pb.GetUserResponse, error) {
//	return bh.grpcClient.GetOne(context.Background(), &pb.GetUserRequest{Id: id})
//}
//func (bh *BaseUserHandler) Create(u *pb.UserItem) (*pb.CreateUserResponse, error) {
//	return bh.grpcClient.Create(context.Background(), &pb.CreateUserRequest{User: u})
//}
//func (bh *BaseUserHandler) Update(u *pb.UserItem) (*pb.UpdateUserResponse, error) {
//	return bh.grpcClient.Update(context.Background(), &pb.UpdateUserRequest{User: u})
//}
//func (bh *BaseUserHandler) Delete(id uint32) (*pb.DeleteUserResponse, error) {
//	return bh.grpcClient.Delete(context.Background(), &pb.DeleteUserRequest{Id: id})
//}
//
