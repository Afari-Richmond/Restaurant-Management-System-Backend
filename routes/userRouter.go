
package routes


import (
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management-system-backend/controllers"
)

func UserRoutes(router *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}