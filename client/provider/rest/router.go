package rest

import (
	"log"
	"net/http"

	controllers3 "github.com/PavelDonchenko/bookstorejRPC/client/controllers/book"
	controllers2 "github.com/PavelDonchenko/bookstorejRPC/client/controllers/bookHandler"
	controllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/user"
	rest2 "github.com/PavelDonchenko/bookstorejRPC/client/provider/rest/book"
	rest3 "github.com/PavelDonchenko/bookstorejRPC/client/provider/rest/bookHistory"
	rest "github.com/PavelDonchenko/bookstorejRPC/client/provider/rest/user"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

var esClient *elasticsearch.Client

func CreateAllRoutes(cu *controllers.BaseUserHandler, cb *controllers3.BaseBookHandler, cbh *controllers2.BaseBookHistoryHandler) {
	ch := rest.NewRouterUserHandler(cu)
	bh := rest2.NewRouterBookHandler(cb)
	bhh := rest3.NewRouterBookHistoryHandler(cbh, esClient)
	router := gin.Default()

	router.GET("/user/:id", ch.GetUser)
	router.GET("/users/:page", ch.GetAllUsers)
	router.POST("/users", ch.CreateUser)
	router.PUT("/users/:id", ch.UpdateUser)
	router.DELETE("/users/:id", ch.DeleteUser)

	router.GET("/book/:id", bh.GetBook)
	router.GET("/books/:page", bh.GetAllBooks)
	router.POST("/books", bh.CreateBook)
	router.PUT("/books/update", bh.UpdateBook)
	router.DELETE("/books/:id", bh.DeleteBook)

	router.POST("/book-history", bhh.CreateBookHistory)
	router.GET("/book-history/:id", bhh.GetOneBookHistory)
	router.DELETE("/book-history/:id", bhh.DeleteBookHistory)
	router.GET("/search", bhh.SearchBookHistory)

	router.Run(":8082")

	log.Fatal(http.ListenAndServe(":8082", nil))
}
