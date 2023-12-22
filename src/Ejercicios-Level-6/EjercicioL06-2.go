package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := foo(s...)
	fmt.Println(s1)

	s2 := fooo(s)
	fmt.Println(s2)
}

func foo(v ...int) int {
	total := 0

	for _, v := range v {
		total += v
	}
	return total
}

func fooo(v []int) int {
	total := 0

	for _, v := range v {
		total += v
	}
	return total
}
