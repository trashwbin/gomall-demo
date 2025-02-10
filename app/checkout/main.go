package main

import (
	"context"
	"github.com/trashwbin/gomall-demo/app/checkout/infra/mq"
	"github.com/trashwbin/gomall-demo/app/checkout/infra/rpc"
	"github.com/trashwbin/gomall-demo/common/mtl"
	"github.com/trashwbin/gomall-demo/common/serversuite"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/trashwbin/gomall-demo/app/checkout/conf"
	"github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/checkout/checkoutservice"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background())
	opts := kitexInit()
	rpc.InitClient()
	mq.Init()
	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

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
