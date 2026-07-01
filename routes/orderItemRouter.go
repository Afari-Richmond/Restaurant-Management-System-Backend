
package routes 

import(
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management-system-backend/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order-items", controller.GetOrderItems())
	incomingRoutes.GET("/order-items/:order_item_id", controller.GetOrderItem())
	incomingRoutes.POST("/order-items", controller.CreateOrderItem())
	incomingRoutes.PATCH("/order-items/:order_item_id", controller.UpdateOrderItem())
	incomingRoutes.DELETE("/order-items/:order_item_id", controller.DeleteOrderItem())
	incomingRoutes.GET("/order-items/order/:order_id", controller.GetOrderItemsByOrderID())
}