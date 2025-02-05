package service

import (
	"context"
	"github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	// RequestContext 是当前请求的上下文
	RequestContext *app.RequestContext
	// Context 是用于处理请求的上下文
	Context context.Context
}

// NewHomeService 创建一个新的HomeService实例
func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

// Run 处理首页逻辑并返回响应内容
func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	// 创建一个响应映射
	var resp = make(map[string]any)

	// 定义首页显示的商品列表
	items := []map[string]any{
		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt-2", "Price": 110, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt-3", "Price": 120, "Picture": "/static/image/t-shirt-2.jpeg"},
		{"Name": "T-shirt-4", "Price": 130, "Picture": "/static/image/notebook.jpeg"},
		{"Name": "T-shirt-5", "Price": 140, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt-6", "Price": 150, "Picture": "/static/image/t-shirt.jpeg"},
	}

	// 设置响应中的标题
	resp["title"] = "Hot Sales"

	// 设置响应中的商品列表
	resp["items"] = items

	// 返回响应内容和错误（如果有的话）
	return resp, nil
}
