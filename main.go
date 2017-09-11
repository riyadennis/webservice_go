package main

import (
	"fmt"
	"./config"
	"./lib"
	"os"
)

func main() {
	fmt.Println("Application starting to run")
	fmt.Println("Reading the configs")
	config := config.GetConfig()
	pwd, _ := os.Getwd()
	fmt.Println(lib.ReadFileWriteToKafka(pwd + config.Kafka.File))
}
