package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/trashwbin/gomall-demo/app/cart/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/app/cart/biz/model"
	"github.com/trashwbin/gomall-demo/app/cart/infra/rpc"
	cart "github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/cart"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.GetProductId()})
	if err != nil {
		return nil, err
	}

	if getProduct.Product == nil || getProduct.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not exist")
	}

	err = model.AddCart(mysql.DB, s.ctx, &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	})

	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.AddItemResp{}, nil
}
