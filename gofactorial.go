package main

import "fmt"

func main() {
	fmt.Println("Factorial of number 6 ", factorial(6))
}

func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
