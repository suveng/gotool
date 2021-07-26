package util

import "testing"

func TestSendEmail(t *testing.T) {
	SendEmail("suveng@163.com", "标题", "内容", SMTPUsername, SMTPPassword, SMTPHost, SMTPPort)
}
