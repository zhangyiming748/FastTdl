package test

import (
	"testing"
	"github.com/zhangyiming748/FastTdl/util"
)

func TestFingNewestFile(t *testing.T) {
	ret:=util.FindLatestFileGo("/Users/zen/github/FastTdl")
	t.Logf("最终结果:%+v",ret)
}
