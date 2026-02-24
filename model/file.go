package model

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"FastTdl/sqlite"
	"gorm.io/gorm"
)

type File struct {
	Id        int64          `gorm:"primaryKey;autoIncrement;comment:主键id"`
	Origin    string         `gorm:"comment:原始url"`
	Channel   string         `gorm:"comment:频道id"`
	FileId    int            `gorm:"comment:文件id"` // 文件id
	Tag       string         `gorm:"comment:主文件夹名 #后面的文件名"`
	Subtag    string         `gorm:"comment:子(二级)文件夹名 &后面的文件名"`
	Filename  string         `gorm:"comment:手动设置的文件名 @后面的文件名"`
	Offset    int            `gorm:"comment:偏移量 如下载当前媒体之后第n个文件 +后面的数字"`
	Capacity  int            `gorm:"comment:下载当前文件和之后的n个文件 %后面的数字"`
	From      string         `gorm:"comment:用来下载的设备"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

/*
下载成功后插入
*/

func (f *File) InsertOne() (int64, error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	device := fmt.Sprintf("%v with %v:%v", hostname, runtime.GOOS, runtime.GOARCH)
	f.From = device

	db := sqlite.GetSqlite()
	result := db.Create(&f)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, result.Error
}

/*
根据分割后的原始url判断是否下载过
*/
func (f *File) FindByOriginURL() (bool, error) {
	db := sqlite.GetSqlite()
	result := db.Where("channel = ? AND file_id = ?", f.Channel, f.FileId).First(&f)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

/*
根据自定义文件名判断是否下载过
*/
func (f *File) FindByFilename() (bool, error) {
	db := sqlite.GetSqlite()
	result := db.Where("filename = ?", f.Filename).First(&f)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

/*
根据fileid判断是否下载过
*/
func (f *File) FindByFileId() (bool, error) {
	db := sqlite.GetSqlite()
	result := db.Where("file_id = ?", f.FileId).First(&f)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

/*
根据fileid和channel判断是否下载过，如果有同一条记录同时这两个值相等，则认为是同一个文件
*/
func (f *File) FindByFileIdAndChannel() (bool, error) {
	db := sqlite.GetSqlite()
	result := db.Where("file_id = ? AND channel = ?", f.FileId, f.Channel).First(&f)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}
