package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

func main() {
	var s Shape
	fmt.Println("value of s is ", s)
	fmt.Printf("type  of s is %T\n", s)
}
