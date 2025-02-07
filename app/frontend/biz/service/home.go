package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/common"
	"github.com/trashwbin/gomall-demo/app/frontend/infra/rpc"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/product"

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
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Hot sale",
		"items": products.Products,
	}, nil
}
