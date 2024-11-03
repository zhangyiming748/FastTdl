package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindUniqueFile(dir string, searchStr string) (string, error) {
	var foundFile string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查是否是文件且文件名包含指定字符串
		if !info.IsDir() && strings.Contains(info.Name(), searchStr) {
			if foundFile != "" {
				// 如果已经找到一个文件，再次找到则返回错误
				return fmt.Errorf("找到多个文件: %s 和 %s", foundFile, path)
			}
			foundFile = path // 记录找到的文件路径
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if foundFile == "" {
		return "", fmt.Errorf("未找到包含 '%s' 的文件", searchStr)
	}
	// 返回找到的文件的绝对路径
	absPath, err := filepath.Abs(foundFile)
	if err != nil {
		return "", err
	}
	return absPath, nil
}
func RenameByKey(key, words string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取用户的个人文件夹目录:", err)
	}
	home = filepath.Join(home, "Downloads", "telegram")
	//key := "6600"
	absFile, err := FindUniqueFile(home, key)
	if err != nil {
		fmt.Println("无法获取用户的指定文件:", err)
	}
	fmt.Printf("absfile: %s", absFile)
	dir := filepath.Dir(absFile)       // 获取目录路径
	fileName := filepath.Base(absFile) // 获取文件名
	fmt.Println("目录路径:", dir)
	fmt.Println("文件名:", fileName)
	suffix := filepath.Ext(fileName)               //扩展名部分 带有.
	prefix := strings.TrimSuffix(fileName, suffix) //文件名部分
	fmt.Println(prefix, suffix)
	newAbsFile := strings.Join([]string{dir, string(os.PathSeparator), words, suffix}, "")
	fmt.Printf("最终的旧文件名:%s\t新文件名:%v\n", absFile, newAbsFile)
	if noRename := os.Rename(absFile, newAbsFile); noRename != nil {
		log.Printf("%s重命名%s失败\n", absFile, newAbsFile)
	} else {
		log.Printf("%s重命名%s成功\n", absFile, newAbsFile)

	}
}
