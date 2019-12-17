package main

import "fmt"

const pi = 3.14

const (
	a,b = iota + 1, iota +2
	c,d
	e,f
)

func main() {
	name := "daniel"
	age := 18
	fmt.Println(name)
	fmt.Printf("%s is good boy %d\n", name, age)
	fmt.Println(pi)
	fmt.Println(a,b,c,d,e,f)
}
