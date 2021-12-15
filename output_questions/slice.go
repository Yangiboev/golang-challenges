package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5}
	s = s[:0]
	fmt.Println(s, len(s), cap(s))
}
