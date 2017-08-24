package models

import (
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Sku              string `json:"sku"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	AgeRestriction   int    `json:"age_restriction"`
}

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func GetProduct(sku int) (Product) {
	InitDB()
	query := fmt.Sprintf("SELECT sku,age_restriction, short_description, description " +
		"FROM product WHERE sku = %d", sku)
	row := db.QueryRow(query)

	defer db.Close()
	var product Product

	err := row.Scan(&product.Sku,
		&product.AgeRestriction,
		&product.ShortDescription,
		&product.Description,
	)
	if err != nil {
		log.Fatal(err)
	}
	return product
}
