package main

import f "fmt"

func main() {
	// ELIMINAR ELEMENTOS
	x := []int{4, 5, 7, 8, 42}
	x = append(x, 44, 55, 66)

	y := []int{333, 444, 666}
	x = append(x, y...)
	f.Println(x)

	x = append(x[:5], x[7:]...)
	f.Println(x)
}
