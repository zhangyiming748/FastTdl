package test

import (
	"github.com/zhangyiming748/FastTdl/model"
	"github.com/zhangyiming748/FastTdl/sqlite"
	"testing"
)

func TestHistoryModel(t *testing.T) {
	sqlite.SetSqlite()

	// 同步表结构
	s := new(model.Save)

	// 插入测试数据
	s.SaveSize = 1.0
	err := s.Insert()
	if err != nil {
		t.Errorf("插入数据失败: %v", err)
	}
}
