package main

import (
	"fmt"
	"log"
	"net"
	uri "net/url"
	"os"
	"runtime"

	"github.com/zhangyiming748/FastTdl/tdl"
	"github.com/zhangyiming748/FastTdl/util"
)

func init() {
	util.SetLog("tdl.log")
	util.SetLevelDB()
}

func main() {
	var urls []string
	if FileExists("/data/post.link") {
		urls = util.ReadByLine("/data/post.link")
	} else if FileExists("post.link") {
		urls = util.ReadByLine("post.link")
	} else {
		log.Fatalln("没有在任何位置找到post.link文件")
	}
	proxy := os.Getenv("PROXY")
	if proxy == "" && runtime.GOOS == "linux" {
		log.Fatalln("容器中未指定外部可用代理")
	}
	if proxy == "" {
		proxy = "http://127.0.0.1:8889"
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
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
