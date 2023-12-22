package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)
	s2 := bar()
	i := s2()
	fmt.Println(i)
}

func foo() string {
	s := "Hola"
	return s
}

func bar() func() int {
	return func() int {
		return 1492
	}
}
