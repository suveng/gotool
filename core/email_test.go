package util

import "testing"

func TestSendEmail(t *testing.T) {
	// qq 测试
	SendEmail("suveng@163.com", "标题", "<h1>html</h1>", "suveng@qq.com", "", "smtp.qq.com", "587", ConvertHTML)
	SendEmail("suveng@163.com", "标题", "内容", "suveng@qq.com", "", "smtp.qq.com", "587", ConvertText)

	// 网易测试
	SendEmail("suveng@qq.com", "申通快递面试邀请", "HR面试邀请 请于20210801到上海青浦区申通大楼面试", "suveng@163.com", "", "smtp.163.com", "25", ConvertText)

	// 网易不能发 HTML 内容, 需要验证还是什么, 自行根据提示的url 去配置
	SendEmail("suveng@qq.com", "申通快递面试邀请", "<h1>HR面试邀请 请于20210801到上海青浦区申通大楼面试</h1> 欢迎来面试!", "suveng@163.com", "", "smtp.163.com", "25", ConvertHTML)
}
