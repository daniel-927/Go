package main

import "fmt"

// Go语言中函数的return不是原子操作,在底层是分为两步来执行
// 第一步:返回值赋值
// defer
// 第二步:真正的RET返回
// 函数中如果存在defer,那么defer执行的时机是在第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x不是返回值的赋值
	}()
	return x // 1.返回值赋值 2.defer 3.真正的RET指令
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值赋值为5后,defer执行+1, 结果为6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值是将x赋值给y,所以等于5后,x+1等于6.  实际返回值5
}
func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中x的副本
	}(x)
	return 5 // 返回值 = x = 5
}
func f5() (x int) {
	defer func(x int) int {
		x++ // 改变的是函数中x的副本
		return x
	}(x)
	return 5 // 返回值 = x = 5
}

// 传一个x的指针到匿名函数中
func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 1.返回值=x=5 2.defer x=6 3.RET返回
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
}
