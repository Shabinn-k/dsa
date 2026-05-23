package main
import "fmt"

func main() {
	//reverese
	str1:="hello world"
	fmt.Println(RS(str1))

	//bubble sort
	arr1:=[]int{5,1,4,2,8}
	BS:=Bubble(arr1)
	fmt.Println(BS)

	//Palindrome
	str2:="malayalam"
	PL:=Palindrome(str2)
	fmt.Println(PL)

	//count v,c
	Vowel("hello how are yopu my friend") 

	//anagram
	str3:="anagram"
	str4:="nagaram"
	VA:=Anagram(str3,str4)
	fmt.Println(VA)

	//second largest from array
	arr2:=[]int{1,2,3,5,6,8,4,10}
	SL:=Second(arr2)
	fmt.Println(SL)
}