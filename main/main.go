package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/iphuket/mail"
)

// Mail smtp
// var Mail = &mail.SMTP{FromMail: "noreply@oovmi.com", PassWord: "kvzuqqubgicgbicj", MailHost: "smtp.qq.com", MailPort: "25"}
var Mail = &mail.SMTP{FromMail: "no-reply@oovmi.com", PassWord: "kvzuqqubgicgbicj", MailHost: "smtp.minis.app", MailPort: "1025"}

func main() {
	go web()
	mail.NewServer()
}
func web() {
	app := gin.Default()
	app.GET("/send", send)
	app.RunTLS(":99", "./2006285_smtp.minis.app.pem", "./2006285_smtp.minis.app.key")
}
func send(c *gin.Context) {
	err := mail.Send(Mail, "494745409@qq.com", "邮件测试昵称", "moko@oossi.com", "邮件测试主题", "<h1>邮件内容</h1>")
	if err != nil {
		fmt.Println(err)
		c.Writer.Write([]byte("error " + fmt.Sprint(err)))
		return
	}
	c.Writer.Write([]byte("mail send success"))
}
