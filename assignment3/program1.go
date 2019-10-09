package main

import "fmt"

type intf interface {
	display()
}

type person struct {
	name string
	age  int
}

//func (p *person) display() {
//	p.age += 2
//}

func (p person) display() {
	p.age += 2
}

func main() {
	per := person{ "Neha", 25 }
	per.display()

	fmt.Println("Name is:", per.name, "Age is:", per.age)
}
