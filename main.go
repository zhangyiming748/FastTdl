package logic

import (
	"github.com/zhangyiming748/FastTdl/util"
	"os"
	"testing"
)

func main(t *testing.T) {
	urls := util.ReadByLine("post.link")
	proxy := os.Getenv("PROXY")
	if proxy == "" {
		proxy = "http://127.0.0.1:8889"
	}
	Downloads(urls, proxy)
}
