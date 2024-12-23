package tdl

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zhangyiming748/FastTdl/constant"
	"github.com/zhangyiming748/FastTdl/util"
)

//	func GenerateDownloadLinkByOffset(of constant.OneFile) {
//		of.AddIdByOffset()
//	}
func GenerateDownloadLinkByCapacity(of constant.OneFile) (ofs []constant.OneFile) {
	c := of.Capacity
	for i := 0; i < c; i++ {
		nof := constant.OneFile{
			Channel:  of.Channel,
			FileId:   of.FileId + i,
			Tag:      of.Tag,
			Subtag:   of.Subtag,
			FileName: "",
			Offset:   0,
			Capacity: 0,
			Success:  false,
		}
		ofs = append(ofs, nof)
	}
	return ofs
}
func DownloadWithFolder(of constant.OneFile, proxy string) constant.OneFile {
	uri := strings.Join([]string{"https://t.me", of.Channel, strconv.Itoa(of.FileId)}, "/")
	fmt.Printf("用户的下载文件夹目录: %s\n", constant.GetMainFolder())
	fmt.Printf("要下载的链接: %s\n", uri)
	_, err := util.GetLevelDB().Get([]byte(uri), nil)
	if errors.Is(err, leveldb.ErrNotFound) {
		log.Println("文件未下载过")
		util.GetLevelDB().Put([]byte(uri), []byte("downloaded"), nil)
	} else {
		log.Println("文件下载过,跳过")
		of.SetStatus()
		return of
	}
	target := constant.GetMainFolder()
	if tag := of.Tag; tag != "" {
		target = filepath.Join(target, tag)
		if subtag := of.Subtag; subtag != "" {
			target = filepath.Join(target, subtag)
		}
	}
	os.MkdirAll(target, 0755)
	if err := util.ExecTdlCommand(proxy, uri, target); err != nil {
		log.Println("下载命令执行出错", uri)
		return of
	}
	of.SetStatus()
	if of.FileName != "" {
		util.RenameByKey(of)
	}
	return of
}
func ParseLines(lines []string, f *os.File) (ofs []constant.OneFile) {
	for _, line := range lines {
		if of, err := parseOneLine(line); err != nil { // 如果解析失败 则写入文件
			log.Printf("解析失败:%s\n", line)
			f.WriteString(line)
			f.WriteString("\n")
			continue
		} else {
			ofs = append(ofs, *of)
		}
	}
	return ofs
}
func parseOneLine(line string) (*constant.OneFile, error) {
	log.Printf("解析行: %s\n", line)
	of := new(constant.OneFile)
	line = strings.Replace(line, "?single", "", -1)
	if channel, id, err := getChannelAndFileID(line); err != nil {
		return nil, fmt.Errorf("URL: %s 不符合格式", line)
	} else {
		of.SetId(id)
		of.SetChannel(channel)
	}
	originUrl := strings.Join([]string{"https://t.me", of.Channel, strconv.Itoa(of.FileId)}, "/")
	params := strings.Replace(line, originUrl, "", 1)
	tag, subtag, filename, offset, capacity, err := getParam(params)
	if err != nil {
		return nil, err
	} else {
		of.SetTag(tag)
		of.SetSubtag(subtag)
		of.SetFileName(filename)
		of.SetOffset(offset)
		of.SetCapacity(capacity)
	}
	log.Printf("解析结果:%+v\n", of)
	return of, nil
}
func getChannelAndFileID(url string) (channel string, file int, err error) {
	//https://t.me/guoman_08/2148#&@+%
	static := "https://t.me/"
	url = strings.Replace(url, static, "", 1)
	if strings.Contains(url, "#") {
		prefix := strings.Split(url, "#")[0]
		channel = strings.Split(prefix, "/")[0]
		file, _ = strconv.Atoi(strings.Split(prefix, "/")[1])
	} else if strings.Contains(url, "&") {
		prefix := strings.Split(url, "&")[0]
		channel = strings.Split(prefix, "/")[0]
		file, _ = strconv.Atoi(strings.Split(prefix, "/")[1])
	} else if strings.Contains(url, "@") {
		prefix := strings.Split(url, "@")[0]
		channel = strings.Split(prefix, "/")[0]
		file, _ = strconv.Atoi(strings.Split(prefix, "/")[1])
	} else if strings.Contains(url, "+") {
		prefix := strings.Split(url, "+")[0]
		channel = strings.Split(prefix, "/")[0]
		file, _ = strconv.Atoi(strings.Split(prefix, "/")[1])
	} else if strings.Contains(url, "%") {
		prefix := strings.Split(url, "%")[0]
		channel = strings.Split(prefix, "/")[0]
		file, _ = strconv.Atoi(strings.Split(prefix, "/")[1])
	} else {
		channel = strings.Split(url, "/")[0]
		file, _ = strconv.Atoi(strings.Split(url, "/")[1])
	}
	return channel, file, nil
}
func getParam(input string) (tag, subtag, filename string, offset, capacity int, err error) {
	/*
		因为 %或+后面不可能再出现其他参数了，这两个属性也不能同时存在，所以单独处理
	*/
	if strings.Contains(input, "%") { //包含容量
		capacity, err = strconv.Atoi(strings.Split(input, "%")[1])
		if err != nil {
			return "", "", "", 0, 0, err
		}
		input = strings.Split(input, "%")[0]
	}
	if strings.Contains(input, "+") { //包含偏移量
		offset, err = strconv.Atoi(strings.Split(input, "+")[1])
		if err != nil {
			return "", "", "", 0, 0, err
		}
		input = strings.Split(input, "+")[0]
	}
	if strings.Contains(input, "@") {
		filename = strings.Split(input, "@")[1]
		input = strings.Split(input, "@")[0]
	}
	if strings.Contains(input, "&") {
		subtag = strings.Split(input, "&")[1]
		input = strings.Split(input, "&")[0]
	}
	if strings.Contains(input, "#") {
		tag = strings.Split(input, "#")[1]
		input = strings.Split(input, "#")[0]
	}
	fmt.Printf("解析参数后剩下的内容:%s\n", input)
	return tag, subtag, filename, offset, capacity, nil
}
