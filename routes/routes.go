package routes

import (
	"payment/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type"}
	r.Use(cors.New(config))
	r.GET("products", handlers.GetProducts)
	r.POST("products", handlers.CreateProduct)
	r.GET("pk", handlers.PublishableKey)
	r.POST("cpi", handlers.HandleCreatePaymentIntent)

	return r
}
