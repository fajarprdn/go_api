package apprequest

type ProductRequest struct {
	ProductCode string `json:"product_code" binding:"required"`
	ProductName string `json:"product_name" binding:"required"`
}
