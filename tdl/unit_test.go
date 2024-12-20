package tdl

import (
	"fmt"
	"testing"
)

func TestGetParam(t *testing.T) {
	urls := []string{
		"https://t.me/TNTsex/27584",
		"https://t.me/TNTsex/27584#主文件夹",
		"https://t.me/TNTsex/27584&子文件夹",
		"https://t.me/TNTsex/27584@文件名",
		"https://t.me/TNTsex/27584+偏移量",
		"https://t.me/TNTsex/27584%容量",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹",
		"https://t.me/TNTsex/27584#主文件夹@文件名",
		"https://t.me/TNTsex/27584#主文件夹+偏移量",
		"https://t.me/TNTsex/27584#主文件夹%容量",
		"https://t.me/TNTsex/27584&子文件夹@文件名",
		"https://t.me/TNTsex/27584&子文件夹+偏移量",
		"https://t.me/TNTsex/27584&子文件夹%容量",
		"https://t.me/TNTsex/27584@文件名+偏移量",
		"https://t.me/TNTsex/27584@文件名%容量",
		"https://t.me/TNTsex/27584+偏移量%容量",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹+偏移量",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹%容量",
		"https://t.me/TNTsex/27584#主文件夹@文件名+偏移量",
		"https://t.me/TNTsex/27584#主文件夹@文件名%容量",
		"https://t.me/TNTsex/27584#主文件夹+偏移量%容量",
		"https://t.me/TNTsex/27584&子文件夹@文件名+偏移量",
		"https://t.me/TNTsex/27584&子文件夹@文件名%容量",
		"https://t.me/TNTsex/27584&子文件夹+偏移量%容量",
		"https://t.me/TNTsex/27584@文件名+偏移量%容量",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名+偏移量",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名%容量",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹+偏移量%容量",
		"https://t.me/TNTsex/27584#主文件夹@文件名+偏移量%容量",
		"https://t.me/TNTsex/27584&子文件夹@文件名+偏移量%容量",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名+偏移量%容量",
	}
	for i, url := range urls {
		line, err := parseOneLine(url)
		if err != nil {
			continue
		}
		fmt.Printf("%d\n%+v\n", i+1, line)
	}
}
