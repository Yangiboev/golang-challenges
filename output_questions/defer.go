package main

import "fmt"

func defFooStart() {
	fmt.Println("defFooStart executed")
}
func defFooEnd() {
	fmt.Println("defFooEnd executed")
}
func defMain() {
	fmt.Println("defMain executed")
}
func foo() {
	fmt.Println("foo started")
	defer defFooStart()
	panic("panic from foo()")
	defer defFooEnd()
	fmt.Println("foo Done")
}
func main() {
	fmt.Println("main() started")
	defer defMain()
	foo()
	fmt.Println("main() Done")

}
