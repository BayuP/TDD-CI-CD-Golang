package repository

import (
	"ecommerce/common"
	"ecommerce/model"
	"fmt"
	"log"
)

func GetAllProduct() (interface{}, error) {
	// rows, err := common.DB.Query(`Select id, product_name, price, qty from ms_product where id = $1`, idProduct)

	rows, err := common.DB.Query(`Select id, product_name, price, qty from ms_product`)

	if err != nil {
		fmt.Println("error dinii", err)
		//log.Fatal(err)
		return nil, err
	}

	var products []model.ProductRes

	for rows.Next() {
		var product model.ProductRes
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Qty); err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil

}
