package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 要求：
// f1(f2)

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func main() {
	// ret := f3(f2, 100, 200)
	f1(f3(f2, 100, 200)) // 把原来需要传递两个int类型的参数包装成了一个不需要传参的函数
}

/*
func main() {
	ret := lixiang(f2, 100, 200)
	ret()
}

func lixiang(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}
*/
