package main

import (
	"fmt"
	"github.com/webservice_go/lib"
	"github.com/webservice_go/config"
	"os"
	"log"
)

func main() {
	fmt.Println("Application starting to run")
	fmt.Println("Reading the configs")
	p, _ := os.Getwd()
	config, err := config.GetConfig(p + "/webservice_go/config.yml")
	if err!=nil {
		log.Fatalf("Unable to load configuration : %s", err.Error())
	}
	pwd, _ := os.Getwd()
	fmt.Println(lib.ReadFileWriteToKafka(pwd + "/webservice_go/"+config.Kafka.File))

	articleReader := lib.ArticleReader{
		Url:config.Article.Url+"?source="+config.Article.Source+"&sortBy=top",
		Body: nil,
		Key: config.Article.Key,
	}
	articleReader.Read()
}
