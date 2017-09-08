package models

import (
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Sku              int    `json:"sku"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	AgeRestriction   int        `json:"age_restriction"`
	IsAlcohol        bool        `json:"is_alcohol"`
	OwnLabel         bool        `json:"own_label"`
}

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func (p Product) Save() (Product) {
	InitDB()
	query := fmt.Sprintf("INSERT IGNORE INTO product(sku, age_restriction, short_description, description, "+
		"is_alcohol, own_label) VALUES (%d , %d, '%s', '%s', %t, %t)",
		p.Sku, p.AgeRestriction, p.ShortDescription, p.Description, p.IsAlcohol, p.OwnLabel)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
func Get(sku int) (Product) {
	InitDB()
	query := fmt.Sprintf("SELECT sku,age_restriction, short_description, description "+
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

func GetAll() ([]*Product, error) {
	InitDB()
	rows, err := db.Query("SELECT decription FROM product limit 10")
	if err != nil {
		log.Fatal(err.Error())
	}
	ps := make([]*Product, 0)

	for rows.Next() {
		p := new(Product)
		err := rows.Scan(&p.Description)
		if err != nil {
			return ps, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}
