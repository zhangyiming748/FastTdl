package test

import (
	"github.com/zhangyiming748/FastTdl/model"
	"github.com/zhangyiming748/FastTdl/mysql"
	"log"
	"testing"
)

func TestConnectAndInit(t *testing.T) {
	mysql.SetMysql()

	err := mysql.GetMysql().Sync(model.File{})
	if err != nil {
		return
	} else {
		t.Log("success")
	}
	oneline := new(model.File)
	oneline.Channel = "of.Channel"
	oneline.FileId = 3
	oneline.Tag = "of.Tag"
	oneline.Subtag = "of.Subtag"
	oneline.Filename = "of.FileName"
	oneline.Offset = 1
	oneline.Capacity = 2
	log.Printf("成功后写入数据库")
	_, err = oneline.InsertOne()
	if err != nil {
		log.Printf("写入数据库失败")
	} else {
		log.Printf("写入数据库成功")
	}
}
