package main

import (
	"fmt"
	"log"
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
		proxy = "http://127.0.0.1:8889"
	}
	tdl.DownloadsHelp(urls, proxy)
}
