package main

import f "fmt"

func main() {
	f.Println(2 == 2 && 3 == 3) //AND
	f.Println(true && false)    //AND
	f.Println(2 == 2 || 3 == 3) //OR
	f.Println(true || false)    //OR
	f.Println(!true)
}
