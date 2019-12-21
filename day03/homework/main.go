package main

import "fmt"

func main() {
	/* 	// 1.判断字符串中汉字的数量
	   	// 如何判断一个字符是汉字?
	   	s1 := "Hello boy , 沙河在什么地方呀?사랑해"
	   	// 1.依次拿到字符串中的字符
	   	var conut int
	   	for _, c := range s1 {
	   		// 2.判断当前这个字符是不是汉字
	   		if unicode.Is(unicode.Han, c) { // Hangul是韩文
	   			conut++
	   		}
	   	}

	   	// 3.把汉字出现的次数累加得到最终结果
	   	fmt.Println(conut) */

	/* 	// how do you do 单词出现的次数
	   	s2 := "how do you do"
	   	// 1.把字符串按照空格切割得到切片
	   	s3 := strings.Split(s2, " ")
	   	// 2.遍历切片存储到一个map
	   	m1 := make(map[string]int, 10)
	   	for _, w := range s3 {
	   		// 1.如果原来map中不存在w这个key那么出现次数=1
	   		// 2.如果map中存在w这个key,那么出现次数+1
	   		if _, ok := m1[w]; !ok {
	   			m1[w] = 1
	   		} else {
	   			m1[w]++
	   		}
	   	}
	   	// 3.累加出现的次数

	   	for key, value := range m1 {
	   		fmt.Println(key, value)
	   	} */

	// 回文判断
	// 字符串从左往右读和从右往左读是一样的,那么就是回文.
	// 上海自来水来自海上
	// 山西运煤车煤运西山
	// 黄山落叶松叶落山黄
	// 解题思路:
	// 把字符串中的字符拿出来放到一个[]rune中

	ss := "黄山落叶松叶落山黄"
	/* 	r := make([]rune, 0, len(ss))
	   	for _, c := range ss {
	   		r = append(r, c)
	   	} */
	r := []rune(ss)
	fmt.Println("[]rune:", r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是的")
}
