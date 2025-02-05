package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/trashwbin/gomall-demo/app/frontend/biz/service"
	"github.com/trashwbin/gomall-demo/app/frontend/biz/utils"
	auth "github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/auth"
	common "github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/common"
)

// Login 处理用户登录请求
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq

	// 绑定并验证请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		// 如果绑定或验证失败，发送错误响应
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 创建LoginService实例并执行登录逻辑
	redirect, err := service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		// 如果登录失败，发送错误响应
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 登录成功后，重定向到指定路径
	c.Redirect(consts.StatusFound, []byte(redirect))
}

// Register 处理用户注册请求
// @router /auth/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterReq

	// 绑定并验证请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		// 如果绑定或验证失败，发送错误响应
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 创建RegisterService实例并执行注册逻辑
	_, err = service.NewRegisterService(ctx, c).Run(&req)
	if err != nil {
		// 如果注册失败，发送错误响应
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 注册成功后，重定向到根路径
	c.Redirect(consts.StatusFound, []byte("/"))
}

// Logout 处理用户登出请求
// @router /auth/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty

	// 绑定并验证请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		// 如果绑定或验证失败，发送错误响应
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 创建LogoutService实例并执行登出逻辑
	_, err = service.NewLogoutService(ctx, c).Run(&req)
	if err != nil {
		// 如果登出失败，发送错误响应
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 登出成功后，重定向到根路径
	c.Redirect(consts.StatusFound, []byte("/"))
}
