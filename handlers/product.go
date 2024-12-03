package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductAPI struct {
	db *pgxpool.Pool
}

func NewProductAPI(db *pgxpool.Pool) *ProductAPI {
	return &ProductAPI{
		db: db,
	}
}

func (h *ProductAPI) ListProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List Product"})
}

func (h *ProductAPI) GetProducts(ctx *gin.Context) {
	var req ProductGetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ProductID == "111" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product 111 not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Found Product : " + req.ProductID})
}
