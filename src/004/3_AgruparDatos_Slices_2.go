package main

import f "fmt"

func main() {
	//TIPO{elementos} //COMPOSITE LITERAL
	x := []int{1, 2, 3, 4, 5}
	f.Println(x)
	f.Println(len(x)) //LARGO
	f.Println(cap(x)) //CAPACIDAD
	f.Println(x[:3])
	f.Println(x[2:])

	for i, v := range x {
		f.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		f.Println(i, x[i])
	}
}
