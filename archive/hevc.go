package archive

import (
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/h2non/filetype"
	"github.com/zhangyiming748/FastMediaInfo"
)

var (
	hasFFmpeg    bool
	hasMediainfo bool
)

func init() {
	// 检查 ffmpeg
	if _, err := exec.LookPath("ffmpeg"); err == nil {
		hasFFmpeg = true
	}

	// 检查 mediainfo
	if _, err := exec.LookPath("mediainfo"); err == nil {
		hasMediainfo = true
	}

	if !hasFFmpeg || !hasMediainfo {
		panic("缺少必要的软件依赖：ffmpeg 或 mediainfo 未安装")
	}
}

// GetAllVideoFiles 返回指定目录下所有文件的绝对路径
func GetAllFiles(root string) ([]string, error) {
	if !hasFFmpeg || !hasMediainfo {
		log.Fatalln("缺少必要的软件依赖：ffmpeg 或 mediainfo 未安装,跳过最终存档步骤")
	}
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
			if isVideo(absPath) && !isH265(absPath) {
				files = append(files, absPath)
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
	if mi.Video.Format == "HEVC" {
		return true
	} else {
		return false
	}
}

func ConvertH265(src string) {
	purgePath:=filepath.Dir(src)
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	b := seed.Intn(2000)
	tmp := strconv.Itoa(b)
	tmp=strings.Join([]string{tmp,".mp4"}, "")
	dst := filepath.Join(purgePath, tmp)
	cmd := exec.Command("ffmpeg", "-i", src, "-c:v", "libx265", "-tag:v", "hvc1", "-c:a", "libmp3lame", dst)
	err := cmd.Run()
	if err != nil {
		log.Fatalln("转换失败：", err)
	}else{
		os.Remove(src)
		os.Rename(dst, src)
	}
}
