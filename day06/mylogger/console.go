package mylogger

import "fmt"

// 往终端写日志相关内容

// Logger 日志结构体
type Logger struct {
}

// NewLog 构造函数
func Newlog() Logger {
	return Logger{}
}

func (l Logger) Debug(msg string) {
	fmt.Println(msg)
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l Logger) Warning(msg string) {
	fmt.Println(msg)
}

func (l Logger) Error(msg string) {
	fmt.Println(msg)
}

func (l Logger) Fatal(msg string) {
	fmt.Println(msg)
}
