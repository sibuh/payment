package routes

import (
	"payment/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("products", handlers.GetProducts)
	r.POST("products", handlers.CreateProduct)
	r.GET("pk", handlers.PublishableKey)
	r.POST("cpi", handlers.HandleCreatePaymentIntent)

	return r
}
