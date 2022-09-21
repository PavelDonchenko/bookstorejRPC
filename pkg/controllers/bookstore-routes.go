package routes

import (
	"github.com/PavelDonchenko/40projects/go-bookstore/pkg/controllers"
	"github.com/PavelDonchenko/40projects/go-bookstore/pkg/middlewares"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	//user routes
	var s *controllers.Server
	router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetAllUsers)).Methods("GET")
	router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUserById)).Methods("GET")
	router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//book routes
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
}
