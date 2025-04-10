package archive

import (
	"fmt"
	"github.com/h2non/filetype"
	"github.com/zhangyiming748/FastMediaInfo"
	"github.com/zhangyiming748/FastTdl/constant"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func ArchiveVideo() {
	p := constant.GetParams()
	files, _ := GetAllVideoFiles(p.GetMainFolder())
	for _, file := range files {
		ConvertH265(file)
	}
}

// GetAllVideoFiles 返回指定目录下所有文件的绝对路径
func GetAllVideoFiles(root string) ([]string, error) {
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
			if isVideo(absPath) {
				if !isH265(absPath) {
					files = append(files, absPath)
				}
				if filepath.Ext(absPath) == ".mkv" {
					files = append(files, absPath)
				}
				if isHev1(absPath) {
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

func isVideo(fp string) bool {
	file, _ := os.Open(fp)
	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)
	return filetype.IsVideo(head)
}

func isH265(fp string) bool {
	mi := FastMediaInfo.GetStandMediaInfo(fp)
	if mi.Video.Format == "HEVC" && mi.Video.CodecID == "hvc1" {
		log.Printf("视频:%s格式为 HEVC,跳过转换\n", fp)
		return true
	} else {
		log.Printf("视频:%s格式为 %s,开始转换\n", fp, mi.Video.Format)
		return false
	}
}

func isHev1(fp string) bool {
	mi := FastMediaInfo.GetStandMediaInfo(fp)
	if mi.Video.CodecID == "hev1" {
		return true
	} else {
		return false
	}
}

func outOfFHD(fp string) bool {
	mi := FastMediaInfo.GetStandMediaInfo(fp)
	height, _ := strconv.Atoi(mi.Video.Height)
	width, _ := strconv.Atoi(mi.Video.Width)
	if height > 1920 || width > 1920 {
		log.Printf("视频:%s大于1080p,转换\n", fp)
		return true
	} else {
		return false
	}
}

func ConvertH265(src string) {
	if !isVideo(src) {
		return
	}
	purgePath := filepath.Dir(src)
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	b := seed.Intn(2000)
	tmp := strconv.Itoa(b)
	tmp = strings.Join([]string{tmp, ".mp4"}, "")
	dst := filepath.Join(purgePath, tmp)
	args := []string{"-i", src}
	if runtime.GOARCH == "arm64" && runtime.GOOS == "linux" {
		args = []string{"-threads", "1", "-i", src}
	}
	args = append(args, "-c:v", "libx265")
	args = append(args, "-tag:v", "hvc1")
	if outOfFHD(src) {
		args = append(args, "-vf", "scale=if(gt(iw\\,ih)\\,1920\\,-2):if(gt(iw\\,ih)\\,-2\\,1080)")
	}
	// args = append(args, "-c:a", "libmp3lame")
	args = append(args, "-c:a", "aac")
	args = append(args, dst)
	cmd := exec.Command("ffmpeg", args...)
	if isHev1(src) {
		cmd = exec.Command("ffmpeg", "-i", src, "-c:v", "copy", "-tag:v", "hvc1", "-c:a", "aac", dst)
	} else {
		if isH265(src) {
			return
		}
	}
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
				fmt.Print(string(buf[:n]))
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
		log.Fatalf("转换失败：%v\n", err)
	} else {
		// 先尝试删除源文件
		if err := os.Remove(src); err != nil {
			log.Fatalf("删除源文件失败：%v\n", err)
		}
		// 源文件删除成功后，等待短暂时间确保文件句柄完全释放
		time.Sleep(100 * time.Millisecond)
		// 尝试重命名
		if err := os.Rename(dst, src); err != nil {
			log.Fatalf("重命名文件失败：%v\n", err)
		}
	}
}
