package main

import "fmt"

func main() {
	s1 := foo()
	x, y := bar()
	fmt.Println(y, s1)
	fmt.Println(y, x)
}

func foo() int {
	return 2023
}

func bar() (int, string) {
	return 2023, "Este es el año"
}
