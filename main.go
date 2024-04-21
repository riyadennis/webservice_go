package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"

	"github.com/riyadennis/webservice_go/lib"
)

func main() {
	fmt.Println("Application starting to run")
	fmt.Println("Reading the configs")

	err := configCheck()
	if err != nil {
		log.Fatalf("Unable to load configuration : %s", err.Error())
	}

	articleReader := lib.ArticleReader{
		Url:  viper.GetString("article_url") + "?source=" + viper.GetString("article_source") + "&sortBy=top",
		Body: nil,
		Key:  viper.GetString("article_key"),
	}

	err = articleReader.Read()
	if err != nil {
		log.Fatalf("Unable to read article : %s", err.Error())
	}
}

func configCheck() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/") // call multiple times to add many search paths
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		}
		return fmt.Errorf("fatal error config file: %w", err)
	}

	return nil
}
