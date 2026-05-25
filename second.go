package main

func Second(arr []int) int {
	large := 0
	second := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > large {
			second = large
			large = arr[i]
		} else if arr[i] > second && arr[i] != large {
			second = arr[i]
		}
	}
	return second
}
