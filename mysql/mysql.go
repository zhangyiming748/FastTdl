package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	useMysql bool
	engine   *xorm.Engine
)

// CREATE DATABASE `tdl` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';
func SetMysql() {
	var err error
	user := "root"
	password := "163453"
	host := "192.168.1.9"
	port := 3306
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/tdl?charset=utf8", user, password, host, port)

	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		useMysql = false
	} else {
		useMysql = true
		log.Printf("连接数据库成功:%v\n", engine)
	}
}

func GetMysql() *xorm.Engine {
	return engine
}

func UseMysql() bool {
	return useMysql
}
