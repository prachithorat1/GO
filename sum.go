package main

import "fmt"

func main() {
	var num int

	fmt.Printf("enter a  number till you want result")
	fmt.Scanf("%d", &num)

	result = summition(n)

	fmt.Printf("Sum is %.2f\n", result)
}

func fact(a int) int {
	if n == 0 {
		return 1
	}
	fact := 1
	for i := 1; i <= num; i++ {
		fact = fact * i
	}
	return fact
}

func summition(n int) int {
	sum := 0
	for i := 1; i <= 1; i++ {
		sum += float64(i) / fact(i)
	}
	return sum
}
