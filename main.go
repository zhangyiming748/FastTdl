package main

import (
	"github.com/zhangyiming748/FastTdl/tdl"
	"github.com/zhangyiming748/FastTdl/util"
	"os"
)

func main() {
	urls := util.ReadByLine("post.link")
	proxy := os.Getenv("PROXY")
	if proxy == "" {
		proxy = "http://127.0.0.1:8889"
	}
	tdl.DownloadsHelp(urls, proxy)
}
