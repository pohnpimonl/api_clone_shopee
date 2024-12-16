package handlers

//comment
type ProductGetReq struct {
	ProductID string `json:"productId" binding:"required"`
}

type CreateProductReq struct {
	Name        string  `json:"productName" binding:"required"`
	Description string  `json:"productDescription" binding:"required"`
	Price       float64 `json:"productPrice" binding:"required"`
}
