package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/h2non/filetype"
)

// ... existing code ...

// FindLatestFileGo 使用 Go 标准库在指定目录中查找最近创建或修改的文件
// dir: 要搜索的目录路径
// returns: 最新文件的绝对路径，如果出错则返回空字符串
func FindLatestFileGo(dir string) string {
	var latestFile string
	var latestTime int64

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("访问路径出错 %s: %v\n", path, err)
			return nil // 继续遍历
		}
		// 跳过目录和隐藏文件
		if info.IsDir() || strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}
		// 跳过非视频文件
		file, _ := os.Open(path)
		defer file.Close()
		// We only have to pass the file header = first 261 bytes
		head := make([]byte, 261)
		file.Read(head)
		if !filetype.IsVideo(head) {
			return nil
		} 
		
		// 更新最新文件
		if modTime := info.ModTime().Unix(); modTime > latestTime {
			latestTime = modTime
			latestFile = path
			log.Printf("找到更新的文件: %s, 修改时间: %v\n", path, info.ModTime())
		}
		return nil
	})

	if err != nil {
		log.Printf("遍历目录出错: %v\n", err)
		return ""
	}
	return latestFile
}
