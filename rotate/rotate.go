package rotate

import (
	"github.com/zhangyiming748/archive"
	"github.com/zhangyiming748/finder"
	"log"
)

func RotateVideos(root,direction string) {
	switch direction {
	case "90":
		direction = archive.ToRight
	case  "270":
		direction = archive.ToLeft
	}
	files := finder.FindAllVideos(root)
	for i, file := range files {
		log.Printf("正在处理第%d/%d个文件: %s\n", i+1, len(files), file)
		archive.RotateVideo(file,direction)
	}
}
