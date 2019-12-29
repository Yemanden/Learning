package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 10; i++ {
		go work()
	}
	runtime.Gosched()
	runtime.Gosched()
}

func work() {
	counter := 0
	for i := 0; i < 5; i++ {
		fmt.Println(counter)
		counter++
	}
}
