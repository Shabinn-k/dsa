package main

import "fmt"

type MinHeap struct {
	arr []int
}

func (h *MinHeap) InsertH(val int) {
	h.arr = append(h.arr, val)
	i := len(h.arr) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.arr[parent] <= h.arr[i] {
			break
		}
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		i = parent
	}
}

func (h *MinHeap)Heapify(i int){
smallest:=i
left:=2*i+1
right:=2*i+2
if left<len(h.arr)&&h.arr[left]<h.arr[smallest]{
	smallest=left
}
if right<len(h.arr)&&h.arr[right]<h.arr[smallest]{
	smallest=right
}
if smallest!=i

}
func (h *MinHeap)MinExtract()int{
	if len(h.arr)==0{
		return -1
	}
	min:=h.arr[0]
	last:=len(h.arr)-1
	h.arr[0]=h.arr[last]
	h.arr=h.arr[:last]

	h.Heapify()
}
func Heap() {
	h := MinHeap{}
	h.InsertH(10)
	h.InsertH(5)
	h.InsertH(20)
	h.InsertH(2)

	fmt.Println(h.arr)
}