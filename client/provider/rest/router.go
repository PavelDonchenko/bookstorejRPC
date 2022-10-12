package rest

import (
	"log"
	"net/http"

	controllers3 "github.com/PavelDonchenko/bookstorejRPC/client/controllers/book"
	controllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/user"
	rest2 "github.com/PavelDonchenko/bookstorejRPC/client/provider/rest/book"
	rest "github.com/PavelDonchenko/bookstorejRPC/client/provider/rest/user"
	"github.com/gin-gonic/gin"
)

func CreateAllRoutes(cu *controllers.BaseUserHandler, cb *controllers3.BaseBookHandler) {
	ch := rest.NewRouterUserHandler(cu)
	bh := rest2.NewRouterBookHandler(cb)
	router := gin.Default()

	router.GET("/users/:id", ch.GetUser)
	router.GET("/users", ch.GetAllUsers)
	router.POST("/users", ch.CreateUser)
	router.PUT("/users/:id", ch.UpdateUser)
	router.DELETE("/users/:id", ch.DeleteUser)

	router.GET("/books/:id", bh.GetBook)
	router.GET("/books", bh.GetAllBooks)
	router.POST("/books", bh.CreateBook)
	router.PUT("/books/update", bh.UpdateBook)
	router.DELETE("/books/:id", bh.DeleteBook)

	router.Run(":8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
