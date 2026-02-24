package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func ReadByLine(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		log.Println("按行读文件出错")
		return []string{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if len(a) == 0 {
			continue
		}
		line := string(a)
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, "single") {
			uri := strings.Replace(line, "?single", "", 1)
			lines = append(lines, uri)
		} else {
			lines = append(lines, line)
		}
	}
	return lines
}

// 按行写文件
func WriteByLine(fp string, s []string) error {
	file, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("无法打开文件: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, v := range s {
		if _, err := writer.WriteString(v + "\n"); err != nil {
			return fmt.Errorf("写入文件失败: %v", err)
		}
	}
	return writer.Flush()
}

// IsExistPath 判断文件夹是否存在
func IsExistPath(folderPath string) bool {
	info, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		log.Printf("文件夹:%v不存在\n", folderPath)
		return false
	}
	if !info.IsDir() {
		log.Printf("路径:%v不是文件夹\n", folderPath)
		return false
	}
	log.Printf("文件夹:%v存在\n", folderPath)
	return true
}

// IsExistFile 判断文件是否存在
func IsExistFile(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Printf("文件:%v不存在\n", filePath)
		return false
	}
	if info.IsDir() {
		log.Printf("路径:%v是文件夹而不是文件\n", filePath)
		return false
	}
	log.Printf("文件:%v存在\n", filePath)
	return true
}
func IsExistCmd(cmds ...string) bool {
	for _, cmd := range cmds {
		//cmd := "ls" // 需要测试的命令
		_, err := exec.LookPath(cmd)
		if err != nil {
			log.Printf("命令:%s不存在\n", cmd)
			return false
		} else {
			log.Printf("命令:%s存在\n", cmd)
		}
	}
	return true
}

func GetAllFileInfoFast(dir, pattern string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), pattern) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}
func ReadInSlice(fp string) []string {
	fileBytes, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []string{}
	}

	// 创建一个bufio.Reader对象
	reader := bufio.NewReader(bytes.NewReader(fileBytes))

	// 按行读取文件内容并存储到字符串切片中
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, line)
	}

	// 打印结果
	//for i, line := range lines {
	//	fmt.Printf("第%d行: %s\n", i+1, line)
	//}
	return lines
}

/*
统计指定目录下的文件数 相当于find . -type f | wc -l
*/
func CountFiles(dirPath string) (int, error) {
	var count int
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			count++
		}
		return nil
	})
	return count, err
}

// CopyDir 递归复制目录
func CopyDir(src string, dst string) error {
	var err error
	var fds []os.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := filepath.Join(src, fd.Name())
		dstfp := filepath.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

// CopyFile 复制单个文件
func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

/*
获取字符串数组理论上需要下载的条目数
*/
func GetExpectedFilesToAdd(urls []string) int {
	var count int
	for _, uri := range urls {
		if strings.Contains(uri, "%") {
			num, err := strconv.Atoi(strings.Split(uri, "%")[1])
			if err != nil {
				count++
			} else {
				count += num
			}
		} else {
			count++
		}

	}
	return count
}
