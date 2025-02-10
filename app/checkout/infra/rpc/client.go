package rpc

import (
	"github.com/trashwbin/gomall-demo/common/clientsuite"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/order/orderservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/trashwbin/gomall-demo/app/checkout/conf"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	if err != nil {
		panic(err)
	}
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	if err != nil {
		panic(err)
	}
}
func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	if err != nil {
		panic(err)
	}
}
