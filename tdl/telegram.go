package tdl

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"FastTdl/constant"
	"FastTdl/model"
	"FastTdl/util"
)

// zh2enMap 存储中文到英文的映射关系
var zh2enMap map[string]string

// init 初始化函数，程序启动时加载中英文映射表
func init() {
	zh2enMap = zh2en("zh_cn2en_us.md")
}

// GenerateDownloadLinkByCapacity 根据容量生成多个下载任务
// of: 原始下载任务
// returns: 生成的下载任务列表
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

// DownloadWithFolder 执行文件下载任务
// of: 下载任务信息
// proxy: 代理服务器地址
// f: 失败记录文件
// returns: 更新后的下载任务信息
func DownloadWithFolder(of constant.OneFile, p constant.Parameter) constant.OneFile {
	// 构建下载链接
	uri := strings.Join([]string{"https://t.me", of.Channel, strconv.Itoa(of.FileId)}, "/")

	// 输出下载信息
	fmt.Printf("用户的下载文件夹目录: %s\n", p.GetMainFolder())
	fmt.Printf("要下载的链接: %s\t%+v\n", uri, of)

	// 检查是否已下载过（Sqlite模式）

	oneline := new(model.File)
	oneline.Channel = of.Channel
	oneline.FileId = of.FileId
	oneline.Filename = of.FileName
	if found, _ := oneline.FindByOriginURL(); found {
		log.Println("相同url的文件下载过,跳过")
		return of
	}
	if found, _ := oneline.FindByFileIdAndChannel(); found {
		log.Println("相同文件下载过,跳过")
		return of
	}
	log.Println("数据库中没有查到相同文件,继续下载")

	// 构建目标文件夹路径
	target := p.GetMainFolder()
	if tag := of.Tag; tag != "" {
		target = filepath.Join(target, tag)
		if subtag := of.Subtag; subtag != "" {
			target = filepath.Join(target, subtag)
		}
	}
	os.MkdirAll(target, 0755)

	// 构建完整的原始链接（包含所有参数）
	origin := uri
	if of.Tag != "" {
		origin = strings.Join([]string{origin, of.Tag}, "#")
	}
	if of.Subtag != "" {
		origin = strings.Join([]string{origin, of.Subtag}, "&")
	}
	if of.FileName != "" {
		origin = strings.Join([]string{origin, of.FileName}, "@")
	}
	if of.Offset != 0 {
		origin = strings.Join([]string{origin, strconv.Itoa(of.Offset)}, "+")
	}
	if of.Capacity != 0 {
		origin = strings.Join([]string{origin, strconv.Itoa(of.Capacity)}, "%")
	}
	if err := util.ExecTdlCommand(p.GetProxy(), uri, target); err != nil {
		log.Printf("下载失败: %v\n", err)
	} else {
		log.Printf("下载成功: %v\n", uri)
		// 下载成功后的数据库记录

		oneline := new(model.File)
		oneline.Origin = origin
		oneline.Channel = of.Channel
		oneline.FileId = of.FileId
		oneline.Tag = of.Tag
		oneline.Subtag = of.Subtag
		oneline.Filename = of.FileName
		oneline.Offset = of.Offset
		oneline.Capacity = of.Capacity
		log.Printf("成功后写入数据库")
		_, err := oneline.InsertOne()
		if err != nil {
			log.Println("写入数据库失败")
		} else {
			log.Println("写入数据库成功")
		}

	}
	// 更新任务状态
	of.SetStatus()
	// 如果指定了文件名，则重命名文件
	if of.FileName != "" {
		util.RenameByKey(of, p)
	}

	return of
}

// ParseLines 解析多行下载链接
// lines: 待解析的链接列表
// f: 解析失败记录文件
// returns: 解析成功的下载任务列表
func ParseLines(lines []string) (ofs []constant.OneFile) {
	for _, line := range lines {
		if of, err := parseOneLine(line); err != nil { // 如果解析失败 则写入文件
			log.Printf("解析失败:%s\n", line)
			continue
		} else {
			ofs = append(ofs, *of)
		}
	}
	return ofs
}

