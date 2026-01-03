package main

import "fmt"

func main() {
	var score float64 = 92.93
	if score >= 90 {
		fmt.Println("you have won the Gold medal\n")
	} else {
		fmt.Println("Silver medal is fixed.. for your country")
	}
}
