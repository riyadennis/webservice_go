package main

import (
	"fmt"
	"log"
	"os"

	"github.com/riyadennis/webservice_go/config"
	"github.com/riyadennis/webservice_go/lib"
)

func main() {
	fmt.Println("Application starting to run")
	fmt.Println("Reading the configs")
	p, _ := os.Getwd()
	cfg, err := config.GetConfig(p + "/webservice_go/cfg.yml")
	if err != nil {
		log.Fatalf("Unable to load configuration : %s", err.Error())
	}
	pwd, _ := os.Getwd()
	fmt.Println(lib.ReadFileWriteToKafka(pwd + "/webservice_go/" + cfg.Kafka.File))

	articleReader := lib.ArticleReader{
		Url:  cfg.Article.Url + "?source=" + cfg.Article.Source + "&sortBy=top",
		Body: nil,
		Key:  cfg.Article.Key,
	}
	articleReader.Read()
}
