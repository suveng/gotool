package util

import "testing"

func TestSendEmail(t *testing.T) {
	SendEmail("suveng@163.com", "标题", "内容", "suveng@qq.com", "", "smtp.qq.com", "587", ConvertHTML)
	SendEmail("suveng@163.com", "标题", "内容", "suveng@qq.com", "", "smtp.qq.com", "587", ConvertText)
}
