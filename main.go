package main

import (
	"fmt"
	"runtime"
	"sync"
)

var a sync.WaitGroup

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GORoutines", runtime.NumGoroutine())

	a.Add(1)
	go oddNum()
	evenNum()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GORoutines", runtime.NumGoroutine())
	a.Wait()
}

// Prints odd numbers
func oddNum() {

	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
	a.Done()
}

// Prints even numbers
func evenNum() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
