package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func Binary() {
	root := &Node{
		data: 10,
		left: &Node{
			data: 20,
			left: &Node{
				data: 40,
			},
		},
		right: &Node{data: 30},
	}

	count := CountNode(root)
	fmt.Println(count)
	PrintTree(root,0)
}

func PrintTree(root *Node, space int) {
	if root == nil {
		return
	}
	space += 5
	PrintTree(root.right, space)
	fmt.Println()
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(root.data)
	PrintTree(root.left, space)
}
