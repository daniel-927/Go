package main

import (
	"fmt"
	"os"
)

// 文件操作

func f1() {
	var (
		fileObj *os.File
		err     error
	)
	fileObj, err = os.Open("./main.go")
	defer fileObj.Close() // 如果返回err有值,则fileObj是空指针
	if err != nil {
		fmt.Println("open file failed, err:%v", err)
		return // 退出时执行defer,然后空指针的引用出现错误
	}
}

func f2() {
	var (
		fileObj *os.File
		err     error
	)
	fileObj, err = os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:%v", err)
		return // 如果返回err有值,则直接退出
	}
	defer fileObj.Close() // 退出了则不会进行空指针的引用
}

func main() {

}
