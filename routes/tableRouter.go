
package routes 

import(
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management-system-backend/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:table_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controller.UpdateTable())
	incomingRoutes.DELETE("/tables/:table_id", controller.DeleteTable())
}