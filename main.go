package main

import (
	"fmt"
	"github.com/zhangyiming748/FastTdl/constant"
	"log"
	"net"
	uri "net/url"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/zhangyiming748/FastTdl/tdl"
	"github.com/zhangyiming748/FastTdl/util"
)

func init() {
	util.SetLog("tdl.log")
	util.SetLevelDB()
}

func main() {
	summaries := []constant.OneFile{}
	failed, err := os.OpenFile("failed.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer failed.Close()
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
	links := tdl.ParseLines(urls, failed)
	failed.Sync()
	for _, link := range links {
		if link.Offset != 0 && link.Capacity == 0 {
			link.Id += link.Offset
			summary := tdl.DownloadWithFolder(link, proxy)
			summaries = append(summaries, summary)
		} else if link.Offset == 0 && link.Capacity != 0 {
			us := tdl.GenerateDownloadLinkByCapacity(link)
			for _, u := range us {
				summary := tdl.DownloadWithFolder(u, proxy)
				summaries = append(summaries, summary)
			}
		} else {
			summary := tdl.DownloadWithFolder(link, proxy)
			summaries = append(summaries, summary)
		}
	}
	for i, status := range summaries {
		if status.Success {
			log.Printf("第%d个文件下载成功\n", i+1)
		} else {
			log.Printf("第%d个文件%+v下载失败\n", i+1, status)
			failed.WriteString(fmt.Sprintf("%+v\n", strings.Join([]string{"https://t.me", status.Channel, strconv.Itoa(status.Id)}, "/")))
			failed.Sync()
		}

	}
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
