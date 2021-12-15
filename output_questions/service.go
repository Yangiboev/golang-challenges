package main

import (
	"fmt"
	"runtime"
)

func service(c chan int) {
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}
func main() {
	fmt.Println("main started")
	c := make(chan int, 3)
	go service(c)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println()
	fmt.Println()

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	fmt.Println()
	fmt.Println()

	fmt.Println(runtime.NumGoroutine())
}
