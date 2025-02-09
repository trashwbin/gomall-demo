// Code generated by Kitex v0.9.1. DO NOT EDIT.

package checkoutservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	checkout "github.com/trashwbin/gomall-demo/rpc_gen/kitex_gen/checkout"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Checkout": kitex.NewMethodInfo(
		checkoutHandler,
		newCheckoutArgs,
		newCheckoutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	checkoutServiceServiceInfo                = NewServiceInfo()
	checkoutServiceServiceInfoForClient       = NewServiceInfoForClient()
	checkoutServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return checkoutServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return checkoutServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return checkoutServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "CheckoutService"
	handlerType := (*checkout.CheckoutService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "checkout",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func checkoutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(checkout.CheckoutReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(checkout.CheckoutService).Checkout(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CheckoutArgs:
		success, err := handler.(checkout.CheckoutService).Checkout(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckoutResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCheckoutArgs() interface{} {
	return &CheckoutArgs{}
}

func newCheckoutResult() interface{} {
	return &CheckoutResult{}
}

type CheckoutArgs struct {
	Req *checkout.CheckoutReq
}

func (p *CheckoutArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(checkout.CheckoutReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CheckoutArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CheckoutArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CheckoutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CheckoutArgs) Unmarshal(in []byte) error {
	msg := new(checkout.CheckoutReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckoutArgs_Req_DEFAULT *checkout.CheckoutReq

func (p *CheckoutArgs) GetReq() *checkout.CheckoutReq {
	if !p.IsSetReq() {
		return CheckoutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckoutArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CheckoutArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CheckoutResult struct {
	Success *checkout.CheckoutResp
}

var CheckoutResult_Success_DEFAULT *checkout.CheckoutResp

func (p *CheckoutResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(checkout.CheckoutResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CheckoutResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CheckoutResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CheckoutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CheckoutResult) Unmarshal(in []byte) error {
	msg := new(checkout.CheckoutResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckoutResult) GetSuccess() *checkout.CheckoutResp {
	if !p.IsSetSuccess() {
		return CheckoutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckoutResult) SetSuccess(x interface{}) {
	p.Success = x.(*checkout.CheckoutResp)
}

func (p *CheckoutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckoutResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Checkout(ctx context.Context, Req *checkout.CheckoutReq) (r *checkout.CheckoutResp, err error) {
	var _args CheckoutArgs
	_args.Req = Req
	var _result CheckoutResult
	if err = p.c.Call(ctx, "Checkout", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
