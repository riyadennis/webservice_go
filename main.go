package main

import (
	"fmt"
	"./config"
	"./lib"
)

func main() {
	fmt.Println("Application starting to run")
	fmt.Println("Reading the configs")
	fmt.Println(config.GetConfig())
	config := config.GetConfig()
	file := config.Kafka.File
	fmt.Println(lib.ReadFileWriteToKafka(file))
}
