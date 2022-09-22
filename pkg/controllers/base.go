package controllers

import (
	"fmt"
	"github.com/PavelDonchenko/40projects/go-bookstore/pkg/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

//var (
//	db *gorm.DB
//)
//
//func Connect() {
//	dns := "pavel:mysqlpaha100688@tcp(127.0.0.1:3306)/testdb2?charset=utf8mb4&parseTime=True&loc=Local"
//	d, err :=
//		gorm.Open(mysql.Open(dns), &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//	db = d
//	fmt.Println(db)
//}
//
//func GetDB() *gorm.DB {
//	return db
//}

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize(Dbdriver string) {

	var err error

	if Dbdriver == "mysql" {
		dns := "pavel:mysqlpaha100688@tcp(127.0.0.1:3306)/testdb2?charset=utf8mb4&parseTime=True&loc=Local"
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
