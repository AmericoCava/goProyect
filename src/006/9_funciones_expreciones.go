package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hola")
	}

	f()

	g := func(x int) {
		fmt.Println(x)
	}

	g(80)
}
