package main

import (
	"fmt"

	"github.com/antfley/go-auth/config"
	"github.com/antfley/go-auth/routes"
	"github.com/antfley/go-auth/services"
	"github.com/antfley/go-auth/services/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	Router *gin.Engine
}

func (server *Server) InitServer() {
	server.Router = gin.Default()
	var Router = routes.RouteLoader{}
	for _, routes := range Router.LoadRoutes() {
		routes.Route(server.Router)
	}
}
func (server *Server) InjectDB() {
	services.InjectDBIntoServices(server.db)
}
func (server *Server) Run() {
	fmt.Println("Rise and shine! 🌞🌞🌞")
	fmt.Println("Listening on port : 5050")
	server.Router.Run("127.0.0.1:5050")
}
func main() {
	DBConfig := db.DBConfig{
		DBHost: config.Config("DB_HOST"),
		DBUser: config.Config("DB_USER"),
		DBName: config.Config("DB_NAME"),
		DBPort: config.Config("DB_PORT"),
		DbPass: config.Config("DB_PASS"),
	}
	app := Server{db: db.InitDB(DBConfig)}
	// db.DBMigrate(app.db)
	app.InjectDB()
	app.InitServer()
	app.Run()
}
