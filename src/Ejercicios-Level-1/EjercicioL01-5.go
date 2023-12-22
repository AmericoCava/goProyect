package main

import (
	"fmt"
)

type (
	numero int
)

var (
	x numero
	y int
)

func main() {
	fmt.Println(x)
	fmt.Printf("El Tipo de: %T\n", x)

	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
