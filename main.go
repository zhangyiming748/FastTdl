package main

import (
	"fmt"
	"github.com/zhangyiming748/FastTdl/tdl"
	"github.com/zhangyiming748/FastTdl/util"
	"os"
	"path/filepath"
)

func main() {
	defer func() {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("无法获取用户的个人文件夹目录:", err)
		}
		dir := filepath.Join(home, "Downloads")
		if path := os.Getenv("TDL"); path != "" {
			dir = path
		}
		fmt.Printf("用户的个人文件夹目录: %s\n", home)
		fmt.Printf("用户的下载文件夹目录: %s\n", dir)
		target := filepath.Join(dir, "telegram")
		util.Chown(target, 1000, 1000)
	}()
	urls := util.ReadByLine("post.link")
	proxy := os.Getenv("PROXY")
	if proxy == "" {
		proxy = "http://127.0.0.1:8889"
	}
	tdl.DownloadsHelp(urls, proxy)
}
