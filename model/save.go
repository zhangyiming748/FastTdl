package model

import (
	"errors"
	"FastTdl/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

type Save struct {
	Id       int64          `gorm:"primaryKey;autoIncrement;comment:主键"`
	FileName string         `gorm:"comment:文件名"`
	Before   string         `gorm:"comment:转换前大小(MB)"`
	After    string         `gorm:"comment:转换后大小(MB)"`
	SaveSize float64        `gorm:"comment:节省(为正数的时候)的空间(MB)"`
	CreateAt time.Time      `gorm:"autoCreateTime;comment:创建时间"`
	UpdateAt time.Time      `gorm:"autoUpdateTime;comment:更新时间"`
	DeleteAt gorm.DeletedAt `gorm:"index;comment:删除时间"`
}

func (s *Save) Sync() {
	log.Printf("开始同步表结构\n")
	if err := sqlite.GetSqlite().AutoMigrate(&Save{}); err != nil {
		log.Fatalf("同步表结构History失败:%s", err.Error())
	}
	log.Printf("同步表结构完成\n")
}

func (s *Save) Insert() error {
	db := sqlite.GetSqlite()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	result := db.Create(&s)
	return result.Error
}

func (s *Save) Update() error {
	db := sqlite.GetSqlite()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	result := db.Save(&s)
	return result.Error
}

func (s *Save) Delete() error {
	db := sqlite.GetSqlite()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	result := db.Delete(&s)
	return result.Error
}

func (s *Save) GetById(id int64) error {
	db := sqlite.GetSqlite()
	if db == nil {
		return errors.New("数据库连接未初始化")
	}
	result := db.First(&s, id)
	return result.Error
}

func (s *Save) GetAll() ([]Save, error) {
	db := sqlite.GetSqlite()
	if db == nil {
		return nil, errors.New("数据库连接未初始化")
	}
	var Saves []Save
	result := db.Find(&Saves)
	return Saves, result.Error
}
