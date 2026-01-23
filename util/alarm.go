package util

import (
	"log"
	"strings"

	"github.com/zhangyiming748/sendEmailAlert"
)

func Alarm(words ...string) {
	info := new(sendEmailAlert.Info)
	info.SetFrom("2352103020@qq.com") //${{ secrets.FROM }}
	tos := []string{
		"578779391@qq.com",
		"zhangyiming748@gmail.com",
		"zhangyiming748@outlook.com",
		"18904892728@163.com",
		"18904892728@189.cn",
		"zhangyiming748@linux.do",
		"zhangyiming748@icloud.com",
	}
	info.SetTo(tos)
	info.SetSubject("程序运行结束通知")
	info.SetText("FastTdl主程序运行结束")
	info.SetHost(sendEmailAlert.QQ.SMTP)
	info.SetPort(sendEmailAlert.QQ.SMTPProt)
	info.SetUsername("2352103020@qq.com") //${{ secrets.FROM }}
	info.SetPassword("ocuplrlgwgelebej")  //${{ secrets.PASSWORD }}
	info.AppendText(strings.Join(words, "<br>"))
	status := info.Send()
	log.Println(status)
}
