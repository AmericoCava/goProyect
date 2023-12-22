package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := []int{52}
	x = append(x, y...)
	fmt.Println(x)

	z := []int{53, 54, 55}
	x = append(x, z...)
	fmt.Println(x)

	w := []int{56, 57, 58, 59, 60}
	x = append(x, w...)
	fmt.Println(x)
}
