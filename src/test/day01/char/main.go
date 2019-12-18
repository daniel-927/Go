package main

import "fmt"

// 字符

func main() {
	s1 := "Golang"
	c1 := 'G'
	fmt.Println(s1, c1)
	s2 := "中国"
	c2 := '中'
	fmt.Println(s2, c2)

	s3 := "hello沙河"
	fmt.Println(len(s3))

	// 遍历字符串
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%c\n", s3[i])
	}
	fmt.Println()
	// for range 循环 是按照rune类型去遍历
	for k, v := range s3 {
		fmt.Printf("%d%c\n", k, v)
	}

	// 强制类型转换
	s5 := "big"
	byterArray := [] byte(s5)
	fmt.Println(byterArray)
	byterArray[0] = 'p'
	fmt.Println(byterArray)
	//将字节数组强制转成字符串类型
	s5 = string(byterArray)
	fmt.Println(s5)
}
