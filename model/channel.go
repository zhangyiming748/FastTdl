package model

import (
	"time"

	"github.com/zhangyiming748/FastTdl/mysql"
)

type Channel struct {
	Id          int64     `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	ChannelName string    `xorm:"comment('频道名') VARCHAR(255)"`
	ChannelUrl  string    `xorm:"comment('频道url字段') VARCHAR(255)"`
	ChannelId   string     `xorm:"comment('频道id') VARCHAR(255)"` // 自动去除 -100前缀 ,默认文件名第一个下划线之前
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
	DeletedAt   time.Time `xorm:"deleted"`
}

func (c *Channel) SetChannel() (int64,error) {
	return mysql.GetMysql().InsertOne(c)
}