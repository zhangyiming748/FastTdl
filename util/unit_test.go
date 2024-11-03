package util

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestOffset(t *testing.T) {
	prefix := "https://t.me/Siwa2024/14935"
	GenerateUrl(1, 200, prefix)
}
func TestGetPercentageSign(t *testing.T) {
	s := "🔮 奇闻异录 与 沙雕时刻 meme collection~ ...  21.3% [........] [0 B in 297ms; 0 B/s]"
	ret := GetPercentageSign(s)
	t.Log(ret)
}
func TestRegexp(t *testing.T) {

	// 要匹配的字符串
	str := "(1249419900):6597 -> /h~ ... done! [184.88 MB in 43.347s; 4.26 MB/s]"

	// 正则表达式，匹配冒号后和箭头前的任意长度数字
	re := regexp.MustCompile(`:(\d+)\s+->`)

	// 查找匹配
	matches := re.FindStringSubmatch(str)
	if len(matches) > 1 {
		fmt.Println("匹配到的数字:", matches[1])
	} else {
		fmt.Println("没有匹配到数字")
	}

}
func TestRename(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取用户的个人文件夹目录:", err)

	}
	home = filepath.Join(home, "Downloads", "telegram")
	key := "6600"
	absFile, err := FindUniqueFile(home, key)
	if err != nil {
		fmt.Println("无法获取用户的个人文件夹目录:", err)
	}
	t.Log(absFile)
	dir := filepath.Dir(absFile)       // 获取目录路径
	fileName := filepath.Base(absFile) // 获取文件名
	fmt.Println("目录路径:", dir)
	fmt.Println("文件名:", fileName)
	suffix := filepath.Ext(fileName)               //扩展名部分 带有.
	prefix := strings.TrimSuffix(fileName, suffix) //文件名部分
	fmt.Println(prefix, suffix)
	newAbsFile := strings.Join([]string{dir, string(os.PathSeparator), "男友视角", suffix}, "")
	fmt.Printf("最终的旧文件名:%s\n新文件名:%v\n", absFile, newAbsFile)
	os.Rename(absFile, newAbsFile)
}
