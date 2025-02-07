package rpc

import (
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/trashwbin/gomall-demo/app/frontend/conf"
	frontendutils "github.com/trashwbin/gomall-demo/app/frontend/utils"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/user/userservice"
)

var (
	// UserClient 是用户服务的RPC客户端实例
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
	// once 确保InitClient函数只执行一次
	once sync.Once
)

// InitClient 初始化RPC客户端，确保只执行一次
func InitClient() {
	once.Do(func() {
		initUserClient()
		initCartClient()
		initProductClient()
	})
}

// initUserClient 初始化用户服务的RPC客户端
func initUserClient() {
	// 创建Consul解析器，用于从Consul获取服务地址
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendutils.MustHandleError(err) // 处理可能发生的错误

	// 使用解析器创建用户服务的RPC客户端
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendutils.MustHandleError(err) // 处理可能发生的错误
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendutils.MustHandleError(err)
}
func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendutils.MustHandleError(err)
}
