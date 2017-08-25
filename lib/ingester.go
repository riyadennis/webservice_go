package lib

import (
	"os"
	"log"
	"fmt"
)

func ReadFile(fileName string){
	file, err := os.Open(fileName)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(file)
}
