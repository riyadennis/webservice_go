package main

import (
	"fmt"
	"math"
)

type welcome struct {

}
//reciever is a struct
func (w *welcome) String() string{
	return "Hello"
}
type Power struct{
	x float64
	y float64
}
// a regular function
func findPower(p Power) float64{
	return math.Pow(p.x, p.y)
}

func main(){
	fmt.Printf("%s \n", new(welcome))
	p := Power{10, 12.2}
	fmt.Printf("%.2f \n", findPower(p))
}
