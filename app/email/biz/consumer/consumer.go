package consumer

import "github.com/trashwbin/gomall-demo/app/email/biz/consumer/email"

// Init 函数用于初始化邮件消费模块
// 该函数没有输入参数和返回值
// 调用email.ConsumerInit()来启动邮件消费服务
func Init() {
	email.ConsumerInit()
}
