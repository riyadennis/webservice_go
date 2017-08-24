package main

import (
	"fmt"
	"./models"
)

func main() {
	fmt.Println("Application starting to run")
	product := models.GetProduct(24)
	fmt.Println(product)
}
