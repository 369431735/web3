package utils

import (
	"net/smtp"
	"net/textproto"
	"pledge-backend/config"

	"github.com/jordan-wright/email"
)

// SendEmail 发送电子邮件
// 根据配置发送纯文本或HTML格式的电子邮件
// data: 邮件内容（字节数组）
// dataType: 内容类型（1=纯文本，2=HTML）
// 返回: 可能的错误信息
func SendEmail(data []byte, dataType int) error {
	// 创建电子邮件对象
	e := &email.Email{
		To:      config.Config.Email.To,      // 收件人列表
		Cc:      config.Config.Email.Cc,      // 抄送人列表
		From:    config.Config.Email.From,    // 发件人
		Subject: config.Config.Email.Subject, // 邮件主题
		Headers: textproto.MIMEHeader{},      // 邮件头
	}
	// 根据数据类型设置邮件内容
	if dataType == 1 {
		e.Text = data // 纯文本内容
	} else {
		e.HTML = data // HTML内容
	}
	// 使用SMTP发送邮件
	return e.Send(config.Config.Email.Host+":"+config.Config.Email.Port,
		smtp.PlainAuth("",
			config.Config.Email.Username,
			config.Config.Email.Pwd,
			config.Config.Email.Host))
}

// SendEmailWithAttach 发送带附件的电子邮件
// 根据配置发送带有附件的纯文本或HTML格式的电子邮件
// data: 邮件内容（字节数组）
// dataType: 内容类型（1=纯文本，2=HTML）
// filename: 附件文件路径
// 返回: 可能的错误信息
func SendEmailWithAttach(data []byte, dataType int, filename string) error {
	// 创建电子邮件对象
	e := &email.Email{
		To:      config.Config.Email.To,      // 收件人列表
		Cc:      config.Config.Email.Cc,      // 抄送人列表
		From:    config.Config.Email.From,    // 发件人
		Subject: config.Config.Email.Subject, // 邮件主题
		Headers: textproto.MIMEHeader{},      // 邮件头
	}
	// 根据数据类型设置邮件内容
	if dataType == 1 {
		e.Text = data // 纯文本内容
	} else {
		e.HTML = data // HTML内容
	}
	// 添加附件
	_, err := e.AttachFile(filename)
	if err != nil {
		return err
	}
	// 使用SMTP发送邮件
	return e.Send(config.Config.Email.Host+config.Config.Email.Port,
		smtp.PlainAuth("",
			config.Config.Email.Username,
			config.Config.Email.Pwd,
			config.Config.Email.Host))
}
