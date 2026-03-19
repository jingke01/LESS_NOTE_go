package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("hello world") //主协程
	for i := 0; i < 2; i++ {
		//runtime.Gosched()
		fmt.Println("Hello")
	}
}
