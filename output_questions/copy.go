package main

import "fmt"

func main() {
	var (
		arr  []int
		arr2 = []int{1, 2, 3}
	)
	num_of_copies := copy(arr, arr2)
	fmt.Printf("n1=%d, arr=%v, arr2=%v", num_of_copies, arr, arr2)
	fmt.Println("arr ==nil", arr==nil)
}
