package main

import "fmt"

// for循环

func main() {

	/* 	// 基本格式
	   	for i :=0;i < 10; i++{
	   		fmt.Println(i)
	   	} */

	/* 	// 变种1
	   	var i = 5
	   	for ; i < 10; i++ {
	   		fmt.Println(i)
	   	} */

	/* 	// 变种2
	   	var i = 5
	   	for i < 10 {
	   		fmt.Println(i)
	   		i++
		} */

	/* 	// 死循环 等同while循环
	   	for {
	   		fmt.Println("hello")
		} */

	/* 	// for range循环
	   	s := "hello沙河"
	   	for i, v := range s {
	   		fmt.Printf("%d %c\n", i, v)
	   	} */

	/* 	age := 18
	   	for age > 0 {
	   		fmt.Println(age)
	   		age--
	   	} */
	// switchDemo1()

	// 九九乘法表

	for x := 1; x < 10; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d*%d=%d\t", y, x, x*y)
		}
		fmt.Println()
	}

	for x := 9; x >= 1; x-- {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d*%d=%d\t", y, x, x*y)
		}
		fmt.Println()
	}
}
