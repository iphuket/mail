package mail

import (
	"net/smtp"
	"strings"
)

// Send is Send Mail

// SMTP 配置
type SMTP struct {
	Identity string
	FromMail string
	PassWord string
	MailHost string
	MailPort string
}

// Send Mail toMail 接收邮箱, FromMail 发送者邮箱,PassWord：smtp 密码,MailHost smtp host,MailPort smtp port,nikename 发送者昵称,UserName 发送者用户名，SubJect 主题/标题,Body 发送内容(支持html)
func Send(sm *SMTP, toMail, NickName, UserName, SubJect, Body string) error {
	auth := smtp.PlainAuth(sm.Identity, sm.FromMail, sm.PassWord, sm.MailHost)
	to := []string{toMail}
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + NickName + "<" + UserName + ">\r\nSubject: " + SubJect + "\r\n" + contentType + "\r\n\r\n" + Body)
	err := smtp.SendMail(sm.MailHost+":"+sm.MailPort, auth, sm.FromMail, to, msg)
	if err != nil {
		return err
	}
	return nil
}
