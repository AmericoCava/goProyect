package main

import "fmt"

func main() {
	func() {
		fmt.Println("re hola")
	}() //parentesis al fiunal para hacerla anonima que significa que no tiene nombre y esta dentro del main

	func(x int) {
		fmt.Println(x)
	}(56)
}
