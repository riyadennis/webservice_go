package main

import (
	"fmt"
	"./config"
)

func main() {
	fmt.Println("Application starting to run")
	fmt.Println("Reading the configs")
	fmt.Println(config.GetConfig())
	//file := "data/BigDataFile.dat"
	//fmt.Println(lib.ReadFileWriteToKafka(file))
}
