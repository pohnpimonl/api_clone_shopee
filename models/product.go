package models

type Product struct {
	ID          string  `json:"productId"`
	Name        string  `json:"productName"`
	Description string  `json:"productDescription"`
	Price       float64 `json:"productPrice"`
}
