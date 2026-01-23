package tdl

import (
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
	t.Log(urls)
}
func TestZh2En(t *testing.T) {
	m := zh2en("../zh_cn2en_us.md")

	for k, v := range m {
		t.Logf("\nkey = %s\nvalue = %s\n", k, v)
	}
}
