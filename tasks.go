package main
import "fmt"

type Nodes struct{
    data int
    left *Node
    right *Node
}
// binary tree
func binary(){
    root:=&Node{data:100,
    right:&Node{data:150,
    right:&Node{data:180,},
    left:&Node{data:130,},
    },
    left:&Node{data:50,
    right:&Node{data:80,},
    left:&Node{data:30,},
},
}

insert(root,10)
Print(root,0)
fmt.Println(Search(root,10))
fmt.Println(sym(root))
fmt.Println(CountN(root))
fmt.Println(Height(root))
fmt.Println(CountL(root))
Inorder(root)
fmt.Println()
Postorder(root)
fmt.Println()
Preorder(root)
fmt.Println()
Levelorder(root)
fmt.Println()
max,min:=high(root)
fmt.Println("max :",max)
fmt.Println("min :",min)
fmt.Println(Sum(root))
fmt.Println(exist(root,100))
Delete(root,10)
Print(root,0)
    
}
// print
func Prints(root *Node,space int){
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
// insert
func insert(root *Node,val int)*Node{
    if root==nil{
        return &Node{data:val}
    }
    if val>root.data{
        root.right=insert(root.right,val)
    }else if val<root.data{
        root.left=insert(root.left,val)
    }
    return root
}
// search bst
func Search(root *Node,val int)bool{
    if root==nil{
        return false
    }
    if root.data==val{
        return true
    }
    if val<root.data{
        return Search(root.left,val)
    }
    return Search(root.right,val)
}

// symmetr
func sym(root *Node)bool{
    if root==nil{
        return true
    }
    return Mirror(root.right,root.left)
}
func Mirr(right,left *Node)bool{
    if left==nil&&right==nil{
        return true
    }
    if right==nil||left==nil{
        return false
    }
    if left.data!=right.data{
        return false
    }
    return Mirror(right.right,left.left)&&Mirror(right.left,left.right)
}

// node
func CountN(root *Node)int{
    if root==nil{
        return 0
    }
    return 1+CountN(root.left)+CountN(root.right)
}

// height
func Height(root *Node)int{
   if root==nil{
       return 0
   }
   l:=Height(root.left)
   r:=Height(root.right)
   if l>r{
       return l+1
   }
   return r+1
}

// leafnode
func CountL(root *Node)int{
    if root==nil{
        return 0
    }
    if root.left==nil&&root.right==nil{
        return 1
    }
    return CountL(root.left)+CountL(root.right)
}

// traversals
func Inorder(root *Node){
    if root==nil{
        return 
    }
    Inorder(root.left)
    fmt.Print(root.data," ")
    Inorder(root.right)
}
func Postorder(root *Node){
    if root==nil{
        return 
    }
    Postorder(root.left)
    Postorder(root.right)
    fmt.Print(root.data," ")
}
func Preorder(root *Node){
    if root==nil{
        return
    }
    fmt.Print(root.data," ")
    Preorder(root.left)
    Preorder(root.right)
}
func Levelorder(root *Node){
    if root==nil{
        return 
    }
    q:=[]*Node{root}
    for len(q)>0{
        curr:=q[0]
        q=q[1:]
        fmt.Print(curr.data," ")
        if curr.left!=nil{
            q=append(q,curr.left)
        }
        if curr.right!=nil{
            q=append(q,curr.right)
        }
    }
}

// max and min
func high(root *Node)(int,int){
    if root==nil{
        return -1,-1
    }
    max:=root.data
    min:=root.data
    if root.right!=nil{
        rmax,rmin:=high(root.right)
        if rmax>max{
            max=rmax
        }
        if rmin<min{
            min=rmin
        }
    }
    if root.left!=nil{
        lmax,lmin:=high(root.left)
        if lmax>max{
            max=lmax
        }
        if lmin<min{
            min=lmin
        }
    }
    return max,min
}

// sum
func Sum(root *Node)int{
    if root==nil{
        return 0
    }
    return root.data+Sum(root.right)+Sum(root.left)
}

// exist
func exist(root *Node,val int)bool{
    if root==nil{
        return false
    }
    if root.data==val{
        return true
    }
    return exist(root.left,val)||exist(root.right,val)
}

// delete
func Findmin(root *Node)*Node{
    for root.left!=nil{
        root=root.left
    }
    return root
}
func Delete(root *Node,val int)*Node{
    if root==nil{
        return nil
    }
    if val<root.data{
        root.left=Delete(root.left,val)
    }else if val>root.data{
        root.right=Delete(root.right,val)
    }else{
        if root.left==nil&&root.right==nil{
            return nil
        }
        if root.left==nil{
            return root.left
        }
        if root.right==nil{
            return root.right
        }
        min:=Findmin(root.right)
        root.data=min.data
		root.right = Delete(root.right, min.data)
    }
    return root
}