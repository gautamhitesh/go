package main

import "fmt"

func main() {
	num := 7
	fmt.Printf("%vth number in fibonacci sequence %v", num, fibo(num))
}

func fibo(num int) int {
	a, b, c := 0, 1, 0
	fmt.Println(0, 1)
	if num == 1 {
		return 0
	}
	if num == 2 {
		return 1
	}
	for i := 0; i < num-1; i++ {
		c = a + b
		a = b
		b = c
		fmt.Println(c)
	}
	return c
}
func recusiveFibo(num int) int {
	if num < 2 {
		return num
	} else {
		return recusiveFibo(num-1) + recusiveFibo(num-2)
	}
}
