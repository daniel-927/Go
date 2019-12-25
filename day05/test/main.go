package main

import "fmt"

type student struct {
	name string
	age  int
}
// test
func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "大王八", age: 9000},
		{name: "娜扎", age: 23},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
}
