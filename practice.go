package main

// import "fmt"

// func main(){
//     // Anagram()
//     binary()
// }
// type Node struct{
//     data int
//     left *Node
//     right *Node
// }

// func binary(){
//     root:=&Node{data:100,
//     right:&Node{data:150,
//     right:&Node{data:180,},
//     left:&Node{data:130,},
// },
// left:&Node{data:50,
// right:&Node{data:80,},
// left:&Node{data:30,},
// },
// }

// insert(root,500)
// insert(root,510)
// insert(root,100)

// Print(root,0)
// fmt.Println(Count(root))
// fmt.Println(LCount(root))
// fmt.Println(Height(root))
// fmt.Println(Element(root))
// fmt.Println(Sum(root))
// fmt.Println(Exist(root,101))
// }

// func Sum(root *Node)int{
//     if root==nil{
//         return 0
//     }
//     return root.data+Sum(root.left)+Sum(root.right)
// }
// func Exist(root *Node,val int)bool{
//     if root==nil{
//         return false
//     }
//     if root.data==val{
//         return true
//     }
//     return Exist(root.left,val)||Exist(root.right,val)
// }
// //largest node
// func Element(root *Node)(int, int){
//     if root==nil{
//         return -1,-1
//     }
//     max:=root.data
//     min:=root.data
    
//     if root.right!=nil{
//         rmax,rmin:=Element(root.right)
//         if rmax>max{
//             max=rmax
//         }
//         if rmin<min{
//             min=rmin
//         }
//     }
//     if root.left!=nil{
//         lmax,lmin:=Element(root.left)
//         if lmax>max{
//             max=lmax
//         }
//         if lmin<min{
//             min=lmin
//         }
//     }
//     return max,min
// }

// //height of tree
// func Height(root *Node)int{
//   if root ==nil{
//       return 0
//   }
//   l:=Height(root.left)
//   r:=Height(root.right)
//     if l>r{
//         return l+1
//     }
//     return r+1
    
// }

// // leaf node total
// func LCount(root *Node)int{
//     if root==nil{
//         return 0
//     }
//     if root.left==nil&&root.right==nil{
//         return 1
//     }
//     return LCount(root.left)+LCount(root.right)
// }
// func insert(root *Node,value int)*Node{
//     if root==nil{
//         return &Node{data:value}
//     }
//     if value<root.data{
//         root.left=insert(root.left,value)
//     }else if value>root.data{
//         root.right=insert(root.right,value)
//     }
//     return root
// }

// func Print(root *Node,space int){
//     if root==nil{
//         return 
//     }
//     space+=5
//     Print(root.right,space)
//     fmt.Println()
//     for i:=0;i<space;i++{
//         fmt.Print(" ")
//     }
//     fmt.Println(root.data)
//     Print(root.left,space)
// }

// func Count(root *Node)int{
//     if root==nil{
//         return 0
//     }
//     return 1+Count(root.left)+Count(root.right)
// }

// // Check if two strings are anagrams
// func Anagram(){
//     a:="tar"
//     b:="rat"
//     m:=make(map[rune]int)
//     if len(a)!=len(b){
//         fmt.Println(false)
//         return
//     }
//     for _,v:=range a{
//         m[v]++
//     }
//     for _,v:=range b{
//         m[v]--
//     }
//     for _,v:=range m{
//         if v!=0{
//         fmt.Println("false")
//     return 
//         }
//     fmt.Println("true")
//     break
//     }
// }
// // Remove duplicate characters
// func duplicates(){
//     arr:=[]int{1,2,3,4,5,6,1,2,4,7}
//     m:=make(map[int]bool)
//     d:=[]int{}
//     for _,v:=range arr{
//         if m[v]{
//             fmt.Println(v)
//         }else{
//             d=append(d,v)
//         }
//         m[v]=true
//     }
//     fmt.Println(d)
// }

// // Find first non-repeating character
// func Repeat(){
//     m:=make(map[rune]bool)
//     n:=make(map[rune]int)
//     str:="golang"
//     for _,v:=range str{
//         if m[v]{
//             fmt.Println("first repeating :",string(v))
//             break
//         }
//         m[v]=true
//     }
//   for _, v := range str {
// 		n[v]++
// 	}
// 	for _,v:=range str{
// 		if n[v]==1{
// 			fmt.Println("first non-repeating :",string(v))
// 			break
// 		}
// 	}
// }

// // Reverse a string
// func Reverse(){
//     str:="shabin"
//     r:=[]rune(str)
//     f:=0
//     l:=len(r)-1
//     for f<l{
//          r[f],r[l]=r[l],r[f]
//     f++
//     l--
// }
//     fmt.Println(string(r))
// }

// // Find the Kth largest element
// func Kelement(arr []int,k int)int{
//     if k>=6{
//         return 0
//     }
//     n:=len(arr)
//     for i:=0;i<n;i++{
//         for j:=0;j<n-i-1;j++{
//             if arr[j]<arr[j+1]{
//                 arr[j],arr[j+1]=arr[j+1],arr[j]
//             }
//         }
//     }
//     return arr[k-1]
// }

// // Sort an array in descending order
// func Descend(arr []int)[]int{
//      n:=len(arr)
//     for i:=0;i<n;i++{
//         for j:=0;j<n-i-1;j++{
//             if arr[j]<arr[j+1]{
//                 arr[j],arr[j+1]=arr[j+1],arr[j]
//             }
//         }
//     }
//     return arr
// }

// // Implement Merge Sort
// func MS(arr []int)[]int{
//     if len(arr)<=1{return arr}
//     mid:=len(arr)/2
//     right:=MS(arr[mid:])
//     left:=MS(arr[:mid])
//     return merge(left,right)
// }
// func merge(left,right []int)[]int{
//     res:=[]int{}
//     i,j:=0,0
//     for i<len(left)&&j<len(right){
//         if left[i]<=right[j]{
//             res=append(res,left[i])
//             i++
//         }else{
//             res=append(res,right[j])
//             j++
//         }
//     }
//     if i<len(left){
//         res=append(res,left[i:]...)
//         i++
//     }
//     if j<len(right){
//         res=append(res,right[j:]...)
//         j++
//     }
//     return res
// }

// // Implement Bubble Sort
// func bubble(arr []int)[]int{
//     n:=len(arr)
//     for i:=0;i<n;i++{
//         for j:=0;j<n-i-1;j++{
//             if arr[j]>arr[j+1]{
//                 arr[j],arr[j+1]=arr[j+1],arr[j]
//             }
//         }
//     }
//     return arr
// }

// // Implement Selection Sort
// func selection(arr []int)[]int{
//     n:=len(arr)
//     for i:=0;i<n-1;i++{
//         min:=i
//         for j:=i+1;j<n;j++{
//             if arr[j]<arr[min]{
//                 min=j
//             }
//         }
//         arr[i],arr[min]=arr[min],arr[i]
//     }
//     return arr
    
// }

// // Implement Insertion  Sort
// func Insertion(arr []int)[]int{
//     for i:=0;i<len(arr);i++{
//         key:=arr[i]
//         j:=i-1
//         for j>=0&&arr[j]>key{
//             arr[j+1]=arr[j]
//             j--
//         }
//         arr[j+1]=key
//     }
//     return arr
// }
