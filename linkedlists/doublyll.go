package main

import "fmt"

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}
type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (l *LinkedList) prepend(data int) bool {
	temp := &Node{data, nil, nil}
	if l.head == nil {
		l.head = temp
		l.tail = temp
		l.length++
		return true
	}
	ptr := l.head
	temp.next = ptr
	ptr.prev = temp
	l.head = temp
	l.length++
	return true
}

func (l *LinkedList) append(data int) bool {
	temp := &Node{data, nil, nil}
	ptr := l.head
	if ptr == nil {
		l.head = temp
		l.tail = temp
		l.length++
		return true
	}
	for ptr.next != nil {
		ptr = ptr.next
	}
	temp.prev = ptr
	ptr.next = temp
	l.tail = temp
	l.length++
	return true
}

func (l *LinkedList) print() {
	ptr := l.head
	for ptr != nil {
		fmt.Printf("%d-->", ptr.value)
		ptr = ptr.next
	}
	fmt.Println()
	return
}
func (l *LinkedList) printBackwards() {
	ptr := l.tail
	for ptr != nil {
		fmt.Printf("%d-->", ptr.value)
		ptr = ptr.prev
	}
	fmt.Println()
	return
}

// To DO
func (l *LinkedList) reverseDll() {
}

func main() {
	var l LinkedList
	for i := 0; i < 10; i++ {
		l.prepend(i + 1)
	}
	l.print()
	l.printBackwards()
	fmt.Println("Size of DLL ", l.length)
	l.append(0)
	l.print()
	l.printBackwards()
	fmt.Println("Size of DLL ", l.length)
	l.reverseDll()
	l.print()
	l.printBackwards()
	fmt.Println("Values ", l.head.value, l.head.next.value, l.tail.prev.value, l.tail.value)
}
