package serversuite

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/trashwbin/gomall-demo/common/mtl"
)

// CommonServerSuite 是一个通用的服务套件结构体，用于封装服务的基本配置和选项。
type CommonServerSuite struct {
	CurrentServiceName string // 当前服务的名称，用于标识服务
	RegistryAddr       string // Consul 注册地址，用于服务发现
}

// Options 返回一组服务器选项，这些选项用于配置服务的行为。
// 该方法实现了 server.Option 接口，允许在启动服务时传递自定义配置。
func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		// 添加元数据处理器，处理 HTTP/2 请求中的元数据信息。
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),

		// 设置服务的基本信息，包括服务名称。
		// 这里使用了 s.CurrentServiceName 来动态设置服务名称。
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),

		// 配置 Prometheus Tracer，用于追踪服务请求。
		// 禁用服务器端追踪，并指定使用 mtl.Registry 作为 Prometheus 的注册表。
		server.WithTracer(prometheus.NewServerTracer("192.168.1.35", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),

		server.WithSuite(tracing.NewServerSuite()), // 配置 OpenTelemetry Tracing 套件。
	}

	register, err := consul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(register))

	return opts
}
