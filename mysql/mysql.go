package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangyiming748/FastTdl/constant"
	"xorm.io/xorm"
)

var (
	useMysql bool
	engine   *xorm.Engine
)

func SetMysql() {

	var err error
	p:=constant.GetParams()
	user := p.GetUser()
	password :=p.GetPassword()
	host := p.GetHost()
	port :=p.GetPort()
	
	// 先连接到 MySQL 服务器（不指定数据库）
	rootDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", user, password, host, port)
	tempEngine, err := xorm.NewEngine("mysql", rootDSN)
	if err != nil {
		log.Printf("连接MySQL服务器失败: %v\n", err)
		useMysql = false
		return
	}

	// 检查数据库是否存在
	rows, err := tempEngine.QueryString("SELECT SCHEMA_NAME FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = 'tdl'")
	if err != nil {
		log.Printf("查询数据库失败: %v\n", err)
		useMysql = false
		return
	}

	// 如果数据库不存在，创建它
	if len(rows) == 0 {
		_, err = tempEngine.Exec("CREATE DATABASE `tdl` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci'")
		if err != nil {
			log.Printf("创建数据库失败: %v\n", err)
			useMysql = false
			return
		}
		log.Println("成功创建数据库 tdl")
	}

	// 关闭临时连接
	tempEngine.Close()

	// 连接到 tdl 数据库
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/tdl?charset=utf8mb4", user, password, host, port)
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Printf("连接tdl数据库失败: %v\n", err)
		useMysql = false
		return
	}

	if err = engine.Ping(); err != nil {
		log.Printf("连接数据库失败: %v\n", err)
		useMysql = false
		return
	}else{
		log.Printf("成功Ping到数据库\n")
		useMysql = true
	}

	log.Printf("成功连接到数据库\n")
}

func GetMysql() *xorm.Engine {
	return engine
}

func UseMysql() bool {
	return useMysql
}
