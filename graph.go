package main

import "fmt"

func Graph() {
	gr := make(map[string][]string)
	gr["A"] = append(gr["A"], "B")
	gr["A"] = append(gr["A"], "C")
	gr["B"] = append(gr["B"], "A")
	gr["C"] = append(gr["C"], "A")
	fmt.Println(gr)
}

func AddEdge(graph map[int][]int,u,v int){
	graph[u]=append(graph[u], v)
	graph[v]=append(graph[v], u)
}

func GR(){
	gr:=make(map[int][]int)
		AddEdge(gr,0,1)
		AddEdge(gr,0,2)
		AddEdge(gr,1,3)
		AddEdge(gr,1,4)
		AddEdge(gr,2,5)	
		AddEdge(gr,2,6)
	fmt.Println(gr)
	visited:=make(map[int]bool)
	DFS(0,gr,visited)
	fmt.Println()
	BFS(0,gr)
}
func BFS(start int,graph map[int][]int){
	visited:=make(map[int]bool)
	que:=[]int{start}
	visited[start]=true
	for len(que)>0{
		node:=que[0]
		que=que[1:]
		fmt.Print(node," ")
		for _,neighbour:=range graph[node]{
			if !visited[neighbour]{
				visited[neighbour]=true
				que = append(que, neighbour)
			}
		}
	}
}
func DFS(node int,graph map[int][]int,visited map[int]bool){
		if visited[node]{
			return 
		}
		visited[node]=true
		fmt.Print(node," ")
		for _,neighbour:=range graph[node]{
			DFS(neighbour,graph,visited)
		}
}

