package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// Logger 日志结构体
type ConsoLogger struct {
	Level LogLevel
}

// NewLog 构造函数
func Newlog(levelStr string) ConsoLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoLogger{
		Level: level,
	}
}

func (c ConsoLogger) enable(logLevel LogLevel) bool {
	return c.Level <= logLevel
}

func (c ConsoLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s][%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (c ConsoLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c ConsoLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

func (c ConsoLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c ConsoLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
