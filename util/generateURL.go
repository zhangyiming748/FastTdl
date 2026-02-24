package util

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GenerateURL(baseURL string, offset int, tag string) error {
	// 生成网址的具体逻辑
	baseURL = strings.Replace(baseURL, "?single", "", -1)
	var links []string
	lastIndex := strings.LastIndex(baseURL, "/")
	if lastIndex == -1 {
		return errors.New("不合法的网址")
	}
	prefix := baseURL[:lastIndex]
	suffix := baseURL[lastIndex+1:]
	log.Printf("prefix:%s, suffix:%s", prefix, suffix)
	start, err := strconv.Atoi(suffix)
	if err != nil {
		return errors.New("不合法的网址")
	}
	for i := 0; i < offset; i++ {
		index := start + i
		link := strings.Join([]string{prefix, strconv.Itoa(index)}, "/")
		if tag != "" {
			link = strings.Join([]string{link, tag}, "#")
		}
		links = append(links, link)
	}
	fmt.Println(links)
	WriteByLine("post.link", links)
	return nil
}
