package model

import (
	"time"

	"FastTdl/sqlite"
	"gorm.io/gorm"
)

type Channel struct {
	Id          int64          `gorm:"primaryKey;autoIncrement;comment:主键id"`
	ChannelName string         `gorm:"comment:频道名"`
	ChannelUrl  string         `gorm:"comment:频道url字段"`
	ChannelId   string         `gorm:"comment:频道id"` // 自动去除 -100前缀 ,默认文件名第一个下划线之前
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (c *Channel) SetChannel() (int64, error) {
	db := sqlite.GetSqlite()
	result := db.Create(&c)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (c *Channel) GetByChannelId(channelId string) error {
	db := sqlite.GetSqlite()
	result := db.Where("channel_id = ?", channelId).First(&c)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return result.Error
	}
	return nil
}

func (c *Channel) GetByChannelName(channelName string) error {
	db := sqlite.GetSqlite()
	result := db.Where("channel_name = ?", channelName).First(&c)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return result.Error
	}
	return nil
}

func (c *Channel) GetByChannelURL(channelURL string) error {
	db := sqlite.GetSqlite()
	result := db.Where("channel_url = ?", channelURL).First(&c)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return result.Error
	}
	return nil
}
