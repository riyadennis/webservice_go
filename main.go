package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"./products"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Application starting to run")
	products.GetProduct()
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
