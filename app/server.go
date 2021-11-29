package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) initialize() {
	fmt.Println("Welcome To Go-Store")

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) run(port string) {
	fmt.Printf("Server Running On Port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server.Router))
}

func Run() {
	server := Server{}
	server.initialize()
	server.run("8080")
}
