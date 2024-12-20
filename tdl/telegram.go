package tdl

import (
	"fmt"
	"github.com/zhangyiming748/FastTdl/constant"
	"github.com/zhangyiming748/FastTdl/util"
	"log"
	"os"
	"path/filepath"
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
