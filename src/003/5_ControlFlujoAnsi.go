package main

import f "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		f.Printf("%d\t%#x\t%#U\n", i, i, i)
	}
}
