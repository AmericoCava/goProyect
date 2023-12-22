package main

import (
	"fmt"
)

var z = 80

func main(){
	var w int
	x := 42+10
	y := "test"
	fmt.Println(x)
	fmt.Println(y)
	x = 50
	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(w)
	numero()
}

func numero(){
	fmt.Println(z)
}