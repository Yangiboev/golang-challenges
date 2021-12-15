package main

import "fmt"

// Question: exit outer loop
func main() {
	// outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(j)
			// break outer
		}
	}
}
