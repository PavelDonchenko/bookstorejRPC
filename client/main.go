package main

import (
	bookControllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/book"
	bookHistoryControllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/bookHistory"
	userControllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/user"
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

	handlerUser := userControllers.NewBaseUserHandler(clientUser)
	handlerBook := bookControllers.NewBaseBookHandler(clientBook)
	handlerBookHistory := bookHistoryControllers.NewBaseBookHistoryHandler(clientBookHistory)

	rest.CreateAllRoutes(handlerUser, handlerBook, handlerBookHistory)
}
