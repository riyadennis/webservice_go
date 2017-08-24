package main

import (
	"fmt"
	"./models"
)

func main() {
	fmt.Println("Application starting to run")
	sku := 24
	newProduct := models.SaveProduct(sku, "testing insert","testing", 18, false, false)
	fmt.Println(newProduct)
	product := models.GetProduct(sku)
	fmt.Println(product)
}
