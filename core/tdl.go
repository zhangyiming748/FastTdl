package core

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"FastTdl/archive"
	"FastTdl/constant"
	"FastTdl/discussions"
	"FastTdl/model"
	"FastTdl/sqlite"
	"FastTdl/tdl"
	"FastTdl/util"
)

var Home string

func init() {
	util.SetLog("tdl.log")
}

func Tdl(mainFolder, postLink, proxy string) {
	if proxy == "" {
		proxy = constant.DEFAULT_PROXY
	}
	//在这之前mainFolder实际上是下载链接文件的路径
	//mainFolder = filepath.Base(mainFolder)
	p := constant.Parameter{
		Proxy:      proxy,
		MainFolder: mainFolder,
		SMTP:       constant.DEFAULT_SMTP,
	}
	if err := isValidParent(p.MainFolder); err != nil {
		log.Fatalf("目前设置的主目录是一个不存在的目录:%s需要再检查一遍:%v\n", p.MainFolder, err)
	}
	if e := constant.Ping(p.Proxy); e != nil {
		log.Fatalf("代理设置错误:%s", e.Error())
	}
	// 添加这一行来初始化MySQL连接
	sqlite.SetSqlite()
	sqlite.GetSqlite().AutoMigrate(&model.Channel{})
	sqlite.GetSqlite().AutoMigrate(&model.File{})
	var urls []string
	urls = util.ReadByLine(postLink)
	for i, url := range urls {
		if strings.Contains(url, "comment") {
			log.Printf("检测到链接中包含comment,可能是评论链接，整体下载\n")
			discussions.Discussion(url, p)
			// 这里下载之后 从url列表里删除这个元素
			// urls = append(urls[:i], urls[i+1:]...)
			continue
		}
		if strings.Contains(url, "@") {
			parts := strings.Split(url, "@")
			if len(parts) > 1 {
				fileName := parts[1]
				if util.IsExistFile(filepath.Join(Home, "Downloads", fileName)) {
					log.Printf("文件名 %s 已经存在，可能会导致下载失败，请检查该文件名是否重复\n", fileName)
					// 从urls中删除这个元素
					urls = append(urls[:i], urls[i+1:]...)
				}
			}
		}
		links := tdl.ParseLines(urls)
		log.Printf("共有%d个文件\n", len(links))

		for index, link := range links {
			// 创建一个带取消功能的上下文，用于控制超时协程
			ctx, cancel := context.WithCancel(context.Background())

			// 启动超时监控协程
			go func(ctx context.Context, index int) {
				// 每30秒警告一次
				ticker := time.NewTicker(60 * time.Second)
				defer ticker.Stop()

				// 从60秒开始计数
				elapsed := 60

				for {
					select {
					case <-ticker.C:
						// 每60秒打印一次警告
						msg := fmt.Sprintf("警告：第%d个文件下载已超过%d秒\n", index+1, elapsed)
						util.TimeoutAlert(msg)
						elapsed += 60
					case <-ctx.Done():
						// 收到取消信号，正常退出
						return
					}
				}
			}(ctx, index)
			log.Printf("开始下载第%d/%d个文件\n", index+1, len(links))
			if link.Offset != 0 && link.Capacity == 0 {
				link.FileId += link.Offset
				tdl.DownloadWithFolder(link, p)

			} else if link.Offset == 0 && link.Capacity != 0 {
				us := tdl.GenerateDownloadLinkByCapacity(link)
				for _, u := range us {
					tdl.DownloadWithFolder(u, p)
				}
			} else {
				tdl.DownloadWithFolder(link, p)
			}

			// 下载完成后，取消超时监控协程
			cancel()

			log.Printf("下载完成第个文件%d/%d\n", index, len(links))
			if index == len(links)-1 {
				log.Println("所有文件下载完成")
				return
			}
		}
	}
}

// isValidParent 检查给定的文件夹路径是否有效
// 参数 folderPath: 需要验证的文件夹路径
// 返回值: 如果路径有效返回nil，否则返回相应的错误信息
func isValidParent(folderPath string) error {
	// 获取文件所在目录路径
	dirPath := filepath.Dir(folderPath)

	// 检查目录是否存在
	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("目录不存在: %s", dirPath)
		}
		return fmt.Errorf("检查目录时发生错误: %v", err)
	}
	return nil
}
func ArchiveAllFiles(roots ...string) {
	for _, root := range roots {
		archive.Videos(root)
		archive.Images(root)
	}
}
