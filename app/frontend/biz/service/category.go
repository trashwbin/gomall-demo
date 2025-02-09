package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	category "github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/category"
	"github.com/trashwbin/gomall-demo/app/frontend/infra/rpc"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/product"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	p, _ := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	return utils.H{
		"title":    "Category",
		"items":    p.Products,
		"cart_num": 10,
	}, nil
}
