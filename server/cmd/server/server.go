package server

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/PavelDonchenko/bookstorejRPC/server/gen/proto"
	grpcHandler "github.com/PavelDonchenko/bookstorejRPC/server/provider/grpc"
	repository "github.com/PavelDonchenko/bookstorejRPC/server/repository/user"
	service "github.com/PavelDonchenko/bookstorejRPC/server/service/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	DB *gorm.DB
)

func InitializeDB(Dbdriver, user, password, host, dbname string) *gorm.DB {
	var err error
	if Dbdriver == "mysql" {
		dns := user + ":" + password + "@tcp(" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(Dbdriver, dns)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Dbdriver)
		}
	}

	return DB
	//RegisterBookStoreRoutes()
}

func RunServer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := InitializeDB("mysql", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))

	sqlDB := db.DB()

	defer sqlDB.Close()

	userRepo := repository.NewUserRepo(db)

	s := grpc.NewServer()
	us := service.NewUserService(userRepo)
	uh := grpcHandler.NewUserHandler(us)

	pb.RegisterUserServer(s, uh)

	fmt.Println("Server successfully started on port :8800")
	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		log.Fatalf("Unable to listen on port :8800: %v", err)
	}

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
