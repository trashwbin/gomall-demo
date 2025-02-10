package rpc

import (
	"github.com/trashwbin/gomall-demo/common/clientsuite"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/trashwbin/gomall-demo/app/cart/conf"
	cartutils "github.com/trashwbin/gomall-demo/app/cart/utils"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	cartutils.MustHandleError(err)
}
