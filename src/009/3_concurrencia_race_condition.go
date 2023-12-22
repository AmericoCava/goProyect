package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("os        \t: ", runtime.GOOS)
	fmt.Println("ARCH      \t: ", runtime.GOARCH)
	fmt.Println("CPUs      \t: ", runtime.NumCPU())
	fmt.Println("Goroutines\t: ", runtime.NumGoroutine())

	contador := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := contador
			v++
			//time.Sleep(time.Second*2)
			runtime.Gosched()
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines\t: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("cuenta", contador)
}
