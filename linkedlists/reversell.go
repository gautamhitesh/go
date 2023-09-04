package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}
type LinkedList struct {
	head *Node
	len  int
}

func main() {
	//create a linked list of 10 elements
	l := new(LinkedList)
	for i := 1; i <= 10; i++ {
		fmt.Println("Added: ", i, l.prepend(i))
	}
	fmt.Println("Size of the Linked List ", l.size())
	l.traverse()
	fmt.Println(l.append(0))
	l.traverse()
	fmt.Println(l.getAt(5).Value)
	fmt.Println(l.addDataBefore(1, -5))
	l.traverse()
	l.reverse()
	fmt.Println(l.size())
	l.traverse()
	l.reverseRec(l.head)
	l.traverse()
	//fmt.Println("Searching -5", l.searchVal(-5))

	fmt.Println(l.searchVal(-5))
	l.deleteVal(-5)
	l.traverse()
	l.deletePos(2)
	l.traverse()
	//fmt.Println(l.getAt(1).Value)

}

func (l *LinkedList) searchVal(data int) int {
	ptr := l.head
	pos := 1
	for ptr != nil {
		if ptr.Value == data {
			return pos
		}
		//fmt.Printf("%d ", ptr.Value)
		ptr = ptr.Next
		pos++
	}
	fmt.Println("Item not found")
	return -1
}

// TO-DO
func (l *LinkedList) deleteVal(data int) {
	pos := l.searchVal(data)
	if pos > 1 {
		fmt.Println("Item found at pos", pos)
		//ptr := l.getAt(pos)
		pptr := l.getAt(pos - 1)
		pptr.Next = pptr.Next.Next
		return
	}
	if pos <= 0 {
		fmt.Println("Value not found")
		return
	}
	if pos == 1 {
		l.head = l.head.Next
	}
}

// TO-DO
func (l *LinkedList) deletePos(pos int) {
	pptr := l.getAt(pos - 1)
	pptr.Next = pptr.Next.Next
	return
}

func (l *LinkedList) reverseRec(ptr *Node) *Node {
	if ptr == nil {
		return nil
	}
	if ptr.Next == nil {
		ptr.Next = nil
		l.head = ptr
		return ptr
	}
	temp := l.reverseRec(ptr.Next)
	temp.Next = ptr
	ptr.Next = nil
	return ptr
}

func (l *LinkedList) reverse() {
	if l.head == nil {
		fmt.Println("Empty linked list")
		return
	}
	ptr := l.head
	var pptr *Node
	var nptr *Node
	for ptr != nil {
		nptr = ptr.Next
		ptr.Next = pptr
		pptr = ptr
		ptr = nptr
	}
	l.head = pptr
}

func (l *LinkedList) prepend(data int) bool {
	temp := &Node{data, nil}
	if l.head == nil {
		fmt.Println("Empty Linked list")
		l.head = temp
		l.len++
		return true
	}
	temp.Next = l.head
	l.head = temp
	l.len++
	return true
}

func (l *LinkedList) append(data int) bool {
	temp := &Node{data, nil}
	if l.head == nil {
		fmt.Println("Empty linked list")
		l.head = temp
		l.len++
		return true
	}
	ptr := l.head
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = temp
	l.len++
	return true
}

func (l *LinkedList) traverse() {
	ptr := l.head
	for ptr != nil {
		fmt.Printf("%d->", ptr.Value)
		ptr = ptr.Next
	}
	fmt.Println()
}

// check edge cases
func (l *LinkedList) addDataBefore(pos, data int) bool {
	if pos == 1 {
		l.prepend(data)
		return true
	}
	ptr := l.getAt(pos)
	prevPtr := l.getAt(pos - 1)
	if prevPtr == nil {
		fmt.Println("Node Position does not exist")
		return false
	}
	if ptr == nil && prevPtr != nil {
		l.append(data)
		return true
	}
	temp := &Node{data, ptr}
	prevPtr.Next = temp
	l.len++
	return true
}

func (l *LinkedList) size() int {
	ptr := l.head
	s := 0
	for ptr != nil {
		s = s + 1
		ptr = ptr.Next
	}
	return s
}

func (l *LinkedList) getAt(pos int) *Node {
	ptr := l.head
	if pos == 1 {
		return l.head
	}
	if pos <= 0 {
		return nil
	}
	if pos > l.len {
		return nil
	}
	for i := 2; i <= pos; i++ {
		ptr = ptr.Next
	}
	return ptr
}
