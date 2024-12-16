package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pohnpimonl/api_clone_shopee/handlers"
)

func Router(db *pgxpool.Pool) *gin.Engine {
	router := gin.New()
	productAPI := handlers.NewProductAPI(db)

	v1 := router.Group("/v1")
	v1.POST("/products.lists", productAPI.ListProducts)
	v1.POST("/products.get", productAPI.GetProducts)
	v1.POST("/products.create", productAPI.CreateProduct)

	return router
}
