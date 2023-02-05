// swap numbers of a list
package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Before swap ", a)
	a = swap(a)
	fmt.Println("After swap ", a)
	b := []byte{'A', 'B', 'C'}
	c := []byte{'A', 'B', 'C'}
	res := bytes.Compare(b, c)
	fmt.Println(res)
}

// pass by value
func swap(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

//pass by reference
//??
