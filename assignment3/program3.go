package main

import (
	"fmt"
	"reflect"
)

func main() {
	test1 := [3]string {"a", "b", "c"}
	fmt.Println("type of i is ", reflect.TypeOf(test1).Kind())
	fmt.Println("value of i is ", reflect.ValueOf(test1))

	test2 := []string {"d", "e"}
	t := reflect.Append(reflect.ValueOf(test2), reflect.ValueOf("f"))
	fmt.Println(t)

	fmt.Println(reflect.Swapper(reflect.ValueOf(test1)))
}
