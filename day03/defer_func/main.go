package main

import "fmt"

// defer

// defer多用于函数结束之前释放资源(文件句柄,数据库连接,socket连接)
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("heiheihei") // defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("wuwuwu")    // 逆序执行
	fmt.Println("end")
}

func main() {
	deferDemo()
}
