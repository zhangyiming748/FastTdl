package util

import (
	"github.com/zhangyiming748/lumberjack"
	"io"
	"log"
	"os"
)

func SetLog(l string) {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   l,
		MaxSize:    1, // MB
		MaxBackups: 1,
		MaxAge:     28, // days
	}
	err := fileLogger.Rotate()
	if err != nil {
		log.Println("转换新日志文件失败", err)
	}
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
