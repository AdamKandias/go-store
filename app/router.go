package app

import (
	"github.com/AdamKandias/go-store/app/controllers"
)

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
