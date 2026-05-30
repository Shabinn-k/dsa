package main

func CountNode(root *Node)int{
	if root==nil{
		return 0
	}
	return 1+CountNode(root.left)+CountNode(root.right)
}