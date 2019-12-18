package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	path := "\"D:\\VM\\CentOS 8\""
	fmt.Println(path)

	s := "I'm OK"
	fmt.Println(s)
	// 多行字符串
	s2 := `
	一朝春尽红颜老
	花落人亡两不知
	`
	fmt.Println(s2)
	s3 := `"D:\VM\CentOS 8"`
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(s3))

	// 字符串拼接
	name := "理想"
	world := "dsb"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	// 分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "理想"))
	// 前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))
	// 定位
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "b"))     // 首次出现位置
	fmt.Println(strings.LastIndex(s4, "b")) // 最后出现位置
	// 拼接
	fmt.Println(strings.Join(ret, "+"))
}
