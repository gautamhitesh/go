package main

import (
	. "dsa/utils"
	"fmt"
	"io"
	"os"
)

type BTree struct {
	root  *Node
	depth int
}
type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	t := new(BTree)
	r := GenerateRandomArray(10)
	fmt.Println(r)
	for i := 0; i < 10; i++ {
		t.insert(r[i])
	}
	inorder(t.root)
	fmt.Println()
	preorder(t.root)
	fmt.Println()
	postorder(t.root)
	fmt.Println()
	t.insert(11)
	t.insert(-1)
	inorder(t.root)
	fmt.Println()
	print(os.Stdout, t.root, 0, 'M')
	fmt.Println(maxDepth(t.root))
}

func print(w io.Writer, n *Node, ns int, ch rune) {
	if n == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, n.data)
	print(w, n.left, ns+2, 'L')
	print(w, n.right, ns+2, 'R')
}

func (t *BTree) insert(data int) *BTree {
	temp := &Node{data, nil, nil}
	if t.root == nil {
		t.root = temp
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &Node{data, nil, nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data, nil, nil}
		} else {
			n.right.insert(data)
		}
	}
}

func inorder(ptr *Node) {
	if ptr == nil {
		return
	}
	inorder(ptr.left)
	fmt.Printf("%d ", ptr.data)
	inorder(ptr.right)
	return
}

func postorder(ptr *Node) {
	if ptr == nil {
		return
	}
	postorder(ptr.left)
	postorder(ptr.right)
	fmt.Printf("%d ", ptr.data)
}
func preorder(ptr *Node) {
	if ptr == nil {
		return
	}
	fmt.Printf("%d ", ptr.data)
	preorder(ptr.left)
	preorder(ptr.right)
}

func maxDepth(ptr *Node) int {
	if ptr == nil {
		return 0
	}
	lDepth := maxDepth(ptr.left)
	rDepth := maxDepth(ptr.right)
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}
