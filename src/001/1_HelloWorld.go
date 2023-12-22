package main

import (
	"fmt"
) 

func main()  {
	nb, err := fmt.Println("Hello World")
	_, _ = fmt.Println("Hello World")
	fmt.Println(nb, err)
}