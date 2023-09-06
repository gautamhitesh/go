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

func serialize(ptr *Node) {
	if ptr == nil {
		fmt.Printf("%d ", -1)
		return
	}
	fmt.Printf("%d ", ptr.data)
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

func main() {
	var root *Node
	//r := utils.GenerateRandomArray(10)
	//for i := 0; i < 10; i++ {
	//	insert(r[i], &root)
	//}
	insert(12, &root)
	insert(13, &root)
	inorder(root)
	fmt.Println()
	serialize(root)
}
