package main


func QuickSort(arr []int,low,high int){
	if low<high{
		pi:=Partition(arr,low,high)
		QuickSort(arr,low,pi-1)
		QuickSort(arr,pi+1,high)
	}
}

func Partition(arr []int,low,high int)int{
	Pivot:=arr[high]
	i:=low-1
	for j:=low;j<high;j++{
		if arr[j]<Pivot{
			i++
			arr[i],arr[j]=arr[j],arr[i]
		}
	}
	arr[i+1],arr[high]=arr[high],arr[i+1]
	return i+1
}