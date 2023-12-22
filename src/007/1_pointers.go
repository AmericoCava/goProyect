package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var v *int = &a
	fmt.Println(v)
	fmt.Println(*v)

	*v = 43
	fmt.Println(a)
}
