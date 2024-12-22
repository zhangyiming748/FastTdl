package constant

import (
	"fmt"
	"os"
	"path/filepath"
)

var proxy = "192.168.1.35:8889"

func GetProxy() string {
	return proxy
}

func SetProxy(p string) {
	proxy = p
}

var (
	MainFolder string
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("无法获取用户的个人文件夹目录:%v", err))
	}
	MainFolder = filepath.Join(home, "Downloads", "media")
	if p := os.Getenv("TDL"); p != "" {
		MainFolder = p
	}
}
func GetMainFolder() string {
	return MainFolder
}

type OneFile struct {
	Channel  string // 频道id
	FileId   int    // 文件id
	Tag      string // 主文件夹名 #后面的文件名
	Subtag   string // 子(二级)文件夹名 &后面的文件名
	FileName string // 手动设置的文件名 @后面的文件名
	Offset   int    // 偏移量 如下载当前媒体之后第n个文件 +后面的数字
	Capacity int    // 下载当前文件和之后的n个文件 %后面的数字
	Success  bool   // 是否下载成功
}

func (f *OneFile) SetChannel(s string) {
	f.Channel = s
}

func (f *OneFile) SetId(i int) {
	f.FileId = i
}

func (f *OneFile) SetTag(s string) {
	f.Tag = s
}

func (f *OneFile) SetSubtag(s string) {
	f.Subtag = s
}

func (f *OneFile) SetFileName(s string) {
	f.FileName = s
}

func (f *OneFile) SetOffset(i int) {
	f.Offset = i
}

func (f *OneFile) SetCapacity(i int) {
	f.Capacity = i
}

func (f *OneFile) SetStatus() {
	f.Success = true
}
