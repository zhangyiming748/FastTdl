package archive

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/h2non/filetype"
	"github.com/zhangyiming748/FastTdl/constant"
)

func ArchiveImage() {
	p := constant.GetParams()
	files, _ := GetAllImageFiles(p.GetMainFolder())
	for _, file := range files {
		ConvertH265(file)
	}
}

// GetAllVideoFiles 返回指定目录下所有文件的绝对路径
func GetAllImageFiles(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 跳过目录，只收集文件
		if !info.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			if isImage(absPath) {
				if filepath.Ext(absPath) != ".avif" {
					files = append(files, absPath)
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}

func isImage(fp string) bool {
	file, _ := os.Open(fp)
	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)
	return filetype.IsImage(head)
}

func ConvertAVIF(src string) {
	if !isImage(src) {
		return
	}
	purgePath := filepath.Dir(src)
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	b := seed.Intn(2000)
	tmp := strconv.Itoa(b)
	tmp = strings.Join([]string{tmp, ".avif"}, "")
	dst := filepath.Join(purgePath, tmp)

	args := []string{"-i", src}
	args = append(args, "-c:v", "libaom-av1")
	args = append(args, "-still-picture", "1")
	args = append(args, dst)
	cmd := exec.Command("ffmpeg", args...)

	// 获取输出和错误管道
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	// 启动命令
	log.Printf("开始执行命令:%s\n", cmd.String())
	if err := cmd.Start(); err != nil {
		log.Fatalln("启动转换失败：", err)
		return
	}

	// 创建一个通道来等待所有输出处理完成
	done := make(chan bool)

	// 在后台处理输出
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if n > 0 {
				fmt.Print(string(buf[:n]))
			}
			if err != nil {
				break
			}
		}
		done <- true
	}()

	// 在后台处理错误输出
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stderr.Read(buf)
			if n > 0 {
				log.Print(string(buf[:n]))
			}
			if err != nil {
				break
			}
		}
		done <- true
	}()

	// 等待输出处理完成
	<-done
	<-done

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		log.Printf("转换失败：%v\n", err)
		return
	} else {
		os.Remove(src)
		src = strings.Replace(src, filepath.Ext(src), ".avif", 1)
		os.Rename(dst, src)
	}
}
