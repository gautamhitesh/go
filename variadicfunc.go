package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(joinString("Hello", "World"))
	fmt.Println(joinString("Welcome", "to", "join", "by", "Variadic", "function"))
}

func joinString(astring ...string) string {
	return strings.Join(astring, "-")
}
