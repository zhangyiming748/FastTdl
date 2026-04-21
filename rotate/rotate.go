package rotate

import (
	"FastTdl/util"
	"log"
	"os"

	"github.com/zhangyiming748/archive"
	"github.com/zhangyiming748/finder"
)

func RotateVideos(root, direction string) {
	util.SetExit()
	var archiveDirection string

	switch direction {
	case "90":
		archiveDirection = archive.ToRight
	case "270":
		archiveDirection = archive.ToLeft
	default:
		log.Printf("警告：无效的旋转方向 '%s'，使用默认方向 90 度\n", direction)
		archiveDirection = archive.ToRight
		direction = "90"
	}

	files := finder.FindAllVideos(root)
	if len(files) == 0 {
		log.Printf("在目录 %s 中未找到视频文件\n", root)
		return
	}

	for i, file := range files {
		log.Printf("正在处理第%d/%d个文件: %s\n", i+1, len(files), file)
		archive.RotateVideo(file, archiveDirection)
		if util.GetExit() {
			log.Printf("接收到退出信号,程序在保证原子操作的情况下正常退出")
			os.Exit(0)
		}
	}
}
