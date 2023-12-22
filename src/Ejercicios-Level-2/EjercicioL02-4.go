package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Println(b)
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}
