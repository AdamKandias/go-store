package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func envConfig(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (server *Server) initialize(appConfig AppConfig) {
	fmt.Println("Welcome To", appConfig.AppName)

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) run(port string) {
	fmt.Printf("Server Running On Port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, server.Router))
}

func Run() {
	server := Server{}
	if err := godotenv.Load(); err != nil {
		log.Fatal("error on load .env file")
	}
	appConfig := AppConfig{}
	appConfig.AppName = envConfig("APP_NAME", "Go-Toko")
	appConfig.AppEnv = envConfig("APP_ENV", "development")
	appConfig.AppPort = envConfig("APP_PORT", "8080")

	server.initialize(appConfig)
	server.run(appConfig.AppPort)
}
