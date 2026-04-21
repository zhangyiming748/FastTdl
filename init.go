package main

import (
	"FastTdl/tdl"
	"FastTdl/util"
	"log"
)

// Initialize 统一初始化函数，在 main 函数开始时调用
func Initialize() {
	log.Println("开始初始化 FastTdl...")

	// 1. 初始化时区
	util.InitTimeZone()
	log.Println("时区初始化完成")

	// 2. 初始化日志系统
	util.SetLog("tdl.log")
	log.Println("日志系统初始化完成")

	// 3. 加载中英文映射表
	tdl.LoadTranslationMap()
	log.Println("中英文映射表加载完成")

	log.Println("FastTdl 初始化完成")
}
