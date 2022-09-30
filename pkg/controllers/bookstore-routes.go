package controllers

import "github.com/PavelDonchenko/bookstoreCRUD/pkg/middlewares"

func (s *Server) RegisterBookStoreRoutes() {
	// login route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//user routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetAllUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUserById)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//book routes
	s.Router.HandleFunc("/books", middlewares.SetMiddlewareJSON(s.CreateBook)).Methods("POST")
	s.Router.HandleFunc("/books", middlewares.SetMiddlewareJSON(s.GetAllBooks)).Methods("GET")
	s.Router.HandleFunc("/books/{id}", middlewares.SetMiddlewareJSON(s.GetBookById)).Methods("GET")
	s.Router.HandleFunc("/books/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateBook))).Methods("PUT")
	s.Router.HandleFunc("/books/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteBook)).Methods("DELETE")
}
