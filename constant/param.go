package constant

import (
	"fmt"
	"net"
	"net/url"
)

type Parameter struct {
	Proxy      string
	MainFolder string
	SMTP       string
}

func (p *Parameter) SetProxy(s string) {
	p.Proxy = s
}

func (p *Parameter) GetProxy() string {
	return p.Proxy
}

func (p *Parameter) SetMainFolder(s string) {
	p.MainFolder = s
}

func (p *Parameter) GetMainFolder() string {
	return p.MainFolder
}

func (p *Parameter) SetSMTP(s string) {
	p.SMTP = s
}

func (p *Parameter) GetSMTP() string {
	return p.SMTP
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

func Ping(proxy string) error {
	u, err := url.Parse(proxy)
	if err != nil {
		fmt.Println("解析URL失败:", err)
		return err
	}
	ip := u.Hostname()
	port := u.Port()

	address := net.JoinHostPort(ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
