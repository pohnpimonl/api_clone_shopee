package handlers

type ProductGetReq struct {
	ProductID string `json:"productId" binding:"required"`
}
