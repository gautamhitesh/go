package main

import "fmt"

//typer assertion gets the interface value's underlying concrete value

func main() {

	var i interface{} = "Hello"
	s := i.(string)
	fmt.Printf("Value of interface %s\n\n", s)
	t, ok := i.(string)
	fmt.Printf("Is the value fetched: %v\nValue of interface %v\n\n", ok, t)
	u, ok := i.(float64)
	fmt.Printf("Is the value fetched: %v\nValue of interface %v\n\n", ok, u)

}
