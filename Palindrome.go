package main

func Palindrome(str string)bool{
	f:=0
	l:=len(str)-1
	p:=true

	for f<l{
		if str[l]!=str[f]{
			p=false
			break
		}
		f++
		l--
	}
	return p
}