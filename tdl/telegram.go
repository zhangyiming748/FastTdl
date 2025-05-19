package tdl

import (
	"errors"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zhangyiming748/FastTdl/constant"
	"github.com/zhangyiming748/FastTdl/model"
	"github.com/zhangyiming748/FastTdl/mysql"
	"github.com/zhangyiming748/FastTdl/util"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// zh2enMap 存储中文到英文的映射关系
var zh2enMap map[string]string

const (
	MaxRetries = 1 // 最大重试次数，当下载失败时重试的次数
)

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
func DownloadWithFolder(of constant.OneFile, proxy string, f *os.File) constant.OneFile {
	// 构建下载链接
	uri := strings.Join([]string{"https://t.me", of.Channel, strconv.Itoa(of.FileId)}, "/")
	p := constant.GetParams()

	// 输出下载信息
	fmt.Printf("用户的下载文件夹目录: %s\n", p.GetMainFolder())
	fmt.Printf("要下载的链接: %s\t%+v\n", uri, of)

	// 检查是否已下载过（MySQL模式）
	if mysql.UseMysql() {
		oneline := new(model.File)
		oneline.Channel = of.Channel
		oneline.FileId = of.FileId
		oneline.Filename = of.FileName
		if oneline.Filename != "" {
			if found, _ := oneline.FindByFilename(); found {
				log.Println("相同文件名的文件下载过,跳过")
				return of
			}
		}
		if found, _ := oneline.FindByOriginURL(); found {
			log.Println("相同url的文件下载过,跳过")
			return of
		}
		log.Println("数据库中没有查到相同文件,继续下载")
	} else {
		_, err := util.GetLevelDB().Get([]byte(uri), nil)
		if errors.Is(err, leveldb.ErrNotFound) {
			log.Println("文件未下载过")
			util.GetLevelDB().Put([]byte(uri), []byte("downloaded"), nil)
		} else {
			log.Println("文件下载过,跳过")
			of.SetStatus()
			return of
		}
	}
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
	// 执行下载，支持重试机制
	var downloadErr error
	for i := 0; i < MaxRetries; i++ {
		if i > 0 {
			log.Printf("第%d次重试下载\n", i+1)
			time.Sleep(time.Second * 3) // 重试前等待3秒，避免频繁请求
		}

		// 执行下载命令
		if err := util.ExecTdlCommand(proxy, uri, target); err == nil {
			log.Printf("下载成功")
			downloadErr = nil
			break
		} else {
			downloadErr = err
			log.Printf("第%d次下载失败: %v\n", i+1, err)
		}
	}

	// 处理下载失败的情况
	if downloadErr != nil {
		log.Printf("达到最大重试次数%d,下载命令执行出错:%+v\n", MaxRetries, of)
		f.WriteString(fmt.Sprintf("%v\n", origin))
		return of
	}

	// 下载成功后的数据库记录
	if mysql.UseMysql() {
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
			log.Printf("写入数据库失败")
		} else {
			log.Printf("写入数据库成功")
		}
	}

	// 更新任务状态
	of.SetStatus()
	// 如果指定了文件名，则重命名文件
	if of.FileName != "" {
		util.RenameByKey(of)
	}

	return of
}

// ParseLines 解析多行下载链接
// lines: 待解析的链接列表
// f: 解析失败记录文件
// returns: 解析成功的下载任务列表
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
		src = strings.Replace(src, k, v, -1)
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
