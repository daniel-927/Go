package calc

import "fmt"

// 包中的标识符（变量名、函数名、结构体、接口等）如果首字母是小写的，表示私有（只能在这个包中使用）
// 首字母大写的标识符可以被外部的包调用

var x = 100

const pi = 3.14

func init() {
	fmt.Println("thi is  calc")
	fmt.Println(x, pi)
}

func Add(x, y int) int {
	return x + y
}
