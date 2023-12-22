package main

import "fmt"

func main() {
	a := incremental()
	b := incremental()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incremental() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
