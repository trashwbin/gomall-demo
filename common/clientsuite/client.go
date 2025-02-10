// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

// CommonGrpcClientSuite 是一个通用的 gRPC 客户端套件结构体，用于封装客户端的基本配置和选项。
type CommonGrpcClientSuite struct {
	CurrentServiceName string // 当前服务的名称，用于标识客户端连接的服务
	RegistryAddr       string // Consul 注册地址，用于服务发现
}

// Options 返回一组客户端选项，这些选项用于配置 gRPC 客户端的行为。
// 该方法实现了 client.Option 接口，允许在创建客户端时传递自定义配置。
func (s CommonGrpcClientSuite) Options() []client.Option {
	opts := []client.Option{
		// 添加元数据处理器，处理 HTTP/2 请求中的元数据信息。
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),

		// 设置传输协议为 gRPC。
		client.WithTransportProtocol(transport.GRPC),

		// 设置客户端的基本信息，包括服务名称。
		// 这里使用了 s.CurrentServiceName 来动态设置服务名称。
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
	}
	// 创建一个新的 Consul 解析器，用于解析服务地址。
	// 参数：
	// - s.RegistryAddr: Consul 注册地址，用于连接到 Consul 服务器。
	resolver, err := consul.NewConsulResolver(s.RegistryAddr)
	if err != nil {
		// 如果创建解析器时发生错误，则触发 panic 并终止程序，输出错误信息。
		panic(err)
	}

	// 将解析器添加到客户端选项列表中，以便客户端可以使用 Consul 进行服务发现。
	opts = append(opts, client.WithResolver(resolver))

	return opts
}
