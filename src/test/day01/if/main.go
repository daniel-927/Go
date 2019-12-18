package main

import "fmt"

func main() {

	// 多个条件判断

	/* 	age := 18
	   	if age > 18 {
	   		fmt.Println("澳门首家线上赌场开业啦")
	   	} else if age < 18 {
	   		fmt.Println("warning...")
	   	} else {
	   		fmt.Println("you are ok!")
		   } */

	// 特殊写法，变量仅作用于此if循环内
	if age := 17; age > 18 {
		fmt.Println("澳门首家线上赌场开业啦")
	} else {
		fmt.Println("该写作业啦")
	}
}
