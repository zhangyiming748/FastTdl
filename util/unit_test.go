package util

import (
	"log"
	"testing"
)

// func TestFind(t *testing.T) {
// 	SetLevelDB()
// 	get, err := GetLevelDB().Get([]byte("key"), nil)
// 	if errors.Is(err, leveldb.ErrNotFound) {
// 		log.Printf("get = %v, err = %v\n", get, err)
// 	}
// 	log.Printf("get = %v, err = %v\n", get, err)
// 	err = GetLevelDB().Put([]byte("key"), []byte("value"), nil)
// 	if err != nil {
// 		return
// 	}
// }

func TestCount(t *testing.T) {
	count, err := CountFiles("/Users/zen/github")
	if err != nil {
		log.Printf("统计文件数量出错：%v\n", err)
		return
	}
	t.Logf("目录下共有 %d 个文件\n", count)
}
