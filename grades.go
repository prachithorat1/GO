package main

import "fmt"

func main() {
	fmt.Println("Enter the Marks Sarthak.. :")
	var marks int
	fmt.Scanln(&marks)
	if marks < 0 || marks > 100 {
		fmt.Println("please enter valid marks Sarthak..")
	} else if marks >= 60 && marks < 70 {
		fmt.Println("C Grade")
	} else if marks >= 70 && marks < 80 {
		fmt.Println("B Grade")
	} else if marks >= 80 && marks < 90 {
		fmt.Println("A Grade")
	}
}
