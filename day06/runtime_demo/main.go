package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller()

func f() {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("runtime.Caller() failed, err:%v\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file) // runtime_demo/main.go
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func f1() {
	f()
}

func main() {
	f1()
}
