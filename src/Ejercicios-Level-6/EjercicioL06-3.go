package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hola, playground")
}

func foo() {
	defer func() {
		fmt.Println("foo recorrio")
	}()
	fmt.Println("foo corrio")
}
