package main
// // import "fmt"

// func mains() {
// // arr:=[]int{3,2,5,1,7,4,6}
// // BB:=BubbleSort(arr)
// // fmt.Println(BB)

// // IN:=Insertional(arr)
// // fmt.Println(IN)

// // SL:=Selection(arr)
// // fmt.Println(SL)

// // MS:=MS(arr)
// // fmt.Println(MS)

// // QS(arr,0,len(arr)-1)
// // fmt.Println(arr)
    
//     // HS(arr)
//     // fmt.Println(arr)
    
//     Graph()
// }
// // func PathExist(start,end int,gr map[int][]int,vis map[int]bool)bool{
// //     if start==end{
// //         return true
// //     }
// //     vis[start]=true
// //     for _,n:=range gr[start]{
// //         if !vis[n]{
// //             if PathExist(n,end,gr,vis){
// //                 return true
// //             }
// //         }
// //     }
// //     return false
// // }
// // func BFS(start int ,gr map[int][]int){
// //     vis:=make(map[int]bool)
// //     que:=[]int{start}
// //     vis[start]=true
// //     for len(que)>0{
// //         node:=que[0]
// //         que=que[1:]
// //         fmt.Print(node," ")
// //         for _,n:=range gr[node]{
// //             if !vis[n]{
// //                 vis[n]=true
// //                 que=append(que,n)
// //             }
// //         }
// //     }
// // }
// // func DFS(node int,gr map[int][]int,vis map[int]bool){
// //     if vis[node]{
// //         return
// //     }
// //     vis[node]=true
// //     fmt.Print(node," ")
// //     for _,n:=range gr[node]{
// //         DFS(n,gr,vis)
// //     }
// // }
// // func AddEdge(gr map[int][]int,u,v int){
// //     gr[u]=append(gr[u],v)
// //     gr[v]=append(gr[v],u)
// // }
// // func Graph(){
// //     gr:=make(map[int][]int)
//     // AddEdge(gr,0,1)
//     // AddEdge(gr,0,2)
//     // AddEdge(gr,1,3)
//     // AddEdge(gr,1,4)
//     // AddEdge(gr,2,5)
//     // AddEdge(gr,2,6)
//     // AddEdge(gr,5,7)
//     // AddEdge(gr,7,8)
    
// // fmt.Println(gr)
// //     BFS(0,gr)
// //     fmt.Println()
// //     vis:=make(map[int]bool)
// //     DFS(0,gr,vis)
// //     fmt.Println()
// //     v:=make(map[int]bool)
// //     fmt.Println(PathExist(1,9,gr,v))
// // }

// func Heapify(arr []int,n,i int){
//     large:=i
//     left:=2*i+1
//     right:=2*i+2
//     if left<n&&arr[left]>arr[large]{
//         large=left
//     }
//     if right<n&&arr[right]>arr[large]{
//         large=right
//     }
//     if large!=i{
//         arr[i],arr[large]=arr[large],arr[i]
//         Heapify(arr,n,large)
//     }
// }
// func HS(arr []int){
//  n:=len(arr)
//  for i:=n/2-1;i>=0;i--{
//      Heapify(arr,n,i)
//  }
//  for i:=n-1;i>0;i--{
//      arr[0],arr[i]=arr[i],arr[0]
//      Heapify(arr,i,0)
//  }
// }

// func BubbleSort(arr []int)[]int{
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

// func Insertional(arr []int)[]int{
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

// // func Selection(arr []int)[]int{
// //     n:=len(arr)
// //     for i:=0;i<n;i++{
// //         min:=i
// //         for j:=i+1;j<n-1;j++{
// //             if arr[j]<arr[min]{
// //                 min=j
// //             }
// //         }
// //         arr[i],arr[min]=arr[min],arr[i]
// //     }
// //     return arr
// // }

// func MS(arr []int)[]int{
//     if len(arr)<=1{
//         return arr
//     }
//     mid:=len(arr)/2
//     left:=MS(arr[:mid])
//     right:=MS(arr[mid:])
    
//     return Merge(left,right)
// }
// func Merge(left,right []int)[]int{
//     res:=[]int{}
//     i,j:=0,0
//     for i<len(left)&&j<len(right){
//         if left[i]<=right[j]{
//             res=append(res,left[i])
//             i++
//         }else{
//             res=append(res,right[j])
//                 j++        
//         }
//     }
//     if i<len(left){
//         res=append(res,left[i:]...)
//     }
//     if j<len(right){
//         res=append(res,right[j:]...)
//     }
//     return res
// }

// func QS(arr []int,low,high int){
//     if low<high{
//         p:=Part(arr,low,high)
//         QS(arr,low,p-1)
//         QS(arr,p+1,high)
//     }
// }
// func Part(arr []int,low,high int)int{
//     pivot:=arr[high]
//     i:=low-1
//     for j:=low;j<high;j++{
//         if arr[j]<pivot{
//             i++
//             arr[i],arr[j]=arr[j],arr[i]
//         }
//     }
//     arr[i+1],arr[high]=arr[high],arr[i+1]
//     return i+1
// }

// package main
// import "fmt"

// func main(){
//     Graph()
// }
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
// func Leaf(gr map[int][]int)int{
//     c:=0
//     for _,n:=range gr{
//         if len(n)==1{
//             c++
//         }
//     }
//     return c
// }
// func Sum(node int,gr map[int][]int,vis map[int]bool,val map[int]int)int{
//     vis[node]=true
//     sum:=val[node]
//     for  _,n:=range gr[node]{
//         if !vis[n]{
//             sum+=Sum(n,gr,vis,val)
//         }
//     }
//     return sum
// }

// func TN(gr map[int][]int)int{
//     return len(gr) 
// }

// func Edge(gr map[int][]int)int{
//         c:=0
//         for _,n:=range gr{
//         c+=len(n)
//         }
//         return c/2
// }
// func SE(gr map[int][]int,node int)int{
//     return len(gr[node])
// }
// func HasCycle(gr map[int][]int,node int,vis map[int]bool,p int)bool{
//     vis[node]=true
//     for _,n:=range gr[node]{
//         if !vis[n]{
//             if HasCycle(gr,n,vis,node){
//                 return true
//             }
//         }else if n!=p{
//             return true
//         }
//     }
//     return false
// }
// func AddEdge(gr map[int][]int,u,v int){
//     gr[u]=append(gr[u],v)
//     gr[v]=append(gr[v],u)
// }
// func Graph(){
//     gr:=make(map[int][]int)
//   AddEdge(gr, 0, 1)
//     AddEdge(gr,0,2)
//     AddEdge(gr,1,3)
//     AddEdge(gr,1,4)
//     AddEdge(gr,2,5)
//     AddEdge(gr,2,6)
//     AddEdge(gr,5,7)
//     AddEdge(gr,7,8)
    
// 	val := map[int]int{
// 		0: 10,
// 		1: 20,
// 		2: 30,
// 		3: 40,
// 		4: 50,
// 	}
// 	v:=make(map[int]bool)
//     total:=Sum(0,gr,v,val)
//     fmt.Println(total)
// fmt.Println(HasCycle(gr,0,v,-1))
// fmt.Println(Edge(gr))
// fmt.Println(SE(gr,0))
// fmt.Println(TN(gr))
// fmt.Println(Reachable(gr,0,make(map[int]bool)))
// fmt.Println(Leaf(gr))
// }