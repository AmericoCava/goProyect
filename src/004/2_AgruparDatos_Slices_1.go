package main

import f "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	f.Println(x)

	for i, v := range x {
		f.Println(i, v)
	}
}
