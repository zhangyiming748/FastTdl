package test

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/zhangyiming748/FastTdl/util"
)

func TestFingNewestFile(t *testing.T) {
	ret := util.FindLatestFileGo("/Users/zen/github/FastTdl")
	t.Logf("最终结果:%+v", ret)
}

// 生成连续的n个下载链接
func Generate(base string, n int) (links []string) {
	//https://t.me/rewu8666/137142
	// 根据最后一个出现的/分割字符串prefix和suffix
	prefix := base[:strings.LastIndex(base, "/")+1]
	suffix := base[strings.LastIndex(base, "/")+1:]
	baseindex, err := strconv.Atoi(suffix)
	if err != nil {
		log.Fatalf("base:%s 转换失败", base)
	}
	for i := 0; i < n; i++ {
		baseindex += 1
		link := fmt.Sprintf("%s%d", prefix, baseindex)
		log.Printf("生成链接:%s", link)
		links = append(links, link)
	}
	return links
}
func TestGenerate(t *testing.T) {
	links := Generate("https://t.me/rewu8666/136970", 94)
	f, err := os.OpenFile("post.link", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("打开文件失败:%s", err)
	}
	for _, link := range links {
		f.WriteString(fmt.Sprintln(link))
	}
}
