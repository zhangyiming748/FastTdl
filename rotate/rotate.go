package rotate

import (
	"github.com/zhangyiming748/GracefullyExit"
	"github.com/zhangyiming748/archive"
	"github.com/zhangyiming748/finder"
	"log"
)

func RotateVideos(root, direction string) {
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
		if GracefullyExit.ShouldExit() {
			log.Println("Exit signal received. Quitting after current operation.")
			break
		}
	}
}
