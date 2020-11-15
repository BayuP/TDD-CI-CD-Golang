package model

type Product struct {
	Base
	ID           string `json:"id"`
	Product_Name string `json:"product_name"`
	Qty          int    `json:"Qty"`
	Price        int    `json:"Price"`
}

type ProductRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Qty   int    `json:"qty"`
	Price int    `json:"price"`
}
