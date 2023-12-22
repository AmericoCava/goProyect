package main

import "fmt"

const (
	a = 2020 + iota
	b = 2021
	c = 2022
	d = 2023
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
