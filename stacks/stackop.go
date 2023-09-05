package main

import "fmt"

//LIFO

type Stack struct {
	head *Node
	len  int
}

type Node struct {
	data int
	next *Node
}

func main() {
	stack := new(Stack)
	for i := 0; i < 10; i++ {
		stack.push(i + 1)
	}
	stack.print()
	stack.pop()
	stack.pop()
	stack.print()
}
func (stack *Stack) print() {
	ptr := stack.head
	for ptr != nil {
		fmt.Printf("%d ", ptr.data)
		ptr = ptr.next
	}
	fmt.Println()
	return
}

func (stack *Stack) push(data int) {
	temp := &Node{data, stack.head}
	stack.head = temp
	stack.len++
	return
}
func (stack *Stack) pop() int {
	if stack.head == nil {
		fmt.Println("Empty stack")
		return -1
	}
	temp := stack.head.data
	stack.head = stack.head.next
	stack.len--
	return temp
}
