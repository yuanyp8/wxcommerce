// Code generated by Kitex v0.9.0. DO NOT EDIT.

package productverifyrecordservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	product_verify_record "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_verify_record"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ProductVerifyRecordAdd": kitex.NewMethodInfo(
		productVerifyRecordAddHandler,
		newProductVerifyRecordAddArgs,
		newProductVerifyRecordAddResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductVerifyRecordList": kitex.NewMethodInfo(
		productVerifyRecordListHandler,
		newProductVerifyRecordListArgs,
		newProductVerifyRecordListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductVerifyRecordUpdate": kitex.NewMethodInfo(
		productVerifyRecordUpdateHandler,
		newProductVerifyRecordUpdateArgs,
		newProductVerifyRecordUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductVerifyRecordDelete": kitex.NewMethodInfo(
		productVerifyRecordDeleteHandler,
		newProductVerifyRecordDeleteArgs,
		newProductVerifyRecordDeleteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	productVerifyRecordServiceServiceInfo                = NewServiceInfo()
	productVerifyRecordServiceServiceInfoForClient       = NewServiceInfoForClient()
	productVerifyRecordServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return productVerifyRecordServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return productVerifyRecordServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return productVerifyRecordServiceServiceInfoForClient
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
	serviceName := "ProductVerifyRecordService"
	handlerType := (*product_verify_record.ProductVerifyRecordService)(nil)
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

func productVerifyRecordAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_verify_record.ProductVerifyRecordAddReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_verify_record.ProductVerifyRecordService).ProductVerifyRecordAdd(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductVerifyRecordAddArgs:
		success, err := handler.(product_verify_record.ProductVerifyRecordService).ProductVerifyRecordAdd(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductVerifyRecordAddResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductVerifyRecordAddArgs() interface{} {
	return &ProductVerifyRecordAddArgs{}
}

func newProductVerifyRecordAddResult() interface{} {
	return &ProductVerifyRecordAddResult{}
}

type ProductVerifyRecordAddArgs struct {
	Req *product_verify_record.ProductVerifyRecordAddReq
}

func (p *ProductVerifyRecordAddArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_verify_record.ProductVerifyRecordAddReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductVerifyRecordAddArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductVerifyRecordAddArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductVerifyRecordAddArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductVerifyRecordAddArgs) Unmarshal(in []byte) error {
	msg := new(product_verify_record.ProductVerifyRecordAddReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductVerifyRecordAddArgs_Req_DEFAULT *product_verify_record.ProductVerifyRecordAddReq

func (p *ProductVerifyRecordAddArgs) GetReq() *product_verify_record.ProductVerifyRecordAddReq {
	if !p.IsSetReq() {
		return ProductVerifyRecordAddArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductVerifyRecordAddArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductVerifyRecordAddArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductVerifyRecordAddResult struct {
	Success *product_verify_record.ProductVerifyRecordAddResp
}

var ProductVerifyRecordAddResult_Success_DEFAULT *product_verify_record.ProductVerifyRecordAddResp

func (p *ProductVerifyRecordAddResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_verify_record.ProductVerifyRecordAddResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductVerifyRecordAddResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductVerifyRecordAddResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductVerifyRecordAddResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductVerifyRecordAddResult) Unmarshal(in []byte) error {
	msg := new(product_verify_record.ProductVerifyRecordAddResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductVerifyRecordAddResult) GetSuccess() *product_verify_record.ProductVerifyRecordAddResp {
	if !p.IsSetSuccess() {
		return ProductVerifyRecordAddResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductVerifyRecordAddResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_verify_record.ProductVerifyRecordAddResp)
}

func (p *ProductVerifyRecordAddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductVerifyRecordAddResult) GetResult() interface{} {
	return p.Success
}

func productVerifyRecordListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_verify_record.ProductVerifyRecordListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_verify_record.ProductVerifyRecordService).ProductVerifyRecordList(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductVerifyRecordListArgs:
		success, err := handler.(product_verify_record.ProductVerifyRecordService).ProductVerifyRecordList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductVerifyRecordListResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductVerifyRecordListArgs() interface{} {
	return &ProductVerifyRecordListArgs{}
}

func newProductVerifyRecordListResult() interface{} {
	return &ProductVerifyRecordListResult{}
}

type ProductVerifyRecordListArgs struct {
	Req *product_verify_record.ProductVerifyRecordListReq
}

func (p *ProductVerifyRecordListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_verify_record.ProductVerifyRecordListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductVerifyRecordListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductVerifyRecordListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductVerifyRecordListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductVerifyRecordListArgs) Unmarshal(in []byte) error {
	msg := new(product_verify_record.ProductVerifyRecordListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductVerifyRecordListArgs_Req_DEFAULT *product_verify_record.ProductVerifyRecordListReq

func (p *ProductVerifyRecordListArgs) GetReq() *product_verify_record.ProductVerifyRecordListReq {
	if !p.IsSetReq() {
		return ProductVerifyRecordListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductVerifyRecordListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductVerifyRecordListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductVerifyRecordListResult struct {
	Success *product_verify_record.ProductVerifyRecordListResp
}

var ProductVerifyRecordListResult_Success_DEFAULT *product_verify_record.ProductVerifyRecordListResp

func (p *ProductVerifyRecordListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_verify_record.ProductVerifyRecordListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductVerifyRecordListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductVerifyRecordListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductVerifyRecordListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductVerifyRecordListResult) Unmarshal(in []byte) error {
	msg := new(product_verify_record.ProductVerifyRecordListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductVerifyRecordListResult) GetSuccess() *product_verify_record.ProductVerifyRecordListResp {
	if !p.IsSetSuccess() {
		return ProductVerifyRecordListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductVerifyRecordListResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_verify_record.ProductVerifyRecordListResp)
}

func (p *ProductVerifyRecordListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductVerifyRecordListResult) GetResult() interface{} {
	return p.Success
}

func productVerifyRecordUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_verify_record.ProductVerifyRecordUpdateReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_verify_record.ProductVerifyRecordService).ProductVerifyRecordUpdate(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductVerifyRecordUpdateArgs:
		success, err := handler.(product_verify_record.ProductVerifyRecordService).ProductVerifyRecordUpdate(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductVerifyRecordUpdateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductVerifyRecordUpdateArgs() interface{} {
	return &ProductVerifyRecordUpdateArgs{}
}

func newProductVerifyRecordUpdateResult() interface{} {
	return &ProductVerifyRecordUpdateResult{}
}

type ProductVerifyRecordUpdateArgs struct {
	Req *product_verify_record.ProductVerifyRecordUpdateReq
}

func (p *ProductVerifyRecordUpdateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_verify_record.ProductVerifyRecordUpdateReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductVerifyRecordUpdateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductVerifyRecordUpdateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductVerifyRecordUpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductVerifyRecordUpdateArgs) Unmarshal(in []byte) error {
	msg := new(product_verify_record.ProductVerifyRecordUpdateReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductVerifyRecordUpdateArgs_Req_DEFAULT *product_verify_record.ProductVerifyRecordUpdateReq

func (p *ProductVerifyRecordUpdateArgs) GetReq() *product_verify_record.ProductVerifyRecordUpdateReq {
	if !p.IsSetReq() {
		return ProductVerifyRecordUpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductVerifyRecordUpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductVerifyRecordUpdateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductVerifyRecordUpdateResult struct {
	Success *product_verify_record.ProductVerifyRecordUpdateResp
}

var ProductVerifyRecordUpdateResult_Success_DEFAULT *product_verify_record.ProductVerifyRecordUpdateResp

func (p *ProductVerifyRecordUpdateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_verify_record.ProductVerifyRecordUpdateResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductVerifyRecordUpdateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductVerifyRecordUpdateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductVerifyRecordUpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductVerifyRecordUpdateResult) Unmarshal(in []byte) error {
	msg := new(product_verify_record.ProductVerifyRecordUpdateResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductVerifyRecordUpdateResult) GetSuccess() *product_verify_record.ProductVerifyRecordUpdateResp {
	if !p.IsSetSuccess() {
		return ProductVerifyRecordUpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductVerifyRecordUpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_verify_record.ProductVerifyRecordUpdateResp)
}

func (p *ProductVerifyRecordUpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductVerifyRecordUpdateResult) GetResult() interface{} {
	return p.Success
}

func productVerifyRecordDeleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_verify_record.ProductVerifyRecordDeleteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_verify_record.ProductVerifyRecordService).ProductVerifyRecordDelete(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductVerifyRecordDeleteArgs:
		success, err := handler.(product_verify_record.ProductVerifyRecordService).ProductVerifyRecordDelete(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductVerifyRecordDeleteResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductVerifyRecordDeleteArgs() interface{} {
	return &ProductVerifyRecordDeleteArgs{}
}

func newProductVerifyRecordDeleteResult() interface{} {
	return &ProductVerifyRecordDeleteResult{}
}

type ProductVerifyRecordDeleteArgs struct {
	Req *product_verify_record.ProductVerifyRecordDeleteReq
}

func (p *ProductVerifyRecordDeleteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_verify_record.ProductVerifyRecordDeleteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductVerifyRecordDeleteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductVerifyRecordDeleteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductVerifyRecordDeleteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductVerifyRecordDeleteArgs) Unmarshal(in []byte) error {
	msg := new(product_verify_record.ProductVerifyRecordDeleteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductVerifyRecordDeleteArgs_Req_DEFAULT *product_verify_record.ProductVerifyRecordDeleteReq

func (p *ProductVerifyRecordDeleteArgs) GetReq() *product_verify_record.ProductVerifyRecordDeleteReq {
	if !p.IsSetReq() {
		return ProductVerifyRecordDeleteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductVerifyRecordDeleteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductVerifyRecordDeleteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductVerifyRecordDeleteResult struct {
	Success *product_verify_record.ProductVerifyRecordDeleteResp
}

var ProductVerifyRecordDeleteResult_Success_DEFAULT *product_verify_record.ProductVerifyRecordDeleteResp

func (p *ProductVerifyRecordDeleteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_verify_record.ProductVerifyRecordDeleteResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductVerifyRecordDeleteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductVerifyRecordDeleteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductVerifyRecordDeleteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductVerifyRecordDeleteResult) Unmarshal(in []byte) error {
	msg := new(product_verify_record.ProductVerifyRecordDeleteResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductVerifyRecordDeleteResult) GetSuccess() *product_verify_record.ProductVerifyRecordDeleteResp {
	if !p.IsSetSuccess() {
		return ProductVerifyRecordDeleteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductVerifyRecordDeleteResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_verify_record.ProductVerifyRecordDeleteResp)
}

func (p *ProductVerifyRecordDeleteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductVerifyRecordDeleteResult) GetResult() interface{} {
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

func (p *kClient) ProductVerifyRecordAdd(ctx context.Context, Req *product_verify_record.ProductVerifyRecordAddReq) (r *product_verify_record.ProductVerifyRecordAddResp, err error) {
	var _args ProductVerifyRecordAddArgs
	_args.Req = Req
	var _result ProductVerifyRecordAddResult
	if err = p.c.Call(ctx, "ProductVerifyRecordAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductVerifyRecordList(ctx context.Context, Req *product_verify_record.ProductVerifyRecordListReq) (r *product_verify_record.ProductVerifyRecordListResp, err error) {
	var _args ProductVerifyRecordListArgs
	_args.Req = Req
	var _result ProductVerifyRecordListResult
	if err = p.c.Call(ctx, "ProductVerifyRecordList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductVerifyRecordUpdate(ctx context.Context, Req *product_verify_record.ProductVerifyRecordUpdateReq) (r *product_verify_record.ProductVerifyRecordUpdateResp, err error) {
	var _args ProductVerifyRecordUpdateArgs
	_args.Req = Req
	var _result ProductVerifyRecordUpdateResult
	if err = p.c.Call(ctx, "ProductVerifyRecordUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductVerifyRecordDelete(ctx context.Context, Req *product_verify_record.ProductVerifyRecordDeleteReq) (r *product_verify_record.ProductVerifyRecordDeleteResp, err error) {
	var _args ProductVerifyRecordDeleteArgs
	_args.Req = Req
	var _result ProductVerifyRecordDeleteResult
	if err = p.c.Call(ctx, "ProductVerifyRecordDelete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
