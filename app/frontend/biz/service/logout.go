package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/common"
)

type LogoutService struct {
	// RequestContext 是当前请求的上下文
	RequestContext *app.RequestContext
	// Context 是用于处理请求的上下文
	Context context.Context
}

// NewLogoutService 创建一个新的LogoutService实例
func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

// Run 处理登出逻辑
func (h *LogoutService) Run(req *common.Empty) (resp *common.Empty, err error) {
	// 获取当前请求的会话对象
	session := sessions.Default(h.RequestContext)

	// 清除会话中的所有数据
	session.Clear()

	// 保存会话
	err = session.Save()
	if err != nil {
		// 如果保存会话失败，返回错误
		return nil, err
	}

	// 登出成功，返回空响应
	return
}
