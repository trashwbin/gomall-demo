package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/kitex_gen/pbapi"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/kitex_gen/pbapi/echo"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/middleware"
	"log"
)

func main() {
	// 创建 Consul 解析器
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}

	// 创建 Kitex 客户端
	c, err := echo.NewClient("demo_proto",
		client.WithResolver(r),
		//client.WithTransportProtocol(transport.GRPC), // windows11 中grpc支持有问题，导致panic
		client.WithShortConnection(), // 使用短连接
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware), // 添加中间件
	)

	if err != nil {
		log.Fatal(err)
	}

	// 创建上下文，对于grpc，key需要大写
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")

	// 调用 Echo 方法
	//res, err := c.Echo(ctx, &pbapi.Request{Message: "hello"})
	res, err := c.Echo(ctx, &pbapi.Request{Message: "error"})

	//var bizErr *kerrors.GRPCBizStatusError
	var bizErr *kerrors.BizStatusError

	if err != nil {
		if errors.As(err, &bizErr) {
			fmt.Printf("%#v", bizErr)
			fmt.Println("111")
			return
		}
		log.Fatal(err)
	}
	log.Println(res)
	log.Fatal(err)

}
