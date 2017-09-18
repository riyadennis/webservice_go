package main

import (
	"fmt"
	"github.com/webservice_go/lib"
)

func main() {
	fmt.Println("Application starting to run")
	fmt.Println("Reading the configs")
	//config := config.GetConfig()
	//pwd, _ := os.Getwd()
	//fmt.Println(lib.ReadFileWriteToKafka(pwd + config.Kafka.File))

	lib.ReadArticles("https://newsapi.org/v1/articles?source=bbc-news&sortBy=top", nil)
}
