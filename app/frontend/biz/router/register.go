// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	auth "github.com/trashwbin/gomall-demo/app/frontend/biz/router/auth"
	cart "github.com/trashwbin/gomall-demo/app/frontend/biz/router/cart"
	category "github.com/trashwbin/gomall-demo/app/frontend/biz/router/category"
	checkout "github.com/trashwbin/gomall-demo/app/frontend/biz/router/checkout"
	home "github.com/trashwbin/gomall-demo/app/frontend/biz/router/home"
	order "github.com/trashwbin/gomall-demo/app/frontend/biz/router/order"
	"github.com/trashwbin/gomall-demo/app/frontend/biz/router/product"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	order.Register(r)

	checkout.Register(r)

	cart.Register(r)

	product.Register(r)

	category.Register(r)

	auth.Register(r)

	home.Register(r)
}
