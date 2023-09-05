package main

import "fmt"

// fifo
type Queue struct {
	head *Node
	len  int
}
type Node struct {
	data int
	next *Node
}

func main() {
	q := new(Queue)
	for i := 0; i < 10; i++ {
		q.push(i + 1)
	}
	q.print()
	q.pop()
	q.print()
	q.pop()
	q.print()

}
func (q *Queue) print() {
	ptr := q.head
	for ptr != nil {
		fmt.Printf("%d ", ptr.data)
		ptr = ptr.next
	}
	fmt.Println()
	return
}

func (q *Queue) push(data int) {
	temp := &Node{data, nil}
	if q.head == nil {

		q.head = temp
		q.len++
		return
	}
	ptr := q.head
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = temp
	q.len++
	return
}
func (q *Queue) pop() int {
	if q.head == nil {
		fmt.Println("Empty queue")
		return -1
	}
	temp := q.head.data
	q.head = q.head.next
	q.len--
	return temp
}
