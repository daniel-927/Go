package main

import "fmt"

// 类型断言1
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串:", str)
	}
}

// 类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串:", t)
	case int:
		fmt.Println("是一个int:", t)
	case int64:
		fmt.Println("是一个int64:", t)
	case bool:
		fmt.Println("是一个bool:", t)
	}
}

func main() {
	assign(100)
	assign2(true)
}
