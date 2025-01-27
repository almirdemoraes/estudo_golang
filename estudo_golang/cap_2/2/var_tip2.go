package main

import "fmt"

func main() {
	numerodebytes, erros := fmt.Println("Hello world!", "Oi galera", 100)
	fmt.Println(numerodebytes, erros)

	x := 16
	y := "string"
	z := true

	fmt.Println(x,y,z)
}