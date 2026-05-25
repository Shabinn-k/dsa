package main

import "fmt"

func Vowel(str string) {
	countV := 0
	countC := 0

	for i := 0; i < len(str); i++ {
		ch := str[i]
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			countV++
		} else {
			countC++
		}
	}
	fmt.Println("Vowels :", countV)
	fmt.Println("Consonants :", countC)
}
