package main

import "fmt"

func first() {
    fmt.Println("First")
}
func second() {
    fmt.Println("Second")
}
func main() {
    defer second()
    first()
    third()
}

//Deferred function calls are executed in Last In First Out order after the surrounding function returns

func third() {
    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }
}
