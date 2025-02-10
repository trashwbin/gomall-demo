package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/trashwbin/gomall-demo/app/cart/biz/dal"
	"github.com/trashwbin/gomall-demo/app/cart/infra/rpc"
	"github.com/trashwbin/gomall-demo/common/mtl"
	"github.com/trashwbin/gomall-demo/common/serversuite"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/trashwbin/gomall-demo/app/cart/conf"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/cart/cartservice"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	_ = godotenv.Load()
	// 初始化 Prometheus 指标服务，配置并启动指标收集和暴露。
	// - ServiceName: 当前服务的名称，用于标识该服务。
	// - conf.GetConf().Kitex.MetricsPort: 配置文件中指定的指标服务监听端口。
	// - RegistryAddr: Consul 注册地址，用于将指标服务注册到 Consul。
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background())
	dal.Init()
	rpc.InitClient()
	opts := kitexInit()

	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// 将 CommonServerSuite 的配置选项添加到服务器选项列表中。
	// 这里使用了 serversuite.CommonServerSuite 结构体来封装服务的基本配置和选项。
	opts = append(opts,
		// 添加一个通用的服务套件，用于配置服务的行为。
		server.WithSuite(serversuite.CommonServerSuite{
			CurrentServiceName: ServiceName,  // 当前服务的名称，用于标识该服务
			RegistryAddr:       RegistryAddr, // Consul 注册地址，用于将服务注册到 Consul
		}),
	)

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
