package archive

import (
	"FastTdl/util"
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
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// 记录警告但继续处理其他文件
			log.Printf("警告: 无法访问 %s: %v\n", path, err)
			return nil // 返回 nil 而不是 err，这样 Walk 会继续处理其他文件
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// formatBytes converts bytes to MB with two decimal places
func formatBytes(bytes int64) float64 {
	return float64(bytes) / (1024 * 1024)
}

func Videos(dir string, fhd bool) {
	log.Printf("开始检查视频依赖...\n")
	archive.CheckVideoDependencies()
	log.Printf("视频依赖检查完成\n")
	//在这里实现计算dir文件夹的大小，并打印出来
	log.Printf("开始计算初始文件夹大小: %s\n", dir)
	initialSize, err := calculateDirSize(dir)
	if err != nil {
		log.Printf("无法计算初始文件夹大小: %v\n", err)
	} else {
		log.Printf("处理前文件夹大小: %.2f MB\n", formatBytes(initialSize))
	}
	log.Printf("开始查找所有子文件夹: %s\n", dir)
	folders := finder.FindAllFolders(dir)
	log.Printf("找到 %d 个子文件夹\n", len(folders))
	for i, folder := range folders {
		log.Printf("正在处理第 %d/%d 个文件夹: %s\n", i+1, len(folders), folder)
		files := finder.FindAllVideosInRoot(folder)
		log.Printf("在文件夹 %s 中找到 %d 个视频文件\n", folder, len(files))
		for j, file := range files {
			if util.GetExit() {
				log.Printf("接收到退出信号,程序在保证原子操作的情况下正常退出")
				os.Exit(0)
			}
			log.Printf("正在处理第%d/%d个文件夹下的第%d/%d个文件: %s\n", i+1, len(folders), j+1, len(files), file)
			archive.Convert2H265(file, fhd)
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
	log.Printf("开始检查图片依赖...\n")
	archive.CheckImageDependencies()
	log.Printf("图片依赖检查完成\n")
	//在这里实现计算dir文件夹的大小，并打印出来
	log.Printf("开始计算初始文件夹大小: %s\n", dir)
	initialSize, err := calculateDirSize(dir)
	if err != nil {
		log.Printf("无法计算初始文件夹大小: %v\n", err)
	} else {
		log.Printf("处理前文件夹大小: %.2f MB\n", formatBytes(initialSize))
	}
	log.Printf("开始查找所有子文件夹: %s\n", dir)
	folders := finder.FindAllFolders(dir)
	log.Printf("找到 %d 个子文件夹\n", len(folders))
	for i, folder := range folders {
		files := finder.FindAllImagesInRoot(folder)
		log.Printf("在文件夹 %s 中找到 %d 个图片文件\n", folder, len(files))
		for j, file := range files {
			if filepath.Ext(file) == ".avif" {
				continue
			}
			log.Printf("正在处理第%d个文件夹下的第%d个文件: %s\n", i+1, j+1, file)
			archive.Convert2AVIF(file)
			if util.GetExit() {
				log.Printf("接收到退出信号,程序在保证原子操作的情况下正常退出")
				os.Exit(0)
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
				archive.ConvertMKV2H265(file, false)
				if util.GetExit() {
					log.Printf("接收到退出信号,程序在保证原子操作的情况下正常退出")
					os.Exit(0)
				}
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
func Audios(dir string) {
	//在这里实现计算dir文件夹的大小，并打印出来
	log.Printf("开始计算初始文件夹大小: %s\n", dir)
	initialSize, err := calculateDirSize(dir)
	if err != nil {
		log.Printf("无法计算初始文件夹大小: %v\n", err)
	} else {
		log.Printf("处理前文件夹大小: %.2f MB\n", formatBytes(initialSize))
	}
	log.Printf("开始查找所有子文件夹: %s\n", dir)
	folders := finder.FindAllFolders(dir)
	log.Printf("找到 %d 个子文件夹\n", len(folders))
	for i, folder := range folders {
		files := finder.FindAllAudiosInRoot(folder)
		log.Printf("在文件夹 %s 中找到 %d 个音频文件\n", folder, len(files))
		for j, file := range files {
			log.Printf("正在处理第%d个文件夹下的第%d个文件: %s\n", i+1, j+1, file)
			archive.Convert2Mp3(file)
			if util.GetExit() {
				log.Printf("接收到退出信号,程序在保证原子操作的情况下正常退出")
				os.Exit(0)
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
		log.Printf("音频处理后文件夹大小变化: %.2f MB\n", diff)
	}
}

func Dji(src ,dst string){
	log.Printf("开始处理Dji文件...\n")
	if src==dst{
		log.Fatalf("源目录和目标目录不能相同\n")
	}
	for _, video := range finder.FindAllVideos(src) {
		base := filepath.Base(src)
		target := filepath.Join(dst, base)
		archive.DjiVideoConvert(video, target)
	}
}