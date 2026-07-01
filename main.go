package main 

import(
	"os"
	"github.com/gin-gonic/gin"
	"golang-restaurant-management-system-backend/database"
	"golang-restaurant-management-system-backend/routes"
	"golang-restaurant-management-system-backend/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)


var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.SetupRoutes(router)
	router.Use(middleware.Authtntication())

	routes.FoodRoutes(router)
	routes.UserRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderRoutes(router)
	routes.orderItemRoutes(router)
	routes.TableRoutes(router)
	routes.orderRoutes(router)
	routes.tableRoutes(router)
	routes.MenuRoutes(router)
	
	router.Run(":" + port)
	
}