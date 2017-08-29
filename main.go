package main

import (
	"fmt"
	"./lib"
)

func main() {
	fmt.Println("Application starting to run")
	file := "data/BigDataFile.dat"
	fmt.Println(lib.ReadFileWriteToKafka(file))
}
