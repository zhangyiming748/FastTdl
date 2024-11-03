package logic

import (
	"fmt"
	"github.com/zhangyiming748/basicGin/util"
	"os"
	"testing"
)

func init() {
	util.SetLog("telegram.log")
}
func TestPATH(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取用户的个人文件夹目录:", err)
	}
	t.Log(home)
}

func TestDownloadsTelegram(t *testing.T) {
	urls := util.ReadByLine("post.link")
	proxy := "http://127.0.0.1:8889"
	Downloads(urls, proxy)
}

func TestFindKey(t *testing.T) {
	u := "https://t.me/acgr18/34406?re#3434@feef"
	ret := findKeyByUrl(u)
	t.Log(ret)
}
func TestUni(t *testing.T) {
	fmt.Printf("\u808F")
	fmt.Printf("\u5C44")
}
