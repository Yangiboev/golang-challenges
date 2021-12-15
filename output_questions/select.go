package main

import "fmt"

func service1(c chan string) {
	c <- fmt.Sprint("Hello from service1")
}
func service2(c chan string) {
	c <- fmt.Sprint("Hello from service2")
}
func main() {
	fmt.Println("main started")
	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res1 := <-chan1:
		fmt.Println(res1)
	case res2 := <-chan2:
		fmt.Println(res2)
	default:
		fmt.Println("No message received")
	}
}
