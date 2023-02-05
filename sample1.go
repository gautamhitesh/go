package main

import (
	"fmt"
)

func main() {
	// var x int
	// arr := [3]int{3, 5, 2}
	// x, arr = arr[0], arr[1:] //? how to take it end value
	arr1 := [2]int{2, 3}
	arr2 := [...]int{2, 3}
	fmt.Println(arr1 == arr2)
}
