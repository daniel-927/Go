package main

import (
	"Go/day06/mylogger"
	"time"
)

// 测试我们自己的日志库
func main() {
	log := mylogger.Newlog()
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条info日志")
		time.Sleep(time.Second)
	}
}
