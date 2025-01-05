package main

import (
	"fmt"
)

func main(){
	x := 10
	//fmt.Printf("%b\n%d\n%x", x,x,x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x", y, y, y)


}