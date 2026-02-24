package util

import "testing"

func TestSendTimeoutEmail(t *testing.T) {
	TimeoutAlert("测试邮件")
}
func TestSendEmail(t *testing.T) {
	Alarm("正文第一行", "正文第二行", "正文第三行", "正文第四行")
}
