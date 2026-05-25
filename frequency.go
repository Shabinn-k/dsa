package main

import "fmt"

func Frequency(arr []string) {
	m := make(map[string]int)
	for _, ch := range arr {
		m[ch]++
	}
	for i, n := range m {
		fmt.Printf("%s :%d\n", i, n)
	}
}

func Characters(str string) {
	m := make(map[rune]int)
	for _, v := range str {
		m[v]++
	}
	for i, n := range m {
		fmt.Printf("%c :%d\n", i, n)
	}
}

func Duplicate(arr []int) {
	m := make(map[int]int)
	dupe := []int{}
	og := []int{}

	for _, ch := range arr {
		m[ch]++
		if m[ch] < 2 {
			og = append(og, ch)
		} else {
			dupe = append(dupe, ch)
		}
	}
	fmt.Println("original array ", arr)
	fmt.Println("after removing ", og)
	fmt.Println("duplicates ", dupe)
}
