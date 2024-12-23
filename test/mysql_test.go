package test

import (
	"github.com/zhangyiming748/FastTdl/model"
	"github.com/zhangyiming748/FastTdl/mysql"
	"testing"
)

func TestConnectAndInit(t *testing.T) {
	mysql.SetMysql()
	metas, err := mysql.GetMysql().DBMetas()
	if err != nil {
		return
	} else {
		t.Log(metas)
	}
	info, err := mysql.GetMysql().TableInfo(model.File{})
	if err != nil {
		return
	} else {
		t.Log(info)
	}
	err = mysql.GetMysql().Sync(model.File{})
	if err != nil {
		return
	} else {
		t.Log("success")
	}
}
