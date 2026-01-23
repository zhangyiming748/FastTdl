package util

import (
	"log"
	"time"
)

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Asia/Shanghai")
}

// CheckHour 检查当前时间是否到达指定小时
// hour: 指定的小时数（24小时制）
// 当到达指定小时时返回 true
func CheckHour(hour string) {
	for {
		currentHour := time.Now().In(loc).Format("15") // 移到循环内部，每次都获取最新时间
		if currentHour == hour {
			return
		}
		log.Printf("当前时间为 %s,未达到 %s 点,等待30分钟后再次检查...\n", currentHour, hour)
		time.Sleep(30 * time.Minute) // 缩短检查间隔为1分钟
	}
}

// CheckExactTime 检查是否到达指定的具体时间点
// timeStr: 指定的时间字符串，格式为 "HH:MM:SS"，例如 "08:30:00"
func CheckExactTime(timeStr string) {
	for {
		// 使用上海时区获取当前时间
		currentTime := time.Now().In(loc).Format("15:04:05")
		if currentTime == timeStr {
			return
		}
		log.Printf("当前时间为 %s,未达到 %s 点,等待30分钟后再次检查...\n", currentTime, timeStr)
		time.Sleep(1 * time.Second)
	}
}
