package main

import (
	"fmt"
)

func main (){
	x:= "oi"
	y:="bom dia"

	z:= fmt.Sprint(x,y)

	println(z)
}

//Sprint "junta" strings e só 