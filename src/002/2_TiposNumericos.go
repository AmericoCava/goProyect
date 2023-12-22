package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 42
	y := 42.32
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
