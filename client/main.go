package main

import (
	controllers3 "github.com/PavelDonchenko/bookstorejRPC/client/controllers/book"
	controllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/bookHandler"
	controllers2 "github.com/PavelDonchenko/bookstorejRPC/client/controllers/user"
	"github.com/PavelDonchenko/bookstorejRPC/client/provider/rest"
	"google.golang.org/grpc"

	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
)

func main() {
	conn, err := grpc.Dial(":8800", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	clientUser := pb.NewUserClient(conn)
	clientBook := pb.NewBookClient(conn)
	clientBookHistory := pb.NewBookHistoryClient(conn)

	handlerUser := controllers2.NewBaseUserHandler(clientUser)
	handlerBook := controllers3.NewBaseBookHandler(clientBook)
	handlerBookHistory := controllers.NewBaseBookHistoryHandler(clientBookHistory)

	rest.CreateAllRoutes(handlerUser, handlerBook, handlerBookHistory)
}
