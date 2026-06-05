package main

func MergeSort(arr []int)[]int{
	if len(arr)<=1{return arr}
	mid:=len(arr)/2
	right:=MergeSort(arr[mid:])
	left:=MergeSort(arr[:mid])
	return merge(left,right) 
}

func merge(left,right []int)[]int{
	res:=[]int{}
	i,j:=0,0
	for i<len(left)&&j<len(right){
		if left[i]<=right[j]{
			res=append(res,left[i])
			i++
		}else{
			res=append(res,right[j])
			j++
		}
	}
	if i<len(left){
		res=append(res,left[i:]...)
	}
	if j<len(right){
		res=append(res,right[j:]...)
	}
	return res
}