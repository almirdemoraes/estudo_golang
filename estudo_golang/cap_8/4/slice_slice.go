package main

import(
	"fmt"
)

func main(){		/*  1        2         3        4        5          6*/
	tenis := []string{"adidas", "rebook", "nike", "umbro", "penalty", "topper"}

	marca := tenis[:]

	fmt.Println(marca)
}