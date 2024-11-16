package util

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

var levelDB *leveldb.DB

func SetLevelDB() {
	location := "leveldb"
	db, err := leveldb.OpenFile(location, nil)
	if err != nil {
		log.Fatalf("leveldb数据库创建失败:%v\n", err)
	}
	levelDB = db
}
func GetLevelDB() *leveldb.DB {
	return levelDB
}
