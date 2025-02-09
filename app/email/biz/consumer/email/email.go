package email

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"github.com/trashwbin/gomall-demo/app/email/infra/mq"
	"github.com/trashwbin/gomall-demo/app/email/infra/notify"
	"google.golang.org/protobuf/proto"

	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/email"
)

// ConsumerInit 初始化邮件消费者。
// 该函数连接到服务器，订阅 'email' 频道，并监听消息以发送邮件。
// 当接收到消息时，它将消息数据反序列化为邮件请求对象，并使用 noop 邮件发送器发送邮件。
// 在关闭时，它会取消订阅频道并关闭连接。
func ConsumerInit() {
	// 订阅 'email' 频道并定义消息处理逻辑
	sub, err := mq.Nc.Subscribe("email", func(m *nats.Msg) {
		// 将接收到的消息数据反序列化为邮件请求对象
		var req email.EmailReq
		err := proto.Unmarshal(m.Data, &req)
		if err != nil {
			klog.Error(err)
		}
		// 创建一个新的 noop 邮件发送器实例
		noopEmail := notify.NewNoopEmail()
		// 使用 noop 邮件发送器发送邮件
		_ = noopEmail.Send(&req)
	})
	if err != nil {
		panic(err)
	}

	// 注册一个关闭钩子，在关闭时取消订阅频道并关闭连接
	server.RegisterShutdownHook(func() {
		// 取消订阅频道
		_ = sub.Unsubscribe() //nolint:errcheck
		// 关闭与服务器的连接
		mq.Nc.Close()
	})
}
