package discussions

import (
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zhangyiming748/FastTdl/util"
)

func Discussions(uri, proxy string) {
    dir, err := os.UserHomeDir()
    if err != nil {
        log.Printf("获取用户目录失败: %v\n", err)
        return
    }
	dir = filepath.Join(dir, "media")
	if strings.Contains(uri, "#") {
		uri = strings.Split(uri, "#")[0]
		dir = filepath.Join(dir, strings.Split(uri, "#")[1])
	}
    channelId, chatId := parseUrl(uri)
    if channelId == "" || chatId == "" {
        log.Printf("解析URL失败: %s\n", uri)
        return
    }
	getJson(channelId, chatId, dir, proxy)
	downloadDiscussions(dir, proxy)
}

func parseUrl(uri string) (string, string) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", ""
	}
	// 从路径中提取 channel ID
	path := strings.TrimPrefix(u.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		channelId := parts[0]
		log.Printf("Channel ID: %s", channelId)
		chatID := parts[1]
		return channelId, chatID
	} else {
		log.Println("Invalid URL")
		return "", ""
	}
}

func getJson(channelId, chatId, dir, proxy string) {
	// tdl chat export -c 2386513401 --proxy http://127.0.0.1:8889  --reply 14288
	exec.Command("tdl", "chat", "export", "-c", channelId, "--dir", dir, "--proxy", proxy, "--reply", chatId).Run()
}

func downloadDiscussions(dir, proxy string) {
	// tdl download --proxy http://127.0.0.1:8889 --threads 8 --file tdl-export.json
	file := filepath.Join(dir, "tdl-export.json")
	cmd := exec.Command("tdl", "download", "--proxy", proxy, "--threads", "8", "--file", file)
	if err:=util.ExecCommand(cmd);err==nil{
		os.RemoveAll(file)
	}
}
