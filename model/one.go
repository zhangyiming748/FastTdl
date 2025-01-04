package model

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/zhangyiming748/FastTdl/mysql"
)

type File struct {
	Id        int64     `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	Origin    string    `xorm:"comment('原始url') VARCHAR(255)"`
	Channel   string    `xorm:"comment('频道id') VARCHAR(255)"`
	FileId    int       `xorm:"comment('文件id') Int"` // 文件id
	Tag       string    `xorm:"comment('主文件夹名 #后面的文件名') VARCHAR(255)"`
	Subtag    string    `xorm:"comment('子(二级)文件夹名 &后面的文件名') VARCHAR(255)"`
	Filename  string    `xorm:"comment('手动设置的文件名 @后面的文件名') VARCHAR(255)"`
	Offset    int       `xorm:"comment('偏移量 如下载当前媒体之后第n个文件 +后面的数字') Int"`
	Capacity  int       `xorm:"comment('下载当前文件和之后的n个文件 %后面的数字') Int"`
	From      string    `xorm:"comment('用来下载的设备') VARCHAR(255)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

/*
下载成功后插入
*/
func init() {
	if mysql.UseMysql() {
		mysql.GetMysql().Sync(File{})
	}
}

func (f *File) InsertOne() (int64, error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	device := fmt.Sprintf("%v with %v:%v", hostname, runtime.GOOS, runtime.GOARCH)
	f.From = device
	return mysql.GetMysql().InsertOne(f)
}

/*
根据分割后的原始url判断是否下载过
*/
func (f *File) FindByOriginURL() (bool, error) {
	return mysql.GetMysql().Where("channel = ? AND file_id = ?", f.Channel, f.FileId).Get(&f)
}

/*
根据自定义文件名判断是否下载过
*/
func (f *File) FindByFilename() (bool, error) {
	return mysql.GetMysql().Where("filename = ?", f.Filename).Get(&f)
}
