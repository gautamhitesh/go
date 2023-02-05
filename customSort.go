package main

import (
	"fmt"
	"runtime"
	"sort"
)

type Human struct {
	name string
	age  int
}

type AgeFactor []Human

func (a AgeFactor) Len() int {
	return len(a)
}

func (a AgeFactor) Less(i, j int) bool {
	return a[i].age < a[j].age
}
func (a AgeFactor) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	audience := []Human{
		{"Alice", 35},
		{"Bob", 42},
		{"Nancy", 31},
	}
	fmt.Println("Before sort ", audience)
	sort.Sort(AgeFactor(audience))
	fmt.Println(runtime.NumCPU())
	fmt.Println(audience)
}
