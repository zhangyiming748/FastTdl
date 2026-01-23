package archive

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/archive"
	"github.com/zhangyiming748/finder"
)

func Videos(dir string) {
	folders := finder.FindAllFolders(dir)
	for i, folder := range folders {
		files := finder.FindAllVideosInRoot(folder)
		for j, file := range files {
			log.Printf("正在处理第%d/%d个文件夹下的第%d/%d个文件: %s\n", i+1, len(folders), j+1, len(files), file)
			archive.Convert2H265(file)
		}
	}
}

func Images(dir string) {
	folders := finder.FindAllFolders(dir)
	for i, folder := range folders {
		files := finder.FindAllImagesInRoot(folder)
		for j, file := range files {
			if filepath.Ext(file) == ".avif" {
				continue
			}
			log.Printf("正在处理第%d个文件夹下的第%d个文件: %s\n", i+1, j+1, file)
			archive.Convert2AVIF(file)
		}
	}
}
func Movies(dir string) {
	folders := finder.FindAllFolders(dir)
	for i, folder := range folders {
		files := finder.FindAllVideosInRoot(folder)
		for j, file := range files {
			if strings.ToUpper(filepath.Ext(file)) == ".MKV" {
				log.Printf("正在处理第%d/%d个文件夹下的第%d/%d个mkv文件: %s\n", i+1, len(folders), j+1, len(files), file)
				archive.ConvertMKV2H265(file)
			}
		}
	}
}
