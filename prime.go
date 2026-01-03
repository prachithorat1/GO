package main 
import "fmt"

func main(){
	var number int 
	fmt.Print("Enter any number yto check it is Prime or not")
	fmt.Scanf("%d", &number)
}

if isPrime(number) {
	fmt.Printf("The number %d is Prime number.\n", number)
}else {
	fmt.Printf("The number %d is Not a Prime number.\n", number)
}
