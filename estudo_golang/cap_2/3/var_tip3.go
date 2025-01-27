package main

import (
	"fmt"
)

var y = "Olá, bom dia"
func main(){
	x:= 10
	y:= "olá, bom dia"

	fmt.Printf("x: %v, %T\n", x,x)
	fmt.Printf("y: %v, %T", y,y)

	x=20
	fmt.Printf("x: %v, %T\n", x,x)
}