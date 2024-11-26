package handlers
//comment
type ProductGetReq struct {
	ProductID string `json:"productId" binding:"required"`
}
