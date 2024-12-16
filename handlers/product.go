package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pohnpimonl/api_clone_shopee/models"
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
	var products []models.Product
	rows, err := h.db.Query(ctx, "SELECT id, name, description, price FROM product ORDER BY created_at DESC")
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something Wrong!!"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something Wrong!!"})
			return
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something Wrong!!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"results": products})
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

func (h *ProductAPI) CreateProduct(ctx *gin.Context) {
	var req CreateProductReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.Must(uuid.NewRandom())

	_, err := h.db.Exec(ctx, "INSERT INTO product (id,name,description,price,created_by,updated_by) VALUES ($1, $2, $3, $4, $5, $6)", id, req.Name, req.Description, req.Price, "joy", "joy")

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something Wrong!!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": "success"})
}
