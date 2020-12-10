package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)                 // windows
	fmt.Println("ARCH\t\t", runtime.GOARCH)             // amd64
	fmt.Println("CPU\t\t", runtime.NumCPU())            // 12
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // 1

	wg.Add(1) // 1가지 기다릴 것이 있다.

	go foo()
	bar()

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // 2
	wg.Wait()                                           // 끝나지 않고 기다린다. wg.Done이 될 때까지
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}
