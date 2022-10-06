package main

import (
	controllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/user"
	"github.com/PavelDonchenko/bookstorejRPC/client/provider/rest"
	"google.golang.org/grpc"

	pb "github.com/PavelDonchenko/bookstorejRPC/client/gen/proto"
)

func main() {
	conn, err := grpc.Dial(":8800", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := pb.NewUserClient(conn)

	c := controllers.NewBaseHandler(client)

	rest.CreateRouter(c)

}
