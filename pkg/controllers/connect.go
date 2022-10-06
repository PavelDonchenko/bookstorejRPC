package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PavelDonchenko/bookstoreCRUD/pkg/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize(Dbdriver, user, password, host, dbname string) {
	var err error
	if Dbdriver == "mysql" {
		dns := user + ":" + password + "@tcp(" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
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
	fmt.Println("Listening to port 8800")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
