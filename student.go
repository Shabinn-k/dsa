package main

import "fmt"

func Student() {
	std := make(map[string]int)

	// add students
	std["shabin"] = 100
	std["john"] = 80
	std["rahul"] = 60
	std["modi"] = 10
	fmt.Println(std)

	// print marks
	fmt.Println(std["shabin"])
	fmt.Println(std["modi"])

	// update marks
	std["john"]=85
	fmt.Println(std["john"])

	// delete a student
	delete(std,"modi")
	fmt.Println(std)
}