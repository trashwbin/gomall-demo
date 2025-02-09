package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/cart"
	common "github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/common"
	"github.com/trashwbin/gomall-demo/app/frontend/infra/rpc"
	frontendutils "github.com/trashwbin/gomall-demo/app/frontend/utils"
	rpccart "github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/cart"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: frontendutils.GetUserIdFromCtx(h.Context),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	return
}
