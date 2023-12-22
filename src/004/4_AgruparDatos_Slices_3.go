package main

import f "fmt"

func main() {
	// AGERGAR ELEMENTOS
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6, 7, 8)
	f.Println(x)

	y := []int{9, 10, 11}
	x = append(x, y...)
	f.Println(x)
}