// parseOneLine 解析单行下载链接
// line: 待解析的链接
// returns: 解析后的下载任务信息，错误信息
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
	if strings.Contains(line, "/c/") {
		originUrl = strings.Join([]string{"https://t.me/c/", of.Channel, strconv.Itoa(of.FileId)}, "/")
	}
	params := strings.Replace(line, originUrl, "", 1)
	tag, subtag, filename, offset, capacity, err := getParam(params)
	if strings.Contains(of.Channel, "luowo007") {
		tag = "dance"
	} else if strings.Contains(of.Channel, "rewu8666") {
		tag = "dance"
	} else if strings.Contains(of.Channel, "swxiu") {
		tag = "dance"
	}
	if tag == "蒂法" || strings.ToUpper(tag) == "TIFA" {
		tag = "最终幻想"
		subtag = "蒂法"
	}
	if strings.ToUpper(tag) == "2B" {
		tag = "NieR"
		subtag = "2B"
	}
	if tag == "爱丽丝" || strings.ToUpper(tag) == "AERITH" {
		tag = "最终幻想"
		subtag = "爱丽丝"
	}
	if err != nil {
		return nil, err
	} else {
		of.SetTag(replace(tag))
		of.SetSubtag(replace(subtag))
		of.SetFileName(filename)
		of.SetOffset(offset)
		of.SetCapacity(capacity)
	}
	log.Printf("解析结果:%+v\n", of)
	return of, nil
}

// getChannelAndFileID 从URL中提取频道名和文件ID
// url: 完整的下载链接
// returns: 频道名，文件ID，错误信息
func getChannelAndFileID(url string) (channel string, file int, err error) {
	//https://t.me/guoman_08/2148#&@+%
	static := "https://t.me/"
	if strings.Contains(url, "/c/") {
		static = "https://t.me/c/"
	}
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

// getParam 解析URL中的附加参数
// input: URL中的参数部分
// returns: 标签，子标签，文件名，偏移量，容量，错误信息
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
		// 替换常见的中文标点符号为英文标点符号
		filename = strings.Replace(filename, "，", ",", -1)  // 中文逗号
		filename = strings.Replace(filename, "。", ".", -1)  // 中文句号
		filename = strings.Replace(filename, "；", ";", -1)  // 中文分号
		filename = strings.Replace(filename, "：", ":", -1)  // 中文冒号
		filename = strings.Replace(filename, "？", "?", -1)  // 中文问号
		filename = strings.Replace(filename, "！", "!", -1)  // 中文感叹号
		filename = strings.Replace(filename, "“", "\"", -1) // 中文左双引号
		filename = strings.Replace(filename, "”", "\"", -1) // 中文右双引号
		filename = strings.Replace(filename, "‘", "'", -1)  // 中文左单引号
		filename = strings.Replace(filename, "’", "'", -1)  // 中文右单引号
		filename = strings.Replace(filename, "（", "(", -1)  // 中文左括号
		filename = strings.Replace(filename, "）", ")", -1)  // 中文右括号
		filename = strings.Replace(filename, "【", "[", -1)  // 中文左方括号
		filename = strings.Replace(filename, "】", "]", -1)  // 中文右方括号
		filename = strings.Replace(filename, "《", "<", -1)  // 中文左书名号
		filename = strings.Replace(filename, "》", ">", -1)  // 中文右书名号

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

// replace 将字符串中的中文替换为对应的英文
// src: 源字符串
// returns: 替换后的字符串
func replace(src string) string {
	for k, v := range zh2enMap {
		// 使用strings.EqualFold进行不区分大小写的比较和替换
		if strings.EqualFold(src, k) {
			src = v
		}
	}
	return src
}

// zh2en 从文件中加载中英文映射关系
// fp: 映射文件路径
// returns: 中英文映射表
func zh2en(fp string) map[string]string {
	result := make(map[string]string)
	seen := make(map[string]bool) // 用于记录已经处理过的key
	content, err := os.ReadFile(fp)
	if err != nil {
		log.Printf("读取文件失败: %v\n", err)
		return result
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") || !strings.Contains(line, "|") || strings.Contains(line, ":---:") {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) != 4 {
			continue
		}

		original := strings.TrimSpace(parts[1])
		translations := strings.TrimSpace(parts[2])
		if original == "" || translations == "" {
			continue
		}

		for _, trans := range strings.Split(translations, ";") {
			trans = strings.TrimSpace(trans)
			if trans != "" && !seen[trans] { // 只处理未见过的key
				result[trans] = original
				seen[trans] = true // 标记该key已处理
			}
		}
	}
	return result
}
