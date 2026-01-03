package main

import "fmt"

func operate(a int, b int) (int, int) {
	return a + b, a * b
}
func main() {
	sum, product := operate(50.0, 7.0)
	fmt.Println("Sum : ", sum)
	fmt.Println("Product : ", product)

}
