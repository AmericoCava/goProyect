package main

import "fmt"

func main() {
	f := func() {
		for i := 0; i <= 2; i++ {
			fmt.Println(i)
		}
	}

	f()

	fmt.Printf("%T\n", f)
	fmt.Println("Listooooo")
}
