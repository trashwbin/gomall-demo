package service

import (
	"context"
	"github.com/hertz-contrib/sessions"
	"github.com/trashwbin/gomall-demo/app/frontend/infra/rpc"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/auth"
	common "github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/common"
)

type RegisterService struct {
	// RequestContext 是当前请求的上下文
	RequestContext *app.RequestContext
	// Context 是用于处理请求的上下文
	Context context.Context
}

// NewRegisterService 创建一个新的RegisterService实例
func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

// Run 处理用户注册逻辑
func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	// 调用用户服务的Register方法进行用户注册
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		// 如果注册失败，返回错误
		return nil, err
	}

	// 获取当前请求的会话对象
	session := sessions.Default(h.RequestContext)

	// 将用户ID存储到会话中
	session.Set("user_id", userResp.UserId)

	// 保存会话
	err = session.Save()
	if err != nil {
		// 如果保存会话失败，返回错误
		return nil, err
	}

	// 注册成功，返回空响应
	return
}
