// Package archive 提供音频文件处理和转换的功能，包括音频文件的检测、转换和存档
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
)

// 音频处理相关的常量定义
const (
	// AudioBookType 有声书类型标识
	AudioBookType = "audiobook"
	// RapMusicType 说唱音乐类型标识
	RapMusicType = "rap"
	// Speed        = "1.54" //等效audition的65%
	// Speed = "1.43" 音频播放速度，等效audition的70%
	// Volume 音频音量增益值
	Volume = "2.7"
)

// ArchiveAudio 处理指定类型的音频文件
// mytype 参数指定音频类型，可以是AudioBookType或RapMusicType
//func ArchiveAudio(mytype string) {
//	p := constant.GetParams()
//	files, _ := GetAllAudioFiles(p.GetMainFolder())
//	for _, file := range files {
//		ConvertAudio(file, mytype)
//	}
//}

// GetAllAudioFiles 返回指定目录下所有音频文件的绝对路径
// root 为要搜索的根目录
// 返回音频文件路径列表和可能的错误
func GetAllAudioFiles(root string) ([]string, error) {
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
			if isAudio(absPath) {
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

// isAudio 检查文件是否为音频文件
// fp 为文件路径
// 返回布尔值表示是否为音频文件
func isAudio(fp string) bool {
	file, err := os.Open(fp)
	if err != nil {
		return false
	}
	defer file.Close()
	head := make([]byte, 261)
	if _, err := file.Read(head); err != nil {
		return false
	}
	return filetype.IsAudio(head)
}

// ConvertAudio 转换音频文件
// src 为源文件路径
// mytype 为音频类型，决定处理方式
func ConvertAudio(src, mytype string) {
	if !isAudio(src) {
		return
	}
	// 生成临时文件路径
	purgePath := filepath.Dir(src)
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	b := seed.Intn(2000)
	tmp := strconv.Itoa(b)
	tmp = strings.Join([]string{tmp, ".mp3"}, "")
	dst := filepath.Join(purgePath, tmp)

	// 构建ffmpeg命令参数
	args := []string{"-i", src}
	ff := audition2ffmpeg("65")
	atempo := strings.Join([]string{"atempo", ff}, "=")
	volume := strings.Join([]string{"volume", Volume}, "=")
	filter := strings.Join([]string{atempo, volume}, ",")

	args = append(args, "-ac", "1")
	args = append(args, "-map_metadata", "-1")
	args = append(args, "-ar", "44100")
	args = append(args, "-ab", "128k")
	// 根据音频类型设置不同的处理参数
	switch mytype {
	case AudioBookType:
		// 有声书加速65% 电平增加
		args = append(args, "-filter:a", filter)
	// 歌曲类只增加电平
	case RapMusicType:
		args = append(args, "-filter:a", volume)
	default:
		// 其他类型
		args = append(args, "-c:a", "aac")
	}
	args = append(args, dst)
	cmd := exec.Command("ffmpeg", args...)

	// 获取输出和错误管道

	// 等待命令完成并处理结果
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("转换失败：%v\n", err)
	} else {
		fmt.Printf("转换成功：%s\n", string(out))
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

// audition2ffmpeg 将Adobe Audition的速度参数转换为ffmpeg的速度参数
// speed 为输入的速度参数
// 返回转换后的ffmpeg速度参数
func audition2ffmpeg(speed string) string {
	audition, err := strconv.ParseFloat(speed, 64)
	if err != nil {
		log.Fatalf("解析加速参数错误:%v,退出程序", err)
	}
	param := 100 / audition
	log.Printf("转换后的原始参数:%v\n", param)
	final := fmt.Sprintf("%.2f", param)
	log.Printf("保留两位小数的原始参数:%v\n", final)
	return final
}
