package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	frontendUtils "github.com/trashwbin/gomall-demo/app/frontend/utils"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

// WarpResponse 包装响应内容，添加用户ID
func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	// 从上下文中获取用户ID
	content["user_id"] = frontendUtils.GetUserIdFromCtx(ctx)
	return content
}
