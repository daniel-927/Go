package main

import "fmt"

// 字符串反转操作

func main() {
	s1 := "hello"
	byteArray := []byte(s1)
	s2 := ""
	for i := len(byteArray) - 1; i >= 0; i-- {
		// i 是 4 3 2 1 0
		// byteArray[i] o l l e h
		s2 += string(byteArray[i])
	}
	fmt.Println(s2)

	// 方法2
	length := len(byteArray)

	for i := 0; i < length/2; i++ {
		byteArray[i], byteArray[length-1-i] = byteArray[length-1-i], byteArray[i]
	}
	fmt.Println(string(byteArray))
}
