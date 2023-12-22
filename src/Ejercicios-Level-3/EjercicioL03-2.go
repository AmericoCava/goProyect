package main

import f "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		f.Println(i)
		for j := 0; j < 3; j++ {
			f.Printf("\t%#U\n", j)
		}
	}
}
