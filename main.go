package main

import (
	"fmt"
	"github.com/webservice_go/lib"
	"github.com/webservice_go/config"
	"os"
)

func main() {
	fmt.Println("Application starting to run")
	fmt.Println("Reading the configs")
	config := config.GetConfig()
	pwd, _ := os.Getwd()
	fmt.Println(lib.ReadFileWriteToKafka(pwd + "/src/github.com/webservice_go/"+config.Kafka.File))
	fmt.Println(config.Article.Key)
	lib.ReadArticles(config.Article.Url+"?source="+config.Article.Source+"&sortBy=top", nil, config.Article.Key)
}
