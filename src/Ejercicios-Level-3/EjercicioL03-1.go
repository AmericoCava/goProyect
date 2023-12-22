package main

import f "fmt"

func main() {
	for i := 0; i <= 10000; i++ {
		if i%2 == 0 {
			f.Println(i)
		}
	}
}
