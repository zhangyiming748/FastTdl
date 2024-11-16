package main

import (
	"fmt"
	"log"
	"net"
	uri "net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/zhangyiming748/FastTdl/tdl"
	"github.com/zhangyiming748/FastTdl/util"
)

func init() {
	util.SetLog("tdl.log")
	util.SetLevelDB()
}

func main() {
	defer func() {
		if runtime.GOOS == "linux" {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("无法获取用户的个人文件夹目录:", err)
			}
			dir := filepath.Join(home, "Downloads")
			if path := os.Getenv("TDL"); path != "" {
				dir = path
			}
			target := filepath.Join(dir, "telegram")
			log.Printf("最终下载的文件夹: %s\n", target)
			util.Chown(target, 1000, 1000)
		}
	}()
	urls := util.ReadByLine("post.link")
	proxy := os.Getenv("PROXY")
	if proxy == "" {
		log.Fatalln("容器中未指定外部可用代理")
	}
	if err := ping(proxy); err != nil {
		log.Fatalf("指定的代理IP地址不可用,错误信息:%v\n", err)
	}
	tdl.DownloadsHelp(urls, proxy)
}

func ping(proxy string) error {
	u, err := uri.Parse(proxy)
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
