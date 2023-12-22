package main

import (
	"fmt"
)

type numero int

var x numero

func main() {
	fmt.Println(x)
	fmt.Printf("El Tipo de: %T\n", x)

	x = 42
	fmt.Println(x)
}
