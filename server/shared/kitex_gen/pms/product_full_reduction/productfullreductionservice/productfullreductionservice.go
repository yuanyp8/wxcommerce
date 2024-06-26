// Code generated by Kitex v0.9.0. DO NOT EDIT.

package productfullreductionservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	product_full_reduction "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_full_reduction"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ProductFullReductionAdd": kitex.NewMethodInfo(
		productFullReductionAddHandler,
		newProductFullReductionAddArgs,
		newProductFullReductionAddResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductFullReductionList": kitex.NewMethodInfo(
		productFullReductionListHandler,
		newProductFullReductionListArgs,
		newProductFullReductionListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductFullReductionUpdate": kitex.NewMethodInfo(
		productFullReductionUpdateHandler,
		newProductFullReductionUpdateArgs,
		newProductFullReductionUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductFullReductionDelete": kitex.NewMethodInfo(
		productFullReductionDeleteHandler,
		newProductFullReductionDeleteArgs,
		newProductFullReductionDeleteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	productFullReductionServiceServiceInfo                = NewServiceInfo()
	productFullReductionServiceServiceInfoForClient       = NewServiceInfoForClient()
	productFullReductionServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return productFullReductionServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return productFullReductionServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return productFullReductionServiceServiceInfoForClient
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
	serviceName := "ProductFullReductionService"
	handlerType := (*product_full_reduction.ProductFullReductionService)(nil)
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
		"PackageName": "pms",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.0",
		Extra:           extra,
	}
	return svcInfo
}

func productFullReductionAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_full_reduction.ProductFullReductionAddReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_full_reduction.ProductFullReductionService).ProductFullReductionAdd(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductFullReductionAddArgs:
		success, err := handler.(product_full_reduction.ProductFullReductionService).ProductFullReductionAdd(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductFullReductionAddResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductFullReductionAddArgs() interface{} {
	return &ProductFullReductionAddArgs{}
}

func newProductFullReductionAddResult() interface{} {
	return &ProductFullReductionAddResult{}
}

type ProductFullReductionAddArgs struct {
	Req *product_full_reduction.ProductFullReductionAddReq
}

func (p *ProductFullReductionAddArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_full_reduction.ProductFullReductionAddReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductFullReductionAddArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductFullReductionAddArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductFullReductionAddArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductFullReductionAddArgs) Unmarshal(in []byte) error {
	msg := new(product_full_reduction.ProductFullReductionAddReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductFullReductionAddArgs_Req_DEFAULT *product_full_reduction.ProductFullReductionAddReq

func (p *ProductFullReductionAddArgs) GetReq() *product_full_reduction.ProductFullReductionAddReq {
	if !p.IsSetReq() {
		return ProductFullReductionAddArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductFullReductionAddArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductFullReductionAddArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductFullReductionAddResult struct {
	Success *product_full_reduction.ProductFullReductionAddResp
}

var ProductFullReductionAddResult_Success_DEFAULT *product_full_reduction.ProductFullReductionAddResp

func (p *ProductFullReductionAddResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_full_reduction.ProductFullReductionAddResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductFullReductionAddResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductFullReductionAddResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductFullReductionAddResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductFullReductionAddResult) Unmarshal(in []byte) error {
	msg := new(product_full_reduction.ProductFullReductionAddResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductFullReductionAddResult) GetSuccess() *product_full_reduction.ProductFullReductionAddResp {
	if !p.IsSetSuccess() {
		return ProductFullReductionAddResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductFullReductionAddResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_full_reduction.ProductFullReductionAddResp)
}

func (p *ProductFullReductionAddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductFullReductionAddResult) GetResult() interface{} {
	return p.Success
}

func productFullReductionListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_full_reduction.ProductFullReductionListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_full_reduction.ProductFullReductionService).ProductFullReductionList(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductFullReductionListArgs:
		success, err := handler.(product_full_reduction.ProductFullReductionService).ProductFullReductionList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductFullReductionListResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductFullReductionListArgs() interface{} {
	return &ProductFullReductionListArgs{}
}

func newProductFullReductionListResult() interface{} {
	return &ProductFullReductionListResult{}
}

type ProductFullReductionListArgs struct {
	Req *product_full_reduction.ProductFullReductionListReq
}

func (p *ProductFullReductionListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_full_reduction.ProductFullReductionListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductFullReductionListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductFullReductionListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductFullReductionListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductFullReductionListArgs) Unmarshal(in []byte) error {
	msg := new(product_full_reduction.ProductFullReductionListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductFullReductionListArgs_Req_DEFAULT *product_full_reduction.ProductFullReductionListReq

func (p *ProductFullReductionListArgs) GetReq() *product_full_reduction.ProductFullReductionListReq {
	if !p.IsSetReq() {
		return ProductFullReductionListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductFullReductionListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductFullReductionListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductFullReductionListResult struct {
	Success *product_full_reduction.ProductFullReductionListResp
}

var ProductFullReductionListResult_Success_DEFAULT *product_full_reduction.ProductFullReductionListResp

func (p *ProductFullReductionListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_full_reduction.ProductFullReductionListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductFullReductionListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductFullReductionListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductFullReductionListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductFullReductionListResult) Unmarshal(in []byte) error {
	msg := new(product_full_reduction.ProductFullReductionListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductFullReductionListResult) GetSuccess() *product_full_reduction.ProductFullReductionListResp {
	if !p.IsSetSuccess() {
		return ProductFullReductionListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductFullReductionListResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_full_reduction.ProductFullReductionListResp)
}

func (p *ProductFullReductionListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductFullReductionListResult) GetResult() interface{} {
	return p.Success
}

func productFullReductionUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_full_reduction.ProductFullReductionUpdateReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_full_reduction.ProductFullReductionService).ProductFullReductionUpdate(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductFullReductionUpdateArgs:
		success, err := handler.(product_full_reduction.ProductFullReductionService).ProductFullReductionUpdate(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductFullReductionUpdateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductFullReductionUpdateArgs() interface{} {
	return &ProductFullReductionUpdateArgs{}
}

func newProductFullReductionUpdateResult() interface{} {
	return &ProductFullReductionUpdateResult{}
}

type ProductFullReductionUpdateArgs struct {
	Req *product_full_reduction.ProductFullReductionUpdateReq
}

func (p *ProductFullReductionUpdateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_full_reduction.ProductFullReductionUpdateReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductFullReductionUpdateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductFullReductionUpdateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductFullReductionUpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductFullReductionUpdateArgs) Unmarshal(in []byte) error {
	msg := new(product_full_reduction.ProductFullReductionUpdateReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductFullReductionUpdateArgs_Req_DEFAULT *product_full_reduction.ProductFullReductionUpdateReq

func (p *ProductFullReductionUpdateArgs) GetReq() *product_full_reduction.ProductFullReductionUpdateReq {
	if !p.IsSetReq() {
		return ProductFullReductionUpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductFullReductionUpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductFullReductionUpdateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductFullReductionUpdateResult struct {
	Success *product_full_reduction.ProductFullReductionUpdateResp
}

var ProductFullReductionUpdateResult_Success_DEFAULT *product_full_reduction.ProductFullReductionUpdateResp

func (p *ProductFullReductionUpdateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_full_reduction.ProductFullReductionUpdateResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductFullReductionUpdateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductFullReductionUpdateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductFullReductionUpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductFullReductionUpdateResult) Unmarshal(in []byte) error {
	msg := new(product_full_reduction.ProductFullReductionUpdateResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductFullReductionUpdateResult) GetSuccess() *product_full_reduction.ProductFullReductionUpdateResp {
	if !p.IsSetSuccess() {
		return ProductFullReductionUpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductFullReductionUpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_full_reduction.ProductFullReductionUpdateResp)
}

func (p *ProductFullReductionUpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductFullReductionUpdateResult) GetResult() interface{} {
	return p.Success
}

func productFullReductionDeleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_full_reduction.ProductFullReductionDeleteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_full_reduction.ProductFullReductionService).ProductFullReductionDelete(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductFullReductionDeleteArgs:
		success, err := handler.(product_full_reduction.ProductFullReductionService).ProductFullReductionDelete(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductFullReductionDeleteResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductFullReductionDeleteArgs() interface{} {
	return &ProductFullReductionDeleteArgs{}
}

func newProductFullReductionDeleteResult() interface{} {
	return &ProductFullReductionDeleteResult{}
}

type ProductFullReductionDeleteArgs struct {
	Req *product_full_reduction.ProductFullReductionDeleteReq
}

func (p *ProductFullReductionDeleteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_full_reduction.ProductFullReductionDeleteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductFullReductionDeleteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductFullReductionDeleteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductFullReductionDeleteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductFullReductionDeleteArgs) Unmarshal(in []byte) error {
	msg := new(product_full_reduction.ProductFullReductionDeleteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductFullReductionDeleteArgs_Req_DEFAULT *product_full_reduction.ProductFullReductionDeleteReq

func (p *ProductFullReductionDeleteArgs) GetReq() *product_full_reduction.ProductFullReductionDeleteReq {
	if !p.IsSetReq() {
		return ProductFullReductionDeleteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductFullReductionDeleteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductFullReductionDeleteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductFullReductionDeleteResult struct {
	Success *product_full_reduction.ProductFullReductionDeleteResp
}

var ProductFullReductionDeleteResult_Success_DEFAULT *product_full_reduction.ProductFullReductionDeleteResp

func (p *ProductFullReductionDeleteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_full_reduction.ProductFullReductionDeleteResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductFullReductionDeleteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductFullReductionDeleteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductFullReductionDeleteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductFullReductionDeleteResult) Unmarshal(in []byte) error {
	msg := new(product_full_reduction.ProductFullReductionDeleteResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductFullReductionDeleteResult) GetSuccess() *product_full_reduction.ProductFullReductionDeleteResp {
	if !p.IsSetSuccess() {
		return ProductFullReductionDeleteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductFullReductionDeleteResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_full_reduction.ProductFullReductionDeleteResp)
}

func (p *ProductFullReductionDeleteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductFullReductionDeleteResult) GetResult() interface{} {
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

func (p *kClient) ProductFullReductionAdd(ctx context.Context, Req *product_full_reduction.ProductFullReductionAddReq) (r *product_full_reduction.ProductFullReductionAddResp, err error) {
	var _args ProductFullReductionAddArgs
	_args.Req = Req
	var _result ProductFullReductionAddResult
	if err = p.c.Call(ctx, "ProductFullReductionAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductFullReductionList(ctx context.Context, Req *product_full_reduction.ProductFullReductionListReq) (r *product_full_reduction.ProductFullReductionListResp, err error) {
	var _args ProductFullReductionListArgs
	_args.Req = Req
	var _result ProductFullReductionListResult
	if err = p.c.Call(ctx, "ProductFullReductionList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductFullReductionUpdate(ctx context.Context, Req *product_full_reduction.ProductFullReductionUpdateReq) (r *product_full_reduction.ProductFullReductionUpdateResp, err error) {
	var _args ProductFullReductionUpdateArgs
	_args.Req = Req
	var _result ProductFullReductionUpdateResult
	if err = p.c.Call(ctx, "ProductFullReductionUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductFullReductionDelete(ctx context.Context, Req *product_full_reduction.ProductFullReductionDeleteReq) (r *product_full_reduction.ProductFullReductionDeleteResp, err error) {
	var _args ProductFullReductionDeleteArgs
	_args.Req = Req
	var _result ProductFullReductionDeleteResult
	if err = p.c.Call(ctx, "ProductFullReductionDelete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
