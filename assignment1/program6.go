package main

import (
	"fmt"
)

func main() {
	var number int
	sum := 0
	fmt.Println("Enter the number to be reversed:")
        fmt.Scanf("%d", &number)

	for number > 0 {
		remainder := number % 10
		sum *= 10
		sum += remainder
		number /= 10
	}
	fmt.Println("Reversed Number is :", sum)

}
