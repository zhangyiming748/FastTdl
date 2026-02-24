package util

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/zhangyiming748/sendEmailAlert"
)

func TimeoutAlert(msg string) {
	log.Println(msg)
	recipients := []string{
		//"zhangyiming748@protonmail.com",
		//"zhangyiming748@gmail.com",
		//"578779391@qq.com",
		"zhangyiming748@outlook.com",
	}
	hostname, _ := os.Hostname()
	email := sendEmailAlert.Info{
		Form:     "2352103020@qq.com",
		To:       recipients,
		Subject:  "单个文件运行超时警告",
		Text:     msg,
		Image:    "",
		Host:     sendEmailAlert.QQ.SMTP,
		Port:     sendEmailAlert.QQ.SMTPProt,
		Username: "2352103020@qq.com",
		Password: "hntplpwyeclmdiji",
	}
	email.AppendText(fmt.Sprintf("计算机名:%v", hostname))
	email.AppendText(fmt.Sprintf("架构:%v", runtime.GOARCH))
	email.AppendText(fmt.Sprintf("操作系统:%v", runtime.GOOS))
	email.AppendText(fmt.Sprintf("核心数:%v", runtime.NumCPU()))
	status := email.Send()
	log.Printf("发送邮件返回的状态:%v\n", status)
}
