package routes

import (
	"github.com/antfley/go-auth/controllers"
	"github.com/gin-gonic/gin"
)

type User struct {}
func (p *User) Route(route *gin.Engine) {
	router := route.Group("/user")
	Controller := controllers.UserController{}
	router.POST("/register", Controller.Register)
	router.POST("/login", Controller.Login)
}
