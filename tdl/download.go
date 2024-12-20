package tdl

import (
	"fmt"
	"github.com/zhangyiming748/FastTdl/constant"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseLines(lines []string, f *os.File) (ofs []constant.OneFile) {
	for _, line := range lines {
		if of, err := parseOneLine(line); err != nil { // 如果解析失败 则写入文件
			f.WriteString(line)
			f.WriteString("\n")
			continue
		} else {
			ofs = append(ofs, of)
		}
	}
	return ofs
}
func parseOneLine(line string) (of constant.OneFile, err error) {
	if of.Channel, of.Id, err = getChannelAndFileID(line); err != nil {
		return constant.OneFile{}, fmt.Errorf("URL: %s 不符合格式\n", line)
	}
	if of.Tag, of.Subtag, of.FileName, of.Offset, of.Capacity, err = getParam(line); err != nil {
		return constant.OneFile{}, fmt.Errorf("URL: %s 不符合格式\n", line)
	}
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

		fmt.Printf("URL: %s\n", url)
		fmt.Printf("第二段: %s\n", secondSegment)

		fmt.Printf("第三段: %s\n", thirdSegment)
		if thirdSegment_int, e := strconv.Atoi(thirdSegment); e != nil {
			return "", 0, fmt.Errorf("URL: %s 不符合格式\n", url)
		} else {
			return secondSegment, thirdSegment_int, nil
		}
	} else {
		return "", 0, fmt.Errorf("URL: %s 不符合格式\n", url)
	}
}
func getParam(url string) (tag, subtag, filename string, offset, capacity int, err error) {
	// 提取属性值
	if strings.Contains(url, "#") {
		parts := strings.SplitN(url, "#", 2)
		if len(parts) > 1 {
			tag = strings.Split(parts[1], "&")[0]
		}
	}
	if strings.Contains(url, "&") {
		parts := strings.SplitN(url, "&", 2)
		if len(parts) > 1 {
			subtag = strings.Split(parts[1], "@")[0]
		}
	}
	if strings.Contains(url, "@") {
		parts := strings.SplitN(url, "@", 2)
		if len(parts) > 1 {
			filename = strings.Split(parts[1], "+")[0]
		}
	}
	if strings.Contains(url, "+") {
		parts := strings.SplitN(url, "+", 2)
		if len(parts) > 1 {
			if offset, err = strconv.Atoi(strings.Split(parts[1], "%")[0]); err != nil {
				return "", "", "", 0, 0, err
			}
			//offset = strings.Split(parts[1], "%")[0]

		}
	}
	if strings.Contains(url, "%") {
		parts := strings.SplitN(url, "%", 2)
		if len(parts) > 1 {
			if capacity, err = strconv.Atoi(parts[1]); err != nil {
				return "", "", "", 0, 0, err
			}
			//capacity = parts[1]
		}
	}
	fmt.Printf("主文件夹名: %s\n子文件夹名: %s\n文件名: %s\n偏移量: %d\n容量: %d\n", tag, subtag, filename, offset, capacity)
	return tag, subtag, filename, offset, capacity, nil
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
