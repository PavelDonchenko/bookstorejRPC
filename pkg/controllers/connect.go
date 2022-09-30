package controllers

import (
	"fmt"
	"github.com/PavelDonchenko/bookstoreCRUD/pkg/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize(Dbdriver string) {
	var err error
	if Dbdriver == "mysql" {
		dns := "pavel:mysqlpaha100688@tcp(bookstore-mysql)/testdb2?charset=utf8mb4&parseTime=True&loc=Local"
		s.DB, err = gorm.Open(Dbdriver, dns)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Dbdriver)
		}
	}

	s.DB.Debug().AutoMigrate(&models.User{}, &models.Book{}) //database migration

	s.Router = mux.NewRouter()

	s.RegisterBookStoreRoutes()
}

func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 6666")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
