package main

import "fmt"

func main() {
	x := 0
	//FOR CON BREAK
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
