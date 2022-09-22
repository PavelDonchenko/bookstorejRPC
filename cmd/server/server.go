package server

import (
	"github.com/PavelDonchenko/40projects/go-bookstore/pkg/controllers"
)

var server = controllers.Server{}

func Run() {

	server.Initialize("mysql")

	server.Run("localhost:6666")

}
