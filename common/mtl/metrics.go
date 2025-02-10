package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Registry 是一个全局变量，用于存储 Prometheus 的指标注册表。
var Registry *prometheus.Registry

// InitMetric 初始化并配置 Prometheus 指标服务。
// 参数：
// - serviceName: 服务名称，用于标识当前服务。
// - metricsPort: 指标服务监听的端口。
// - registryAddr: Consul 注册地址，用于服务发现。
func InitMetric(serviceName string, metricsPort string, registryAddr string) (registry.Registry, *registry.Info) {
	// 创建一个新的 Prometheus 注册表，并将其赋值给全局变量 Registry。
	Registry = prometheus.NewRegistry()

	// 注册 Go 运行时的默认收集器，用于收集 Go 运行时的指标。
	Registry.MustRegister(collectors.NewGoCollector())

	// 注册进程级别的收集器，用于收集进程相关的指标。
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	// 创建一个新的 Consul 客户端，用于将指标服务注册到 Consul。
	r, _ := consul.NewConsulRegister(registryAddr)

	// 解析 metricsPort 字符串为 TCP 地址，以便后续使用。
	addr, _ := net.ResolveTCPAddr("tcp", "192.168.1.35"+metricsPort)

	// 构建要注册的服务信息。
	registryInfo := &registry.Info{
		ServiceName: "prometheus",                              // 服务名称为 "prometheus"
		Addr:        addr,                                      // 服务地址为解析后的 TCP 地址
		Weight:      1,                                         // 权重设置为 1
		Tags:        map[string]string{"service": serviceName}, // 添加标签以标识服务名称
	}

	// 将指标服务注册到 Consul。
	_ = r.Register(registryInfo)

	// 注册一个关闭钩子，在程序退出时注销 Consul 中的服务。
	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo)
	})

	// 设置 HTTP 路由，处理 "/metrics" 请求，返回 Prometheus 格式的指标数据。
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))

	// 启动 HTTP 服务器，监听指定的 metricsPort 端口，提供指标服务。
	// 注意：这里忽略了启动服务器时可能发生的错误。建议在生产环境中处理该错误。
	go http.ListenAndServe(metricsPort, nil)

	return r, registryInfo
}
