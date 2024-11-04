package tdl

import (
	"fmt"
	"github.com/zhangyiming748/FastTdl/util"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func DownloadsHelp(urls []string, proxy string) {
	var status string
	var count int
	defer func() {
		status = fmt.Sprintf("全部下载结束,失败 %d / %d 个\n", count, len(urls))
		log.Println(status)
	}()
	f, err := os.OpenFile("failed.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, url := range urls {
		if strings.Contains(url, " ") { //如果url包含空格 需要循环判下载
			base := strings.Split(url, " ")[0]                  //https://t.me/acgr18/34406 3
			step, _ := strconv.Atoi(strings.Split(url, " ")[1]) //9
			//https://t.me/acgr18/34406 9
			prefix, suffix, _ := Split(base)
			var uris []string
			for i := 0; i < step; i++ {
				uri := strings.Join([]string{prefix, strconv.Itoa(suffix + i)}, "/")
				log.Printf("Add url %s with %s\n", uri, proxy)
				uris = append(uris, uri)
			}
			Downloads(uris, proxy, f)
		} else { //如果url不含空格
			Downloads(urls, proxy, f)
		}
	}
	f.Sync()
}
func Downloads(urls []string, proxy string, f *os.File) {
	var status string
	var count int
	defer func() {
		status = fmt.Sprintf("全部下载结束,失败 %d / %d 个\n", count, len(urls))
		log.Println(status)
	}()
	for _, url := range urls {
		if strings.Contains(url, "@") {
			if strings.Contains(url, "#") {
				base := strings.Split(url, "#")[0] //https://t.me/acgr18/34406#3434@feef
				dir := strings.Split(url, "#")[1]
				dir = strings.Split(dir, "@")[0]
				fname := strings.Split(url, "@")[1]
				fail := DownloadWithFolder(base, proxy, dir)
				if fail != nil {
					count++
					out := fmt.Sprintf("download fail :%s\n", url)
					f.WriteString(out)
				}
				if key := findKeyByUrl(url); key != "" {
					util.RenameByKey(key, fname)
				}
			} else {
				base := strings.Split(url, "@")[0] //https://t.me/acgr18/34406@feef
				fname := strings.Split(url, "@")[1]
				fail := Download(base, proxy)
				if fail != nil {
					count++
					out := fmt.Sprintf("download fail :%s\n", url)
					f.WriteString(out)
				}
				if key := findKeyByUrl(url); key != "" {
					util.RenameByKey(key, fname)
				}
			}
		} else if strings.Contains(url, "#") {
			base := strings.Split(url, "#")[0] //https://t.me/acgr18/34406
			dir := strings.Split(url, "#")[1]
			fail := DownloadWithFolder(base, proxy, dir)
			if fail != nil {
				count++
				out := fmt.Sprintf("download fail :%s\n", url)
				f.WriteString(out)
			}
		} else {
			fail := Download(url, proxy)
			if fail != nil {
				count++
				out := fmt.Sprintf("download fail :%s\n", url)
				f.WriteString(out)
			}
		}
	}
	f.Sync()
}

// https://github.com/iyear/tdl.git
func Download(uri, proxy string) error {
	var status string
	defer func() {
		log.Println(status)
	}()
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取用户的个人文件夹目录:", err)
		return err
	}
	dir := filepath.Join(home, "Downloads")
	if path := os.Getenv("TDL"); path != "" {
		dir = path
	}
	fmt.Printf("用户的个人文件夹目录: %s\n", home)
	fmt.Printf("用户的下载文件夹目录: %s\n", dir)
	target := filepath.Join(dir, "telegram")
	os.MkdirAll(target, 0755)
	err = util.ExecTdlCommand(proxy, uri, target)
	if err != nil {
		log.Println("下载命令执行出错", uri)
		status = strings.Join([]string{status, "下载失败"}, "")
		return err
	} else {
		status = strings.Join([]string{status, "下载成功"}, "")
		return nil
	}
}

func DownloadWithFolder(uri, proxy, fname string) error {
	var status string
	defer func() {
		log.Println(status)
	}()
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取用户的个人文件夹目录:", err)
		return err
	}
	dir := filepath.Join(home, "Downloads")
	if path := os.Getenv("TDL"); path != "" {
		dir = path
	}
	fmt.Printf("用户的个人文件夹目录: %s\n", home)
	fmt.Printf("用户的下载文件夹目录: %s\n", dir)
	target := filepath.Join(dir, "telegram", fname)
	os.MkdirAll(target, 0755)
	err = util.ExecTdlCommand(proxy, uri, target)
	if err != nil {
		log.Println("下载命令执行出错", uri)
		status = strings.Join([]string{status, "下载失败"}, "")
		return err
	} else {
		status = strings.Join([]string{status, "下载成功"}, "")
		return nil
	}
}

func Split(s string) (prefix string, suffix int, err error) {
	lastSlashIndex := strings.LastIndex(s, "/")
	if lastSlashIndex != -1 {
		// 分割字符串
		beforeLastSlash := s[:lastSlashIndex]
		afterLastSlash, _ := strconv.Atoi(s[lastSlashIndex+1:])
		return beforeLastSlash, afterLastSlash, nil
	} else {
		return "", -1, err
	}
}
func findKeyByUrl(u string) string {
	//u = "https://t.me/FFLL05/57137?single" // 你可以替换为其他 URL 进行测试
	var prefix string
	if strings.Contains(u, "?") {
		prefix = strings.Split(u, "?")[0]
	} else if strings.Contains(u, "#") {
		prefix = strings.Split(u, "#")[0]
	} else if strings.Contains(u, "@") {
		prefix = strings.Split(u, "@")[0]
	}
	lastSlashIndex := strings.LastIndex(prefix, "/")
	if lastSlashIndex == -1 {
		fmt.Println("没有找到 '/'")
		return ""
	}
	// 提取最后一个 '/' 之后的部分
	lastPart := prefix[lastSlashIndex+1:]
	// 如果需要确保只提取数字，可以使用 strings.TrimSpace 去除空格
	// 这里假设最后一部分是数字
	fmt.Println("提取的数字:", lastPart)
	return lastPart
}
