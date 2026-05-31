package main

import "fmt"
 
type Node struct{
	data int
	right *Node
	left *Node
}

func Binary(){
	root:=&Node{data: 50,
	right: &Node{data:70,
	right: &Node{data:80},
	left:&Node{data:60,},},
	left: &Node{data: 30,
	right: &Node{data: 40},
left: &Node{data: 20},},
}

	
	count := CountNode(root)
	fmt.Println(count)
	Print(root,0)
	fmt.Println(Symmetric(root))

	PreOrder(root)
	PostOrder(root)
	InOrder(root)
	LevelOrder(root)
}

func Print(root *Node,space int){
	if root==nil{
		return
	}
	space+=5
	Print(root.right,space)
	fmt.Println()
	for i:=0;i<space;i++{
		fmt.Print(" ")
	}
	fmt.Println(root.data)
	Print(root.left,space)
}

func InOrder(root *Node){
	if root==nil{
		return
	}
	InOrder(root.left)
	fmt.Print(root.data," ")
	InOrder(root.right)
}

func PreOrder(root *Node){
	if root==nil{
		return
	}
	fmt.Print(root.data," ")
	PreOrder(root.left)
	PreOrder(root.right)
}	

func PostOrder(root *Node){
	if root==nil{
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Print(root.data," ")
}
func LevelOrder(root *Node){
	if root==nil{
		return
	}
	q:=[]*Node{root}
	for len(q)>0{
		curr:=q[0]
		q=q[1:]
		fmt.Println(curr.data," ")
		if curr.left!=nil{
q=append(q, curr.left)
		}
		if curr.right!=nil{
			q=append(q, curr.right)
		}
	}
}

func Symmetric(root *Node)bool{
	if root==nil{
		return true
	}
	return Mirror(root.left,root.right)
}
func Mirror(left,right *Node)bool{
	if left==nil&&right==nil{
		return true
	}
	if left==nil||right==nil{
		return false
	}
	if left.data!=right.data{
		return false
	}
	return Mirror(right.right,left.left)&&Mirror(right.left,left.right)
}

func Insert(root *Node,value int)*Node{
	if root==nil{
		return &Node{data: value}
	}
	if value<root.data{
		root.left=Insert(root.left,value)
	}else if value>root.data{
		root.right=Insert(root.right,value)
	}
	return root
}