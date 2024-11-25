package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pohnpimonl/api_clone_shopee/handlers"
)

func Router() *gin.Engine {
	router := gin.New()
	v1 := router.Group("/v1")
	v1.POST("/products.lists", handlers.ListProducts)
	v1.POST("/products.get", handlers.GetProducts)

	return router
}
