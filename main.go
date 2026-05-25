package main
import "fmt"

func main() {
	//reverese
	str1:="hello world"
	fmt.Println(RS(str1))

	//bubble sort
	arr1:=[]int{7 ,4 ,1 ,9 ,3}
	BS:=Bubble(arr1)
	fmt.Println(BS)

	arr2:=[]int{2 ,8 ,1 ,5 ,4}
	BSD:=BubbleDesc(arr2)
	fmt.Print(BSD)

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
	arr3:=[]int{1,2,3,5,6,8,4,10}
	SL:=Second(arr3)
	fmt.Println(SL)

	Student()

	arr4:=[]string{"go", "java", "go", "python", "go"}
	Frequency(arr4)

	Characters("golang")

	arr5:=[]int{1,2,2,3,4,3,5,4,5,6,7,8,9,7}
	Duplicate(arr5)

	Repeat("programming")

	arr6:=[]int{2,3,4,5}
	target:=9
	TwoSum(target,arr6)

}