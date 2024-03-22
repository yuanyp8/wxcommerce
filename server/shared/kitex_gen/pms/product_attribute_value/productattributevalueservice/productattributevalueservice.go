// Code generated by Kitex v0.9.0. DO NOT EDIT.

package productattributevalueservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	product_attribute_value "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_attribute_value"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ProductAttributeValueAdd": kitex.NewMethodInfo(
		productAttributeValueAddHandler,
		newProductAttributeValueAddArgs,
		newProductAttributeValueAddResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductAttributeValueList": kitex.NewMethodInfo(
		productAttributeValueListHandler,
		newProductAttributeValueListArgs,
		newProductAttributeValueListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductAttributeValueUpdate": kitex.NewMethodInfo(
		productAttributeValueUpdateHandler,
		newProductAttributeValueUpdateArgs,
		newProductAttributeValueUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductAttributeValueDelete": kitex.NewMethodInfo(
		productAttributeValueDeleteHandler,
		newProductAttributeValueDeleteArgs,
		newProductAttributeValueDeleteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	productAttributeValueServiceServiceInfo                = NewServiceInfo()
	productAttributeValueServiceServiceInfoForClient       = NewServiceInfoForClient()
	productAttributeValueServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return productAttributeValueServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return productAttributeValueServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return productAttributeValueServiceServiceInfoForClient
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
	serviceName := "ProductAttributeValueService"
	handlerType := (*product_attribute_value.ProductAttributeValueService)(nil)
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

func productAttributeValueAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_attribute_value.ProductAttributeValueAddReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_attribute_value.ProductAttributeValueService).ProductAttributeValueAdd(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductAttributeValueAddArgs:
		success, err := handler.(product_attribute_value.ProductAttributeValueService).ProductAttributeValueAdd(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductAttributeValueAddResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductAttributeValueAddArgs() interface{} {
	return &ProductAttributeValueAddArgs{}
}

func newProductAttributeValueAddResult() interface{} {
	return &ProductAttributeValueAddResult{}
}

type ProductAttributeValueAddArgs struct {
	Req *product_attribute_value.ProductAttributeValueAddReq
}

func (p *ProductAttributeValueAddArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_attribute_value.ProductAttributeValueAddReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductAttributeValueAddArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductAttributeValueAddArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductAttributeValueAddArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductAttributeValueAddArgs) Unmarshal(in []byte) error {
	msg := new(product_attribute_value.ProductAttributeValueAddReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductAttributeValueAddArgs_Req_DEFAULT *product_attribute_value.ProductAttributeValueAddReq

func (p *ProductAttributeValueAddArgs) GetReq() *product_attribute_value.ProductAttributeValueAddReq {
	if !p.IsSetReq() {
		return ProductAttributeValueAddArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductAttributeValueAddArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductAttributeValueAddArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductAttributeValueAddResult struct {
	Success *product_attribute_value.ProductAttributeValueAddResp
}

var ProductAttributeValueAddResult_Success_DEFAULT *product_attribute_value.ProductAttributeValueAddResp

func (p *ProductAttributeValueAddResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_attribute_value.ProductAttributeValueAddResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductAttributeValueAddResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductAttributeValueAddResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductAttributeValueAddResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductAttributeValueAddResult) Unmarshal(in []byte) error {
	msg := new(product_attribute_value.ProductAttributeValueAddResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductAttributeValueAddResult) GetSuccess() *product_attribute_value.ProductAttributeValueAddResp {
	if !p.IsSetSuccess() {
		return ProductAttributeValueAddResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductAttributeValueAddResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_attribute_value.ProductAttributeValueAddResp)
}

func (p *ProductAttributeValueAddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductAttributeValueAddResult) GetResult() interface{} {
	return p.Success
}

func productAttributeValueListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_attribute_value.ProductAttributeValueListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_attribute_value.ProductAttributeValueService).ProductAttributeValueList(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductAttributeValueListArgs:
		success, err := handler.(product_attribute_value.ProductAttributeValueService).ProductAttributeValueList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductAttributeValueListResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductAttributeValueListArgs() interface{} {
	return &ProductAttributeValueListArgs{}
}

func newProductAttributeValueListResult() interface{} {
	return &ProductAttributeValueListResult{}
}

type ProductAttributeValueListArgs struct {
	Req *product_attribute_value.ProductAttributeValueListReq
}

func (p *ProductAttributeValueListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_attribute_value.ProductAttributeValueListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductAttributeValueListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductAttributeValueListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductAttributeValueListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductAttributeValueListArgs) Unmarshal(in []byte) error {
	msg := new(product_attribute_value.ProductAttributeValueListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductAttributeValueListArgs_Req_DEFAULT *product_attribute_value.ProductAttributeValueListReq

func (p *ProductAttributeValueListArgs) GetReq() *product_attribute_value.ProductAttributeValueListReq {
	if !p.IsSetReq() {
		return ProductAttributeValueListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductAttributeValueListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductAttributeValueListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductAttributeValueListResult struct {
	Success *product_attribute_value.ProductAttributeValueListResp
}

var ProductAttributeValueListResult_Success_DEFAULT *product_attribute_value.ProductAttributeValueListResp

func (p *ProductAttributeValueListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_attribute_value.ProductAttributeValueListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductAttributeValueListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductAttributeValueListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductAttributeValueListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductAttributeValueListResult) Unmarshal(in []byte) error {
	msg := new(product_attribute_value.ProductAttributeValueListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductAttributeValueListResult) GetSuccess() *product_attribute_value.ProductAttributeValueListResp {
	if !p.IsSetSuccess() {
		return ProductAttributeValueListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductAttributeValueListResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_attribute_value.ProductAttributeValueListResp)
}

func (p *ProductAttributeValueListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductAttributeValueListResult) GetResult() interface{} {
	return p.Success
}

func productAttributeValueUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_attribute_value.ProductAttributeValueUpdateReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_attribute_value.ProductAttributeValueService).ProductAttributeValueUpdate(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductAttributeValueUpdateArgs:
		success, err := handler.(product_attribute_value.ProductAttributeValueService).ProductAttributeValueUpdate(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductAttributeValueUpdateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductAttributeValueUpdateArgs() interface{} {
	return &ProductAttributeValueUpdateArgs{}
}

func newProductAttributeValueUpdateResult() interface{} {
	return &ProductAttributeValueUpdateResult{}
}

type ProductAttributeValueUpdateArgs struct {
	Req *product_attribute_value.ProductAttributeValueUpdateReq
}

func (p *ProductAttributeValueUpdateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_attribute_value.ProductAttributeValueUpdateReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductAttributeValueUpdateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductAttributeValueUpdateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductAttributeValueUpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductAttributeValueUpdateArgs) Unmarshal(in []byte) error {
	msg := new(product_attribute_value.ProductAttributeValueUpdateReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductAttributeValueUpdateArgs_Req_DEFAULT *product_attribute_value.ProductAttributeValueUpdateReq

func (p *ProductAttributeValueUpdateArgs) GetReq() *product_attribute_value.ProductAttributeValueUpdateReq {
	if !p.IsSetReq() {
		return ProductAttributeValueUpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductAttributeValueUpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductAttributeValueUpdateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductAttributeValueUpdateResult struct {
	Success *product_attribute_value.ProductAttributeValueUpdateResp
}

var ProductAttributeValueUpdateResult_Success_DEFAULT *product_attribute_value.ProductAttributeValueUpdateResp

func (p *ProductAttributeValueUpdateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_attribute_value.ProductAttributeValueUpdateResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductAttributeValueUpdateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductAttributeValueUpdateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductAttributeValueUpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductAttributeValueUpdateResult) Unmarshal(in []byte) error {
	msg := new(product_attribute_value.ProductAttributeValueUpdateResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductAttributeValueUpdateResult) GetSuccess() *product_attribute_value.ProductAttributeValueUpdateResp {
	if !p.IsSetSuccess() {
		return ProductAttributeValueUpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductAttributeValueUpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_attribute_value.ProductAttributeValueUpdateResp)
}

func (p *ProductAttributeValueUpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductAttributeValueUpdateResult) GetResult() interface{} {
	return p.Success
}

func productAttributeValueDeleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_attribute_value.ProductAttributeValueDeleteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_attribute_value.ProductAttributeValueService).ProductAttributeValueDelete(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductAttributeValueDeleteArgs:
		success, err := handler.(product_attribute_value.ProductAttributeValueService).ProductAttributeValueDelete(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductAttributeValueDeleteResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductAttributeValueDeleteArgs() interface{} {
	return &ProductAttributeValueDeleteArgs{}
}

func newProductAttributeValueDeleteResult() interface{} {
	return &ProductAttributeValueDeleteResult{}
}

type ProductAttributeValueDeleteArgs struct {
	Req *product_attribute_value.ProductAttributeValueDeleteReq
}

func (p *ProductAttributeValueDeleteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_attribute_value.ProductAttributeValueDeleteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductAttributeValueDeleteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductAttributeValueDeleteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductAttributeValueDeleteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductAttributeValueDeleteArgs) Unmarshal(in []byte) error {
	msg := new(product_attribute_value.ProductAttributeValueDeleteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductAttributeValueDeleteArgs_Req_DEFAULT *product_attribute_value.ProductAttributeValueDeleteReq

func (p *ProductAttributeValueDeleteArgs) GetReq() *product_attribute_value.ProductAttributeValueDeleteReq {
	if !p.IsSetReq() {
		return ProductAttributeValueDeleteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductAttributeValueDeleteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductAttributeValueDeleteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductAttributeValueDeleteResult struct {
	Success *product_attribute_value.ProductAttributeValueDeleteResp
}

var ProductAttributeValueDeleteResult_Success_DEFAULT *product_attribute_value.ProductAttributeValueDeleteResp

func (p *ProductAttributeValueDeleteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_attribute_value.ProductAttributeValueDeleteResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductAttributeValueDeleteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductAttributeValueDeleteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductAttributeValueDeleteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductAttributeValueDeleteResult) Unmarshal(in []byte) error {
	msg := new(product_attribute_value.ProductAttributeValueDeleteResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductAttributeValueDeleteResult) GetSuccess() *product_attribute_value.ProductAttributeValueDeleteResp {
	if !p.IsSetSuccess() {
		return ProductAttributeValueDeleteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductAttributeValueDeleteResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_attribute_value.ProductAttributeValueDeleteResp)
}

func (p *ProductAttributeValueDeleteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductAttributeValueDeleteResult) GetResult() interface{} {
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

func (p *kClient) ProductAttributeValueAdd(ctx context.Context, Req *product_attribute_value.ProductAttributeValueAddReq) (r *product_attribute_value.ProductAttributeValueAddResp, err error) {
	var _args ProductAttributeValueAddArgs
	_args.Req = Req
	var _result ProductAttributeValueAddResult
	if err = p.c.Call(ctx, "ProductAttributeValueAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductAttributeValueList(ctx context.Context, Req *product_attribute_value.ProductAttributeValueListReq) (r *product_attribute_value.ProductAttributeValueListResp, err error) {
	var _args ProductAttributeValueListArgs
	_args.Req = Req
	var _result ProductAttributeValueListResult
	if err = p.c.Call(ctx, "ProductAttributeValueList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductAttributeValueUpdate(ctx context.Context, Req *product_attribute_value.ProductAttributeValueUpdateReq) (r *product_attribute_value.ProductAttributeValueUpdateResp, err error) {
	var _args ProductAttributeValueUpdateArgs
	_args.Req = Req
	var _result ProductAttributeValueUpdateResult
	if err = p.c.Call(ctx, "ProductAttributeValueUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductAttributeValueDelete(ctx context.Context, Req *product_attribute_value.ProductAttributeValueDeleteReq) (r *product_attribute_value.ProductAttributeValueDeleteResp, err error) {
	var _args ProductAttributeValueDeleteArgs
	_args.Req = Req
	var _result ProductAttributeValueDeleteResult
	if err = p.c.Call(ctx, "ProductAttributeValueDelete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
