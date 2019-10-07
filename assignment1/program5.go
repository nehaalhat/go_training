package main

import (
	"fmt"
)

func main() {
	num1, num2 := 1, 1
	num3 := 0

	for i := 0; i<=10; i++{
		fmt.Print(num1, " ")
		num3 = num1 + num2
		num1, num2 = num2, num3
	}
}

/* OUTPUT
1 1 2 3 5 8 13 21 34 55 89
 */
