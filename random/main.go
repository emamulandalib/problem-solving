package main

import "fmt"

func main() {
	//unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	a := BinarySearch([]int{3, 4, 6, 7}, 9)

	fmt.Println(a)
	// sorted = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
}
