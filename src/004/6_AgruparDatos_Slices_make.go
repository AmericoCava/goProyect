package main

import f "fmt"

func main() {
	x := make([]int, 10, 10)
	f.Println(x)
	f.Println(len(x))
	f.Println(cap(x))

	x[3] = 42
	x[6] = 43
	f.Println(x)
	f.Println(len(x))
	f.Println(cap(x))

	x = append(x, 45)
	f.Println(x)
	f.Println(len(x))
	f.Println(cap(x))
}
