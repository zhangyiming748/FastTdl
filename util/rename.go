package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/zhangyiming748/FastTdl/constant"
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
func RenameByKey(of constant.OneFile) {
	p := constant.GetParams()
	key := strconv.Itoa(of.FileId)
	home := p.GetMainFolder()
	if of.Tag != "" {
		home = filepath.Join(home, of.Tag)
		if of.Subtag != "" {
			home = filepath.Join(home, of.Subtag)
		}
	}
	//key := "6600"
	if path := os.Getenv("TDL"); path != "" {
		home = path
	}
	absFile, err := FindUniqueFile(home, key)
	if err != nil {
		log.Printf("无法获取用户的指定文件:%v\t放弃重命名\n", err)
		return
	}
	fmt.Printf("absfile: %s", absFile)
	dir := filepath.Dir(absFile)       // 获取目录路径
	fileName := filepath.Base(absFile) // 获取文件名
	fmt.Println("目录路径:", dir)
	fmt.Println("文件名:", fileName)
	suffix := filepath.Ext(fileName)               //扩展名部分 带有.
	prefix := strings.TrimSuffix(fileName, suffix) //文件名部分
	fmt.Println(prefix, suffix)
	newAbsFile := strings.Join([]string{dir, string(os.PathSeparator), of.FileName, suffix}, "")
	log.Printf("最终的旧文件名:%s\t新文件名:%v\n", absFile, newAbsFile)
	if noRename := os.Rename(absFile, newAbsFile); noRename != nil {
		log.Printf("%s重命名%s失败\n", absFile, newAbsFile)
	} else {
		log.Printf("%s重命名%s成功\n", absFile, newAbsFile)
	}
}
