package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s, e := fmt.Printf("%v\t%v\t%v", x, y, z)
	fmt.Println(s, e)
}
