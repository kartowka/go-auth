package routes

import (
	"github.com/gin-gonic/gin"
)

type RouterInterface interface{
	Route(*gin.Engine)
}
type RouteLoader struct {}

func (r *RouteLoader) LoadRoutes() []RouterInterface {
	user := new (User)
	return []RouterInterface{user}
}
