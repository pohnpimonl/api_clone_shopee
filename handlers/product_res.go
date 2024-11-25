package handlers

type ProductGetRes struct {
	ProductID string `json:"productId" binding:"required"`
}
