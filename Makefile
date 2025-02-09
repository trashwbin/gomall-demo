#windows 下不直接支持make命令，需要安装mingw-w64，建议直接复制到命令行执行
#抽取后在cmd使用时需手动替换为ROOT_MOD
export ROOT_MOD = github.com/trashwbin/gomall-demo
.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module ${ROOT_MOD}/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server -I ../../idl --type RPC --module ${ROOT_MOD}/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: gen-frontend-home
gen-frontend-home:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/checkout_page.proto --service frontend -module github.com/trashwbin/gomall-demo/app/frontend -I ../../idl
.PHONY: gen-frontend-auth
gen-frontend-auth:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module github.com/trashwbin/gomall-demo/app/frontend -I ../../idl


#使用cwgo生成user服务的客户端代码, 生成的代码在rpc_gen/user目录下,所有的客户端代码都在rpc_gen目录下
.PHONY: gen-user-client
gen-user-client:
	@cd rpc_gen && cwgo client  --type RPC --service user --module github.com/trashwbin/gomall-demo/rpc_gen --I ../idl --idl ../idl/user.proto

#服务端代码生成在app/user目录下，客户端代码生成在rpc_gen/user目录下，我们指定服务端引用的一些实现的接口，这里是kitex_gen
# --pass 是cwgo的一个选项，将后续命令直接传递给底层的代码生成工具，这里是kitex_gen
# --use 是kitex_gen的一个选项，可以控制服务端生成的时候不再去生成客户端的代码，而直接使用指定的这个模块
# 这样做的目的是为了app下所有的微服务不再维护客户端代码，而把客户端代码全部集中在rpc_gen目录下，这样可以更好的维护和管理
.PHONY: gen-user-server
gen-user-server:
	@cd app/user && cwgo server --type RPC --service user --module github.com/trashwbin/gomall-demo/app/user --pass "-use github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

