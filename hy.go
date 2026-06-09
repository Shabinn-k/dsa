package main

// import "fmt"

// func main(){
//     Graph()
// }

// func Graph(){
//     graph:=make(map[int][]int)
//     AddEdge(graph,0,1)
//     AddEdge(graph,0,2)
//     AddEdge(graph,1,3)
//     AddEdge(graph,2,4)
//     AddEdge(graph,4,5)
//     AddEdge(graph,4,6)
    
//     values:=map[int]int{
//         0:10,
//         1:20,
//         3:30,
//         4:40,
//         5:50,
//         6:60,
//     }
    
//     //normal graph
//     fmt.Println(graph)
    
//     //sum
//     fmt.Println("Total of values :",Sum(0,graph,make(map[int]bool),values))
    
//     //leaf nodes
//     fmt.Println("Total leaf nodes :",Leaf(graph))
    
//     //degree 
//     fmt.Println("Total Edges :",Degree(graph))
    
//     //individual
//     fmt.Println("Edges of 4 :",IE(graph,4))
    
//     //reachable
//     fmt.Println(Reachable(graph,1,make(map[int]bool)))
    
//     //total nodes
//     fmt.Println("Total nodes :",Total(graph))
    
//     //cycle 
//     fmt.Println("If cycle :",IsCycle(graph,0,make(map[int]bool),-1))
    
//     //max element
//     fmt.Println(Max(values))
    
//     //max depth
//     fmt.Println(Depth(graph,0,make(map[int]bool)))
    
//     //path exist
//     fmt.Println(Path(graph,0,1,make(map[int]bool)))
    
//     //bfs
//     BFS(0,graph)
    
//     fmt.Println()
//     //dfs
//     DFS(0,graph,make(map[int]bool))
    
// }

// func AddEdge(gr map[int][]int,u,v int){
//         gr[u]=append(gr[u],v)
//         gr[v]=append(gr[v],u)
// }

// // sum
// func Sum(node int,gr map[int][]int,vis map[int]bool,val map[int]int)int{
//     vis[node]=true
//     c:=val[node]
//     for _,n:=range gr[node]{
//         if !vis[n]{
//             c+=Sum(n,gr,vis,val)
//         }
//     }
//     return c
// }

// //leaf node 
// func Leaf(gr map[int][]int)int{
//     c:=0
//     for _,n:=range gr{
//         if len(n)==1{
//         c++    
//         }
//     }
//     return c
// }

// //degree
// func Degree(gr map[int][]int)int{
//     c:=0
//     for _,n:=range gr{
//         c+=len(n)
//     }
//     return c/2
// }

// //individual edges
// func IE(gr map[int][]int,node int)int{
//     return len(gr[node])
// }

// //total nodes
// func Total(gr map[int][]int)int{
//     return len(gr)    
// }

// //reachable 
// func Reachable(gr map[int][]int,node int,vis map[int]bool)int{
//     vis[node]=true
//     c:=1
//     for _,n:=range gr[node]{
//         if !vis[n]{
//             c+=Reachable(gr,n,vis)
//         }
//     }
//     return c
// }
// //bfs
// func BFS(start int,gr map[int][]int){
//     vis:=make(map[int]bool)
//     que:=[]int{start}
//     vis[start]=true
//     for len(que)>0{
//         node:=que[0]
//         que=que[1:]
//         fmt.Print(node," ")
//         for _,n:=range gr[node]{
//             if !vis[n]{
//                 vis[n]=true
//             que=append(que,n)
//             }
//         }
//     }
// }
// //dfs
// func DFS(node int,gr map[int][]int,vis map[int]bool){
//     if vis[node]{
//         return
//     }
//     vis[node]=true
//     fmt.Print(node," ")
//     for _,n:=range gr[node]{
//         if !vis[n]{
//             DFS(n,gr,vis)
//         }
//     }
// }
// //cycle
// func IsCycle(gr map[int][]int,node int,vis map[int]bool,parent int)bool{
//     vis[node]=true
//     for _,n:=range gr[node]{
//         if !vis[n]{
//             if IsCycle(gr,n,vis,node){
//                 return true
//             }
//         }else if n!=parent{
//             return true
//         }
//     }
//     return false
    
// }
// //maximum elemnt
// func Max(val map[int]int)(int,int){
//     maxNode:=-1
//     maxValue:=-1
//     for n,v:=range val{
//         if v>maxValue{
//             maxValue=v
//             maxNode=n
//         }
//     }
//     return maxNode,maxValue
// }

// //depth
// func Depth(gr map[int][]int,node int,vis map[int]bool)int{
//     vis[node]=true
//     max:=0
//     for _,n:=range gr[node]{
//         if !vis[n]{
//             d:=Depth(gr,n,vis)
//             if d>max{
//                 max=d
//             }
//         }
//     }
//     return max+1
// }

// //path
// func Path(gr map[int][]int,start,end int,vis map[int]bool)bool{
//     if start==end{
//         return true
//     }
//     vis[start]=true
//     for _,n:=range gr[start]{
//         if !vis[n]{
//             if Path(gr,n,end,vis){
//                 return true
//             }
//         }
//     }
//     return false
// }


// //sorting
// //qs
// // func QuickSort()
// //ms
// //ss
// //is
// //bs
// //hs

