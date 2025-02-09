package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/trashwbin/gomall-demo/app/frontend/infra/rpc"
	frontendutils "github.com/trashwbin/gomall-demo/app/frontend/utils"
	rpccart "github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/product"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/trashwbin/gomall-demo/app/frontend/hertz_gen/frontend/common"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	var items []map[string]string
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	var total float32

	for _, v := range carts.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: v.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
		total += float32(v.Quantity) * p.Price
	}

	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
