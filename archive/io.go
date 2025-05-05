package archive

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GetFinalSubDirs 获取给定目录下所有最终子目录（不包含子文件夹的目录）
func GetFinalSubDirs(root string) ([]string, error) {
	var finalDirs []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理目录
		if !info.IsDir() {
			return nil
		}

		// 检查当前目录是否包含子目录
		hasSubDir := false
		entries, err := os.ReadDir(path)
		if err != nil {
			log.Printf("读取目录 %s 失败: %v\n", path, err)
			return nil
		}

		for _, entry := range entries {
			if entry.IsDir() {
				hasSubDir = true
				break
			}
		}

		// 如果是最终子目录（不包含其他子目录），则添加到结果中
		if !hasSubDir {
			finalDirs = append(finalDirs, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("遍历目录失败: %v", err)
	}

	return finalDirs, nil
}
