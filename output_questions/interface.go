package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Rect2 struct {
	width  float64
	height float64
}

func (r Rect2) Area() float64 {
	return r.width * r.height
}
func (r Rect2) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var s Shape
	fmt.Printf("type of s is %T\n", s)
	s = Rect{5.0, 3.0}
	r := Rect{5.0, 4.0}
	r2 := Rect2{5.0, 4.0}
	fmt.Println(s == r)
	fmt.Println(r2 == r)
	fmt.Printf("type of s is %T\n", s)

	area := s.Area()
	fmt.Println(area)
}
