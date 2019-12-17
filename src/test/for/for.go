package main

// for循环

import "fmt"

func switchDemo1() {
	finger := 5

	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效")
	}
}

func main() {
	age := 18
	for age > 0 {
		fmt.Println(age)
		age--
	}
	switchDemo1()
}
