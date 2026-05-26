package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func Binary() {
	root := Node{data: 10}
	root.left = &Node{data: 20}
	root.right = &Node{data: 30}

	fmt.Println(root.data)
	fmt.Println(root.left.data)
	fmt.Println(root.right.data)
}