package main

import (
	"fmt"
)

// A type switch is a construct that permits several type assertions in series
func main() {
	checkType(9)
	checkType("Hello")
	checkType(9.87)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	default:
		fmt.Printf("Unknown type %T\n", v)
	case string:
		fmt.Println("Type STRING")
	case int64:
		fmt.Println("Type INTEGER")
	case float64:
		fmt.Println("Type FLOAT")
	}
}
