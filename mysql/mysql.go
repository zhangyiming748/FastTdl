package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var (
	useMysql bool
	engine   *xorm.Engine
)

// CREATE DATABASE `tdl` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';
func SetMysql() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:163453@tcp(192.168.1.9:3306)/tdl?charset=utf8")
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
