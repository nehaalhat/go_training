package main

import (
	"fmt"
)

func main() {
	k := 1
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(k, " ")
			k += 1
		}
		fmt.Println("")
	}
}

/* OUTPUT
1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
 */
