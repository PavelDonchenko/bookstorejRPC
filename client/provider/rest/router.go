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

	router.GET("/users/:id", ch.GetOne)
	router.GET("/users", ch.GetAll)
	router.POST("/users", ch.Create)
	router.PUT("/users/update", ch.Update)
	router.DELETE("/users/:id", ch.Delete)

	router.GET("/books/:id", bh.GetOne)
	router.GET("/books", bh.GetAll)
	router.POST("/books", bh.Create)
	router.PUT("/books/update", bh.Update)
	router.DELETE("/books/:id", bh.Delete)

	router.Run(":8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}

//func CreateUserRouter(cu *controllers.BaseUserHandler) {
//	ch := rest.NewRouterHandler(cu)
//	router := gin.Default()
//
//	router.GET("/users/:id", ch.GetOne)
//	router.GET("/users", ch.GetAll)
//	router.POST("/users", ch.Create)
//	router.PUT("/users/update", ch.Update)
//	router.DELETE("/users/:id", ch.Delete)
//
//	router.Run(":8081")
//
//	log.Fatal(http.ListenAndServe(":8081", nil))
//}
//
//func CreateBookRouter(cb *controllers3.BaseBookHandler)  {
//	bh := rest.NewRouterHandler(cb)
//	router := gin.Default()
//
//	router.GET("/book/:id", bh.GetOne)
//	router.GET("/book", bh.GetAll)
//	router.POST("/book", bh.Create)
//	router.PUT("/book/update", bh.Update)
//	router.DELETE("/book/:id", bh.Delete)
//}
