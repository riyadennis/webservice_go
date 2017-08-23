package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"./developers"
	//if its mysql _ "github.com/ziutek/mymysql/godrv" //will import the package and execute its init function
	_ "github.com/mattn/go-sqlite3" //will import the package and execute its init function
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Application starting to run")
	r.HandleFunc("/developer/{id}", developers.GetDeveloper)
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
