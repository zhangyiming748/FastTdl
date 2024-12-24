package mysql

import (
	"log"
	"os"
	"strings"

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
	var host string
	if env := os.Getenv("MYSQL"); env == "" {
		host = "192.168.1.9:8889"
	}
	session := strings.Join([]string{"root:163453@tcp(", host, ")/tdl?charset=utf8&parseTime=true&loc=Asia%2FShanghai"}, "")
	engine, err = xorm.NewEngine("mysql", session)
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
