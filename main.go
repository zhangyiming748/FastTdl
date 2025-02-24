package main

import (
	"fmt"
	"log"
	
	"os"
	"strconv"
	"strings"

	"github.com/zhangyiming748/FastTdl/discussions"
	"github.com/zhangyiming748/FastTdl/model"

	"github.com/zhangyiming748/FastTdl/archive"
	"github.com/zhangyiming748/FastTdl/constant"
	"github.com/zhangyiming748/FastTdl/mysql"

	"github.com/zhangyiming748/FastTdl/tdl"
	"github.com/zhangyiming748/FastTdl/util"
)

func init() {
	util.SetLog("tdl.log")
	util.SetLevelDB()
	mysql.SetMysql()
	if mysql.UseMysql() {
		mysql.GetMysql().Sync(model.File{})
	}
}

type Info struct {
	URL  string
	Base constant.OneFile
}

func main() {
	summaries := []constant.OneFile{}
	failed, err := os.OpenFile("failed.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer failed.Close()
	p:=constant.GetParams()
	
	defer discussions.DownloadAllDiscussions(p.GetProxy())
	var urls []string
	if util.IsExistFile("/data/post.link") {
		urls = util.ReadByLine("/data/post.link")
	} else if util.IsExistFile("post.link") {
		urls = util.ReadByLine("post.link")
	} else {
		log.Println("没有在任何位置找到post.link文件")
	}

	links := tdl.ParseLines(urls, failed)
	failed.Sync()
	//var current Info
	defer archive.Archive()
	for index, link := range links {
		log.Printf("开始下载第%d/%d个文件\n", index+1, len(links))
		//current.URL = strings.Join([]string{"https://t.me", link.Channel, strconv.Itoa(link.FileId)}, "/")
		//current.Base = link
		
		if link.Offset != 0 && link.Capacity == 0 {
			link.FileId += link.Offset
			summary := tdl.DownloadWithFolder(link, p.GetProxy(), failed)
			summaries = append(summaries, summary)
		} else if link.Offset == 0 && link.Capacity != 0 {
			us := tdl.GenerateDownloadLinkByCapacity(link)
			for _, u := range us {
				summary := tdl.DownloadWithFolder(u, p.GetProxy(), failed)
				summaries = append(summaries, summary)
			}
		} else {
			summary := tdl.DownloadWithFolder(link, p.GetProxy(), failed)
			summaries = append(summaries, summary)
		}
		log.Printf("下载完成第个文件%d/%d\n", index, len(links))
	}
	for i, status := range summaries {
		if status.Success {
			log.Printf("第%d个文件下载成功\n", i+1)
		} else {
			log.Printf("第%d个文件%+v下载失败\n", i+1, status)
			failed.WriteString(fmt.Sprintf("%+v\n", strings.Join([]string{"https://t.me", status.Channel, strconv.Itoa(status.FileId)}, "/")))
			failed.Sync()
		}
	}
}

