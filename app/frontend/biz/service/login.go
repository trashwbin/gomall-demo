package service

import (
	"context"
	"github.com/hertz-contrib/sessions"
	"github.com/trashwbin/gomall-demo/app/frontend/infra/rpc"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/auth"
)

type LoginService struct {
	// RequestContext 是当前请求的上下文
	RequestContext *app.RequestContext
	// Context 是用于处理请求的上下文
	Context context.Context
}

// NewLoginService 创建一个新的LoginService实例
func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

// Run 处理登录请求并返回重定向路径
func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	// 调用用户服务的Login方法进行登录验证
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}

	// 获取当前请求的会话对象
	session := sessions.Default(h.RequestContext)

	// 将用户ID存储到会话中
	session.Set("user_id", resp.UserId)

	// 保存会话
	err = session.Save()
	if err != nil {
		return "", err
	}

	// 设置重定向路径，默认为根路径"/"
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}

	return redirect, nil
}
