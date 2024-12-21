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
		"https://t.me/TNTsex/27584+3",
		"https://t.me/TNTsex/27584%10",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹",
		"https://t.me/TNTsex/27584#主文件夹@文件名",
		"https://t.me/TNTsex/27584#主文件夹+3",
		"https://t.me/TNTsex/27584#主文件夹%10",
		"https://t.me/TNTsex/27584&子文件夹@文件名",
		"https://t.me/TNTsex/27584&子文件夹+3",
		"https://t.me/TNTsex/27584&子文件夹%10",
		"https://t.me/TNTsex/27584@文件名+3",
		"https://t.me/TNTsex/27584@文件名%10",
		"https://t.me/TNTsex/27584+3%10",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹+3",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹%10",
		"https://t.me/TNTsex/27584#主文件夹@文件名+3",
		"https://t.me/TNTsex/27584#主文件夹@文件名%10",
		"https://t.me/TNTsex/27584#主文件夹+3%10",
		"https://t.me/TNTsex/27584&子文件夹@文件名+3",
		"https://t.me/TNTsex/27584&子文件夹@文件名%10",
		"https://t.me/TNTsex/27584&子文件夹+3%10",
		"https://t.me/TNTsex/27584@文件名+3%10",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名+3",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名%10",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹+3%10",
		"https://t.me/TNTsex/27584#主文件夹@文件名+3%10",
		"https://t.me/TNTsex/27584&子文件夹@文件名+3%10",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名+3%10",
	}
	for _, url := range urls {

		_, err := parseOneLine(url)
		if err != nil {
			continue
		}
	}
}
func TestSuffix(t *testing.T) {
	urls := []string{
		"https://t.me/TNTsex/27584",
		"https://t.me/TNTsex/27584#主文件夹",
		"https://t.me/TNTsex/27584&子文件夹",
		"https://t.me/TNTsex/27584@文件名",
		"https://t.me/TNTsex/27584+3",
		"https://t.me/TNTsex/27584%10",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹",
		"https://t.me/TNTsex/27584#主文件夹@文件名",
		"https://t.me/TNTsex/27584#主文件夹+3",
		"https://t.me/TNTsex/27584#主文件夹%10",
		"https://t.me/TNTsex/27584&子文件夹@文件名",
		"https://t.me/TNTsex/27584&子文件夹+3",
		"https://t.me/TNTsex/27584&子文件夹%10",
		"https://t.me/TNTsex/27584@文件名+3",
		"https://t.me/TNTsex/27584@文件名%10",
		"https://t.me/TNTsex/27584+3%10",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹+3",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹%10",
		"https://t.me/TNTsex/27584#主文件夹@文件名+3",
		"https://t.me/TNTsex/27584#主文件夹@文件名%10",
		"https://t.me/TNTsex/27584#主文件夹+3%10",
		"https://t.me/TNTsex/27584&子文件夹@文件名+3",
		"https://t.me/TNTsex/27584&子文件夹@文件名%10",
		"https://t.me/TNTsex/27584&子文件夹+3%10",
		"https://t.me/TNTsex/27584@文件名+3%10",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名+3",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名%10",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹+3%10",
		"https://t.me/TNTsex/27584#主文件夹@文件名+3%10",
		"https://t.me/TNTsex/27584&子文件夹@文件名+3%10",
		"https://t.me/TNTsex/27584#主文件夹&子文件夹@文件名+3%10",
	}
	for i, url := range urls {

		if prefix, capacity := getCapacity(url); capacity != 0 {
			fmt.Printf("%d\n%+v\n%v\n", i+1, prefix, capacity)
		}
		if prefix, offset := getOffset(url); offset != 0 {
			fmt.Printf("%d\n%+v\n%v\n", i+1, prefix, offset)
		}
	}
}
