package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/trashwbin/gomall-demo/app/cart/biz/dal/mysql"
	"github.com/trashwbin/gomall-demo/app/cart/biz/model"
	cart "github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/cart"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	err = model.EmptyCart(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	return &cart.EmptyCartResp{}, nil
}
