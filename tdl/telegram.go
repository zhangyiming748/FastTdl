package tdl

import (
	"errors"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zhangyiming748/FastTdl/constant"
	"github.com/zhangyiming748/FastTdl/util"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//	func GenerateDownloadLinkByOffset(of constant.OneFile) {
//		of.AddIdByOffset()
//	}
func GenerateDownloadLinkByCapacity(of constant.OneFile) (ofs []constant.OneFile) {
	c := of.Capacity
	for i := 0; i < c; i++ {
		nof := constant.OneFile{
			Channel:  of.Channel,
			Id:       of.Id + i,
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
	uri := strings.Join([]string{"https://t.me", of.Channel, strconv.Itoa(of.Id)}, "/")
	fmt.Printf("用户的下载文件夹目录: %s\n", constant.GetMainFolder())
	fmt.Printf("要下载的链接: %s\n", uri)
	_, err := util.GetLevelDB().Get([]byte(uri), nil)
	if errors.Is(err, leveldb.ErrNotFound) {
		log.Println("文件未下载过")
		util.GetLevelDB().Put([]byte(uri), []byte("downloaded"), nil)
	} else {
		log.Println("文件下载过,跳过")
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
		key := strconv.Itoa(of.Id)
		util.RenameByKey(key, of.FileName)
	}
	return of
}
func ParseLines(lines []string, f *os.File) (ofs []constant.OneFile) {
	for _, line := range lines {
		if of, err := parseOneLine(line); err != nil { // 如果解析失败 则写入文件
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
	if prefix, capacity := getCapacity(line); capacity != 0 {
		fmt.Printf("获取容量之后:%+v\n%v\n", prefix, capacity)
		of.SetCapacity(capacity)
	}

	if channel, id, err := getChannelAndFileID(line); err != nil {
		return nil, fmt.Errorf("URL: %s 不符合格式\n", line)
	} else {
		of.SetId(id)
		of.SetChannel(channel)
	}
	if prefix, offset := getOffset(line); offset != 0 {
		fmt.Printf("获取偏移量之后:prefix:%+v\nsuffix:%v\n", prefix, offset)
		fmt.Printf("此时的of:%+v\n", of)
		of.SetOffset(offset + of.Id)
		fmt.Printf("此时的of:%+v\n", of)
	}
	if tag, subtag, filename, _, _, err := getParam(line); err != nil {
		return nil, fmt.Errorf("URL: %s 不符合格式\n", line)
	} else {
		of.SetTag(tag)
		of.SetSubtag(subtag)
		of.SetFileName(filename)
	}
	log.Printf("解析结果:%+v\n", of)
	return of, nil
}
func getChannelAndFileID(url string) (channel string, file int, err error) {
	// 定义正则表达式
	pattern := `https?://t\.me/([a-zA-Z0-9]+)/([0-9]+)([@&#+%]?)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(url)
	if len(matches) == 4 {
		secondSegment := matches[1] // 第二段
		thirdSegment := matches[2]  // 第三段

		//fmt.Printf("URL: %s\n", url)
		//fmt.Printf("第二段: %s\n", secondSegment)
		//fmt.Printf("第三段: %s\n", thirdSegment)

		if thirdSegment_int, e := strconv.Atoi(thirdSegment); e != nil {
			return "", 0, fmt.Errorf("URL: %s 不符合格式\n", url)
		} else {
			return secondSegment, thirdSegment_int, nil
		}
	} else {
		return "", 0, fmt.Errorf("URL: %s 不符合格式\n", url)
	}
}
func getParam(input string) (tag, subtag, filename string, offset, capacity int, err error) {
	// 定义正则表达式来匹配属性
	re := regexp.MustCompile(`(#([^&@+%]*))|(&([^@+%]*))|(@([^+%]*))|(\+([^%]*))|(%(.*?))`)
	matches := re.FindAllStringSubmatch(input, -1)
	// 创建一个 Attributes 结构体实例
	// 遍历匹配结果并填充结构体
	for _, match := range matches {
		if match[2] != "" {
			tag = match[2]
		} else if match[4] != "" {
			subtag = match[4]
		} else if match[6] != "" {
			filename = match[6]
		} else if match[8] != "" {
			offset, err = strconv.Atoi(match[8])
			if err != nil {
				return "", "", "", 0, 0, err
			}
		} else if match[10] != "" {
			capacity, err = strconv.Atoi(match[10])
			if err != nil {
				return "", "", "", 0, 0, err
			}
		}
	}
	//fmt.Printf("tag = %v\nsubtag = %v\nfilename = %v\noffset = %v\ncapacity = %d\n", tag, subtag, filename, offset, capacity)
	return tag, subtag, filename, 0, 0, nil
}

/*
https://t.me/TNTsex/27584#其他文字
https://t.me/TNTsex/27584&其他文字
https://t.me/TNTsex/27584@其他文字
https://t.me/TNTsex/27584+其他文字
https://t.me/TNTsex/27584%其他文字
go实现分割网址部分和以第一个出现的特殊符号为分割的其他内容
perfix变量保存url部分 如https://t.me/TNTsex/27584
suffix保存包含用来分隔的这个特殊符号的其他文字 如 %其他文字
*/
func splitUrlAndParams(input string) (string, string) {
	// 定义特殊符号
	specialChars := []string{"#", "&", "@", "+", "%"}
	var firstSpecialChar string
	var index int

	// 找到第一个出现的特殊符号
	for _, char := range specialChars {
		if i := strings.Index(input, char); i != -1 {
			if firstSpecialChar == "" || i < index {
				firstSpecialChar = char
				index = i
			}
		}
	}
	// 如果找到了特殊符号，进行分割
	if firstSpecialChar != "" {
		prefix := input[:index]
		suffix := input[index:] // 包含特殊符号
		return prefix, suffix
	}
	// 如果没有找到特殊符号，返回原始字符串和空字符串
	return input, ""
}

/*
因为 %或+后面不可能再出现其他参数了，这两个属性也不能同时存在，所以单独处理
*/
func getOffset(s string) (line string, offset int) {
	if strings.Contains(s, "+") {
		// 偏移量
		prefix := strings.Split(s, "+")[0]
		suffix := strings.Split(s, "+")[1]
		offset, _ = strconv.Atoi(suffix)
		return prefix, offset
	}
	return s, 0
}
func getCapacity(s string) (line string, capacity int) {
	if strings.Contains(s, "%") {
		prefix := strings.Split(s, "%")[0]
		suffix := strings.Split(s, "%")[1]
		capacity, _ = strconv.Atoi(suffix)
		return prefix, capacity
	}
	return s, 0
}
