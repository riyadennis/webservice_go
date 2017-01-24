package main

import (
	"fmt"
	"net/http"

	"github.com/mux"
	"github.com/webservice_go/developers"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Application starting to run")
	r.HandleFunc("/developer/{id}", developers.GetDeveloper)
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
