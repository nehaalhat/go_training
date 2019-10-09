/*
 * Implement interface with N number of methods and implement
 * only N-1 method and call all methods using type associate with them
 */

package main

import "fmt"

type intf interface {
	method1()
	method2()
	method3()
}

func (p person)method1(){
	fmt.Println("Inside method 1")
}

func (p person)method2(){
	fmt.Println("Inside method 2")
}

//func (p person)method3(){
//
//	fmt.Println("Inside method 3")
//
//}

type person struct {
	name string
	age  int
}

func main () {
	fmt.Println("Inside main")
	per := person{}
	per.method1()
	per.method2()
	//per.method3()
	//var test intf
	//test = per
	//test.contactInfo()
}
