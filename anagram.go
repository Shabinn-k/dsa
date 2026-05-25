package main

func Anagram(str1, str2 string) bool {
	m := make(map[rune]int)
	if len(str1) != len(str2) {
		return false
	}
	for _, v := range str1 {
		m[v]++
	}
	for _, v := range str2 {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
