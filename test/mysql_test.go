package test

import (
	"fmt"
	"testing"

	"github.com/zhangyiming748/FastTdl/model"
	"github.com/zhangyiming748/FastTdl/mysql"
)

func TestConnectAndInit(t *testing.T) {
	mysql.SetMysql()

	err1 := mysql.GetMysql().Sync2(model.File{})

	err2 := mysql.GetMysql().Sync2(model.Channel{})
	fmt.Printf("err1:%v,err2:%v\n", err1, err2)
}
func TestCreateTable(t *testing.T) {
	mysql.SetMysql()
	mysql.GetMysql().Sync(model.Channel{})
}
func TestInsertChannel(t *testing.T) {
	mysql.SetMysql()
	err := mysql.GetMysql().Sync2(model.Channel{})
	fmt.Printf("err1:%v\n", err)
	c := new(model.Channel)
	c.ChannelId = "23826513401"
	c.ChannelUrl = "https://t.me/wsj008"
	c.ChannelName = "舞世界【热舞】"
	c.SetChannel()

	fmt.Printf("id :=%v\n", c.Id)
}
