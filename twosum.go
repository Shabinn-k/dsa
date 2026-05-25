package main

import "fmt"

func TwoSum(target int, arr []int) {
	m := make(map[int]int)
	for i, v := range arr {
		get := target - v

		if ind, foun := m[get]; foun {
			fmt.Println(i, ind)
			return
		}
		m[v] = i
	}
	fmt.Println("Nothing")
}
