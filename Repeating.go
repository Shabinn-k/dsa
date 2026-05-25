package main

import "fmt"

func Repeat(str string) {
	m := make(map[rune]bool) 
	n:=make(map[rune]int)
	for _, v := range str {
		if m[v] {
			fmt.Println("First repeating character :",string(v))
			break
		}
		m[v]=true
	}
	for _,v:=range str{
		n[v]++
		if n[v]==1{
			fmt.Println("First non-repeating character :",string(v))
			break
		}
	}	 
}