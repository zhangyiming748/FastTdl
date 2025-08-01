package archive

import (
	"github.com/h2non/filetype"
	"github.com/zhangyiming748/FastMediaInfo"
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

func ArchiveVideo(dir string) {
	files, _ := GetAllVideoFiles(dir)
	for _, file := range files {
		ConvertH265(file)
	}
}

// GetAllVideoFiles 返回指定目录下所有文件的绝对路径
func GetAllVideoFiles(root string) ([]string, error) {
	var files []string
	count := 0
	maxFiles := 100 // 每批处理的最大文件数

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if count >= maxFiles {
			return filepath.SkipAll // 找到足够的文件后停止遍历
		}
		if err != nil {
			return err
		}
		if !info.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			if isVideo(absPath) {
				if !isH265(absPath) || filepath.Ext(absPath) == ".mkv" || isHev1(absPath) {
					files = append(files, absPath)
					count++
				}
			}
		}
		return nil
	})

	log.Printf("本次找到 %d 个待处理文件\n", count)
	return files, err
}

func isVideo(fp string) bool {
	file, err := os.Open(fp)
	if err != nil {
		return false
	}
	defer file.Close()
	head := make([]byte, 261)
	if _, err := file.Read(head); err != nil {
		return false
	}
	return filetype.IsVideo(head)
}

func isH265(fp string) bool {
	mi := FastMediaInfo.GetStandMediaInfo(fp)
	if mi.Video.Format == "HEVC" && mi.Video.CodecID == "hvc1" && filepath.Ext(fp) == ".mp4" {
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
		log.Printf("视频:%s高度为%s,宽度为%s\n", fp, mi.Video.Height, mi.Video.Width)
		return true
	} else {
		return false
	}
}

func ConvertH265(src string) {
	if !isVideo(src) {
		return
	}
	// if strings.Contains(src, "vp9") {
	// 	return
	// }
	purgePath := filepath.Dir(src)
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	b := seed.Intn(2000) + 1000
	tmp := strconv.Itoa(b)
	tmp = strings.Join([]string{tmp, ".mp4"}, "")
	dst := filepath.Join(purgePath, tmp)
	args := []string{"-i", src}
	if runtime.GOARCH == "arm64" && runtime.GOOS == "linux" {
		args = []string{"-threads", "1", "-i", src}
	}
	if isH265(src) && filepath.Ext(src) == ".mkv" {
		args = append(args, "-c:v", "copy")
	} else {
		args = append(args, "-c:v", "libx265")
	}

	args = append(args, "-tag:v", "hvc1")
	if outOfFHD(src) {
		args = append(args, "-vf", "scale=if(gt(iw\\,ih)\\,iw*1080/ih\\,1920):if(gt(iw\\,ih)\\,1080\\,ih*1920/iw)")
		// args = append(args, "-vf", "scale=w=1920:h=1080:force_original_aspect_ratio=decrease")
	}
	// args = append(args, "-c:a", "libmp3lame")
	args = append(args, "-c:a", "aac")
	args = append(args, dst)
	cmd := exec.Command("ffmpeg", args...)
	if isHev1(src) {
		cmd = exec.Command("ffmpeg", "-i", src, "-c:v", "copy", "-tag:v", "hvc1", "-c:a", "aac", dst)
	}
	// 启动命令
	log.Printf("开始执行命令:%s\n", cmd.String())

	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("转换失败：%v\n", err)
		return
	} else {
		log.Printf("转换成功：%s\n", string(out))
		// 先尝试删除源文件
		if err := os.Remove(src); err != nil {
			log.Fatalf("删除源文件失败：%v\n", err)
		}
		// 源文件删除成功后，等待短暂时间确保文件句柄完全释放
		time.Sleep(100 * time.Millisecond)
		// 尝试重命名
		src = strings.Replace(src, filepath.Ext(src), ".mp4", 1)
		if err := os.Rename(dst, src); err != nil {
			log.Fatalf("重命名文件失败：%v\n", err)
		}
	}
}
