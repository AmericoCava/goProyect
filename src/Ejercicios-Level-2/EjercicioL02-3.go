package main

import "fmt"

const (
	x        = 42
	y string = "42"
)

func main() {
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
