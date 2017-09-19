package main

import (
	"fmt"
	"github.com/webservice_go/lib"
	"github.com/webservice_go/config"
)

func main() {
	fmt.Println("Application starting to run")
	fmt.Println("Reading the configs")
	config := config.GetConfig()
	//pwd, _ := os.Getwd()
	//fmt.Println(lib.ReadFileWriteToKafka(pwd + config.Kafka.File))

	lib.ReadArticles(config.Article.Url+"?"+config.Article.Source+"&"+config.Article.SortOption, nil, config.Article.Key)
}
