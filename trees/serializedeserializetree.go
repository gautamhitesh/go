package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func insert(data int, ptr **Node) {
	if *ptr == nil {
		*ptr = &Node{data, nil, nil}
		return
	}
	if data <= (*ptr).data {
		insert(data, &(*ptr).left)
	}
	if data > (*ptr).data {
		insert(data, &(*ptr).right)
	}
	return
}

var serialized []int

func serialize(ptr *Node) {
	if ptr == nil {
		fmt.Printf("%d ", -1)
		return
	}
	fmt.Printf("%d ", ptr.data)
	serialized = append(serialized, ptr.data)
	serialize(ptr.left)
	serialize(ptr.right)
}

func inorder(ptr *Node) {
	if ptr == nil {
		//fmt.Println("empty")
		return
	}
	inorder(ptr.left)
	fmt.Printf("%d ", ptr.data)
	inorder(ptr.right)
}

func deserialize(ptr **Node, data int) {
	if data == -1 {
		return
	}
	if *ptr == nil {
		//fmt.Println("Nil node")
		*ptr = &Node{data, nil, nil}
		return
	}
	if data < (*ptr).data {
		//fmt.Println("going left")
		deserialize(&(*ptr).left, data)
	}
	if data > (*ptr).data {
		//fmt.Println("going right")
		deserialize(&(*ptr).right, data)
	}
}

func main() {
	var root *Node
	//r := utils.GenerateRandomArray(10)
	//for i := 0; i < 10; i++ {
	//	insert(r[i], &root)
	//}
	insert(20, &root)
	insert(8, &root)
	insert(4, &root)
	insert(12, &root)
	insert(10, &root)
	insert(14, &root)
	inorder(root)
	fmt.Println()
	serialize(root)
	fmt.Println()
	fmt.Println(serialized)
	var ds *Node
	for i := 0; i < 6; i++ {
		deserialize(&ds, serialized[i])
	}
	fmt.Println()
	serialized = nil
	serialize(root)
	fmt.Println()
	fmt.Println(serialized)
}
