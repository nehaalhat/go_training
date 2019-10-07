package main

import (
	"fmt"
)

func main(){
	for i := 1; i <= 5; i++{
		for j := i; j < 5; j++{
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++{
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}

/*OUTPUT
    *
   **
  ***
 ****
*****
 */