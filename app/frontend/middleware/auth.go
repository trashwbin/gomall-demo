package middleware

import (
	"context"

	frontendUtils "github.com/trashwbin/gomall-demo/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

// GlobalAuth 是一个全局认证中间件，用于将用户ID添加到上下文中
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取当前请求的会话对象
		s := sessions.Default(c)

		// 将用户ID从会话中取出并添加到上下文中
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id"))

		// 调用下一个处理函数
		c.Next(ctx)
	}
}

// Auth 是一个认证中间件，用于检查用户是否已登录
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取当前请求的会话对象
		s := sessions.Default(c)

		// 从会话中获取用户ID
		userId := s.Get("user_id")

		// 如果用户ID为空，表示用户未登录，则重定向到登录页面，并携带当前请求路径作为next参数
		if userId == nil {
			c.Redirect(consts.StatusFound, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}

		// 用户已登录，继续调用下一个处理函数
		c.Next(ctx)
	}
}
