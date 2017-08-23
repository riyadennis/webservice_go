package products

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func GetProduct() {
	db, err := sql.Open("mysql", "root:@/api_products?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT sku,age_restriction, short_description, description FROM product WHERE sku=24")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()
	defer db.Close()
	var product Product
	for rows.Next() {
		err := rows.Scan(&product.Sku,
			&product.AgeRestriction,
			&product.ShortDescription,
			&product.Description,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(product)
	}

}
func createResponse(statusCode int, message string, w http.ResponseWriter) {
	res := Response{
		Code:   statusCode,
		Result: fmt.Sprintf("%s", message),
	}

	jsonMsg, err := json.Marshal(res)
	if err != nil {
		log.Fatal("Unable to marshal the json")
	}
	w.Header().Set("application/type", "json")
	w.WriteHeader(statusCode)
	w.Write(jsonMsg)
}
