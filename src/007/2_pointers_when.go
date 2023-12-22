package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x antes", x)
	fmt.Println("x antes", &x)
	foo(&x)
	fmt.Println("x despues", x)
	fmt.Println("x despues", &x)
}

func foo(y *int) {
	fmt.Println("y antes", y)
	fmt.Println("y antes", *y)
	*y = 42
	fmt.Println("y despues", y)
	fmt.Println("y despues", *y)
}
