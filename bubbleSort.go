package main

func Bubble(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func BubbleDesc(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
	


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
	return Mirror(left.left,right.right)&&Mirror(left.right,right.left)
}