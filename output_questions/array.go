package main

import "fmt"

func makeSquare(arr [10]int) {
	for index, elem := range arr {
		arr[index] = elem * elem
	}
}
func main() {
	var arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	makeSquare(arr)

	fmt.Println(arr)

}
