package main

import (
	"fmt"

	"./mail"
)

// Mail smtp
// var Mail = &mail.SMTP{FromMail: "noreply@oovmi.com", PassWord: "kvzuqqubgicgbicj", MailHost: "smtp.qq.com", MailPort: "25"}
var Mail = &mail.SMTP{FromMail: "noreply@oovmi.com", PassWord: "kvzuqqubgicgbicj", MailHost: "smtp.qq.com", MailPort: "25"}

func main() {
	mail.NewServer()
	send()
}

func send() {
	err := mail.Send(Mail, "494745409@qq.com", "邮件测试昵称", "moko@oossi.com", "邮件测试主题", "<h1>邮件内容</h1>")
	if err != nil {
		fmt.Println(err)
	}

}
