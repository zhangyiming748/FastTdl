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
	"github.com/zhangyiming748/FastTdl/constant"
)

// 音频处理相关的常量定义
const (
	// AudioBookType 有声书类型标识
	AudioBookType = "audiobook"
	// RapMusicType 说唱音乐类型标识
	RapMusicType = "rap"
	//AudioBook = "1.54" //等效audition的65%
	// Speed 音频播放速度，等效audition的70%
	Speed = "1.43"
	// Volume 音频音量增益值
	Volume = "3.0"
)





// ArchiveAudio 处理指定类型的音频文件
// mytype 参数指定音频类型，可以是AudioBookType或RapMusicType
func ArchiveAudio(mytype string) {
	p := constant.GetParams()
	files, _ := GetAllAudioFiles(p.GetMainFolder())
	for _, file := range files {
		ConvertAudio(file, mytype)
	}
}

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
	file, _ := os.Open(fp)
	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)
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
	ff := audition2ffmpeg(Speed)
	atempo := strings.Join([]string{"atempo", ff}, "=")
	volume := strings.Join([]string{"volume", Volume}, "=")
	filter := strings.Join([]string{atempo, volume}, ",")

	// 根据音频类型设置不同的处理参数
	switch mytype {
	case AudioBookType:
		// 有声书加速65% 电平增加
		args = append(args, "-filter:a", filter)
		args = append(args, "-c:a", "aac")
	// 歌曲类只增加电平
	case RapMusicType:
		args = append(args, "-filter:a", volume)
		args = append(args, "-c:a", "aac")
	}

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

	// 在后台处理标准输出
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

	// 等待命令完成并处理结果
	if err := cmd.Wait(); err != nil {
		log.Printf("转换失败：%v\n", err)
		return
	} else {
		// 先尝试删除源文件
		if err := os.Remove(src); err != nil {
			log.Printf("删除源文件失败：%v\n", err)
			return
		}
		// 源文件删除成功后，等待短暂时间确保文件句柄完全释放
		time.Sleep(100 * time.Millisecond)
		// 尝试重命名
		if err := os.Rename(dst, src); err != nil {
			log.Printf("重命名文件失败：%v\n", err)
			return
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
