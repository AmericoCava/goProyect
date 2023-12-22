package main

import "fmt"

func main() {
	x := "Batman"

	if x == "robin" {
		fmt.Println(x)
	} else if x == "Batman" {
		fmt.Println(x + " new")
	} else {
		fmt.Println("Ninguno")
	}
}
