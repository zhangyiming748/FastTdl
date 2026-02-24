package archive

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/archive"
	"github.com/zhangyiming748/finder"
)

// calculateDirSize calculates the total size of a directory in bytes
func calculateDirSize(dir string) (int64, error) {
	var size int64
	err := filepath.Walk(dir, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// formatBytes converts bytes to MB with two decimal places
func formatBytes(bytes int64) float64 {
	return float64(bytes) / (1024 * 1024)
}

func Videos(dir string) {
	//在这里实现计算dir文件夹的大小，并打印出来
	initialSize, err := calculateDirSize(dir)
	if err != nil {
		log.Printf("无法计算初始文件夹大小: %v\n", err)
	} else {
		log.Printf("处理前文件夹大小: %.2f MB\n", formatBytes(initialSize))
	}
	folders := finder.FindAllFolders(dir)
	for i, folder := range folders {
		files := finder.FindAllVideosInRoot(folder)
		for j, file := range files {
			log.Printf("正在处理第%d/%d个文件夹下的第%d/%d个文件: %s\n", i+1, len(folders), j+1, len(files), file)
			archive.Convert2H265(file)
		}
	}
	//在这里再次实现计算dir文件夹的大小，并打印出来
	finalSize, err := calculateDirSize(dir)
	if err != nil {
		log.Printf("无法计算最终文件夹大小: %v\n", err)
	} else {
		log.Printf("处理后文件夹大小: %.2f MB\n", formatBytes(finalSize))
		//在这里实现计算两次大小的差值并打印以M为单位，保留两位小数的输出
		diff := formatBytes(initialSize - finalSize)
		log.Printf("视频处理后文件夹大小变化: %.2f MB\n", diff)
	}
}

func Images(dir string) {
	//在这里实现计算dir文件夹的大小，并打印出来
	initialSize, err := calculateDirSize(dir)
	if err != nil {
		log.Printf("无法计算初始文件夹大小: %v\n", err)
	} else {
		log.Printf("处理前文件夹大小: %.2f MB\n", formatBytes(initialSize))
	}
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
	//在这里再次实现计算dir文件夹的大小，并打印出来
	finalSize, err := calculateDirSize(dir)
	if err != nil {
		log.Printf("无法计算最终文件夹大小: %v\n", err)
	} else {
		log.Printf("处理后文件夹大小: %.2f MB\n", formatBytes(finalSize))
		//在这里实现计算两次大小的差值并打印以M为单位，保留两位小数的输出
		diff := formatBytes(initialSize - finalSize)
		log.Printf("图片处理后文件夹大小变化: %.2f MB\n", diff)
	}
}
func Movies(dir string) {
	//在这里实现计算dir文件夹的大小，并打印出来
	initialSize, err := calculateDirSize(dir)
	if err != nil {
		log.Printf("无法计算初始文件夹大小: %v\n", err)
	} else {
		log.Printf("处理前文件夹大小: %.2f MB\n", formatBytes(initialSize))
	}
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
	//在这里再次实现计算dir文件夹的大小，并打印出来
	finalSize, err := calculateDirSize(dir)
	if err != nil {
		log.Printf("无法计算最终文件夹大小: %v\n", err)
	} else {
		log.Printf("处理后文件夹大小: %.2f MB\n", formatBytes(finalSize))
		//在这里实现计算两次大小的差值并打印以M为单位，保留两位小数的输出
		diff := formatBytes(initialSize - finalSize)
		log.Printf("电影处理后文件夹大小变化: %.2f MB\n", diff)
	}
}
