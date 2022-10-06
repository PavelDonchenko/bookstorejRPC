package rest

import (
	"log"
	"net/http"

	controllers "github.com/PavelDonchenko/bookstorejRPC/client/controllers/user"
	rest "github.com/PavelDonchenko/bookstorejRPC/client/provider/rest/user"
	"github.com/gin-gonic/gin"
)

func CreateRouter(c *controllers.BaseHandler) {
	h := rest.NewRouterHandler(c)
	router := gin.Default()

	router.GET("/users/:id", h.GetOne)
	router.GET("/users", h.GetAll)
	router.POST("/users", h.Create)
	router.PUT("/users", h.Update)
	router.DELETE("/users/:id", h.Delete)

	router.Run(":8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
