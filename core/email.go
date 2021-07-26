package util

import (
	"fmt"
	"log"
	"net/smtp"
)

// 这里使用 qq 的 smtp 发送邮件, 开启smtp服务 https://service.mail.qq.com/cgi-bin/help?subtype=1&&id=28&&no=331
// 1. 准备qq账号 2. 开启smtp 发送邮箱服务, 获取授权码 3. 将授权码放在password作为授权, 可以重复获取, 一旦泄露可以改qq密码应急

const (
	SMTPHost     = "smtp.qq.com"
	SMTPPort     = "587"
	SMTPUsername = "suveng@qq.com"
	SMTPPassword = ""
)

func SendEmail(to string, subject string, content string, from string, password string, hostname string, port string) {
	auth := smtpAuthorize(from, password, hostname)

	smtpURL := createSMTPURL(hostname, port)

	// 格式化消息体
	format := "Subject: %s\r\n\r\n" + "%s\r\n"
	emailContent := fmt.Sprintf(format, subject, content)

	// 发送
	send(to, smtpURL, auth, from, emailContent)
}

func send(to string, smtpURL string, auth smtp.Auth, from string, emailContent string) {
	err := smtp.SendMail(smtpURL, auth, from, []string{to}, []byte(emailContent))
	if err != nil {
		panic(err)
	}
}

// 创建 smtp url
func createSMTPURL(hostname string, port string) string {
	smtpURL := hostname + ":" + port
	return smtpURL
}

// smtp 服务 授权
func smtpAuthorize(from string, passwrod string, hostname string) smtp.Auth {
	// todo:// 这里identify是干什么的, 不用传
	auth := smtp.PlainAuth("", from, passwrod, hostname)
	return auth
}

func sendHTMLEmail(receiver string, html []byte) {
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	msg := append([]byte("Subject: 这里是标题内容\r\n"+
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"),
		html...)

	err := smtp.SendMail(SMTPHost+SMTPPort, auth, SMTPUsername, []string{receiver}, msg)
	if err != nil {
		log.Fatal("failed to send email:", err)
	}
}

func main() {

}
