package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello golang")

	a, b, c := 1, true, "hello"
	// a := 1
	// b := true
	// c := "hello"
	// var (
	// 	a int    = 1
	// 	b bool   = true
	// 	c string = "hello"
	// )
	fmt.Printf("%v, %v %v", a, b, c)
}
