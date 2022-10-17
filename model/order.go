package model

type Orders struct {
	TotalPrice int64 `json:"totalPrice"`
	Product    []Product
}
type Product struct {
	ProductName string `json:"productName"`
	Quantity    int32  `json:"quantity"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}

type ProductDal struct {
	ProductName string `json:"productName"`
	Quantity    int32  `json:"quantity"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}
