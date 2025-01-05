package main

import "fmt"

var x int = 42
var y string ="James Bond"
var z bool =true

func main(){
    //fmt.Printf("%v\n%v\n%v\n", x,y,z)
    s:= fmt.Sprintf("%v\t%v\t%v", x, y, z)

    fmt.Printf(s)
    
}