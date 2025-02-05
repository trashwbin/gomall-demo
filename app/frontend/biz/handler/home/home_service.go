package home

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/trashwbin/gomall-demo/app/frontend/biz/service"
	"github.com/trashwbin/gomall-demo/app/frontend/biz/utils"
	common "github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/common"
)

// Home 处理首页请求
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty

	// 绑定并验证请求参数
	err = c.BindAndValidate(&req)
	if err != nil {
		// 如果绑定或验证失败，发送错误响应
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 创建HomeService实例并执行首页逻辑
	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		// 如果首页逻辑执行失败，发送错误响应
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 包装响应内容并渲染首页模板
	c.HTML(consts.StatusOK, "home", utils.WarpResponse(ctx, c, resp))
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
