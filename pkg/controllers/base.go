package config

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

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Book{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
