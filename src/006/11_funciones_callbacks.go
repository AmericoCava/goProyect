package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println(s)

	s1 := sumPar(sum, ii...)
	fmt.Println(s1)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}

func sumPar(f func(xi ...int) int, vi ...int) int {
	var y []int

	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
