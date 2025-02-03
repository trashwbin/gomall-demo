package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/kitex_gen/pbapi"
	"github.com/trashwbin/gomall-demo/demo/demo_proto/kitex_gen/pbapi/echo"
	"log"
)

func main() {
	// 创建 Consul 解析器
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}

	// 创建 Kitex 客户端
	c, err := echo.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	// 调用 Echo 方法
	res, err := c.Echo(context.TODO(), &pbapi.Request{Message: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
