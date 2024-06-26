// Code generated by Kitex v0.9.0. DO NOT EDIT.

package productattributecategoryservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	product_attribute_category "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_attribute_category"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ProductAttributeCategoryAdd": kitex.NewMethodInfo(
		productAttributeCategoryAddHandler,
		newProductAttributeCategoryAddArgs,
		newProductAttributeCategoryAddResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductAttributeCategoryList": kitex.NewMethodInfo(
		productAttributeCategoryListHandler,
		newProductAttributeCategoryListArgs,
		newProductAttributeCategoryListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductAttributeCategoryUpdate": kitex.NewMethodInfo(
		productAttributeCategoryUpdateHandler,
		newProductAttributeCategoryUpdateArgs,
		newProductAttributeCategoryUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ProductAttributeCategoryDelete": kitex.NewMethodInfo(
		productAttributeCategoryDeleteHandler,
		newProductAttributeCategoryDeleteArgs,
		newProductAttributeCategoryDeleteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	productAttributeCategoryServiceServiceInfo                = NewServiceInfo()
	productAttributeCategoryServiceServiceInfoForClient       = NewServiceInfoForClient()
	productAttributeCategoryServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return productAttributeCategoryServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return productAttributeCategoryServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return productAttributeCategoryServiceServiceInfoForClient
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
	serviceName := "ProductAttributeCategoryService"
	handlerType := (*product_attribute_category.ProductAttributeCategoryService)(nil)
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

func productAttributeCategoryAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_attribute_category.ProductAttributeCategoryAddReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_attribute_category.ProductAttributeCategoryService).ProductAttributeCategoryAdd(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductAttributeCategoryAddArgs:
		success, err := handler.(product_attribute_category.ProductAttributeCategoryService).ProductAttributeCategoryAdd(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductAttributeCategoryAddResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductAttributeCategoryAddArgs() interface{} {
	return &ProductAttributeCategoryAddArgs{}
}

func newProductAttributeCategoryAddResult() interface{} {
	return &ProductAttributeCategoryAddResult{}
}

type ProductAttributeCategoryAddArgs struct {
	Req *product_attribute_category.ProductAttributeCategoryAddReq
}

func (p *ProductAttributeCategoryAddArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_attribute_category.ProductAttributeCategoryAddReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductAttributeCategoryAddArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductAttributeCategoryAddArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductAttributeCategoryAddArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductAttributeCategoryAddArgs) Unmarshal(in []byte) error {
	msg := new(product_attribute_category.ProductAttributeCategoryAddReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductAttributeCategoryAddArgs_Req_DEFAULT *product_attribute_category.ProductAttributeCategoryAddReq

func (p *ProductAttributeCategoryAddArgs) GetReq() *product_attribute_category.ProductAttributeCategoryAddReq {
	if !p.IsSetReq() {
		return ProductAttributeCategoryAddArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductAttributeCategoryAddArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductAttributeCategoryAddArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductAttributeCategoryAddResult struct {
	Success *product_attribute_category.ProductAttributeCategoryAddResp
}

var ProductAttributeCategoryAddResult_Success_DEFAULT *product_attribute_category.ProductAttributeCategoryAddResp

func (p *ProductAttributeCategoryAddResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_attribute_category.ProductAttributeCategoryAddResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductAttributeCategoryAddResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductAttributeCategoryAddResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductAttributeCategoryAddResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductAttributeCategoryAddResult) Unmarshal(in []byte) error {
	msg := new(product_attribute_category.ProductAttributeCategoryAddResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductAttributeCategoryAddResult) GetSuccess() *product_attribute_category.ProductAttributeCategoryAddResp {
	if !p.IsSetSuccess() {
		return ProductAttributeCategoryAddResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductAttributeCategoryAddResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_attribute_category.ProductAttributeCategoryAddResp)
}

func (p *ProductAttributeCategoryAddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductAttributeCategoryAddResult) GetResult() interface{} {
	return p.Success
}

func productAttributeCategoryListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_attribute_category.ProductAttributeCategoryListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_attribute_category.ProductAttributeCategoryService).ProductAttributeCategoryList(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductAttributeCategoryListArgs:
		success, err := handler.(product_attribute_category.ProductAttributeCategoryService).ProductAttributeCategoryList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductAttributeCategoryListResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductAttributeCategoryListArgs() interface{} {
	return &ProductAttributeCategoryListArgs{}
}

func newProductAttributeCategoryListResult() interface{} {
	return &ProductAttributeCategoryListResult{}
}

type ProductAttributeCategoryListArgs struct {
	Req *product_attribute_category.ProductAttributeCategoryListReq
}

func (p *ProductAttributeCategoryListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_attribute_category.ProductAttributeCategoryListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductAttributeCategoryListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductAttributeCategoryListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductAttributeCategoryListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductAttributeCategoryListArgs) Unmarshal(in []byte) error {
	msg := new(product_attribute_category.ProductAttributeCategoryListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductAttributeCategoryListArgs_Req_DEFAULT *product_attribute_category.ProductAttributeCategoryListReq

func (p *ProductAttributeCategoryListArgs) GetReq() *product_attribute_category.ProductAttributeCategoryListReq {
	if !p.IsSetReq() {
		return ProductAttributeCategoryListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductAttributeCategoryListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductAttributeCategoryListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductAttributeCategoryListResult struct {
	Success *product_attribute_category.ProductAttributeCategoryListResp
}

var ProductAttributeCategoryListResult_Success_DEFAULT *product_attribute_category.ProductAttributeCategoryListResp

func (p *ProductAttributeCategoryListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_attribute_category.ProductAttributeCategoryListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductAttributeCategoryListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductAttributeCategoryListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductAttributeCategoryListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductAttributeCategoryListResult) Unmarshal(in []byte) error {
	msg := new(product_attribute_category.ProductAttributeCategoryListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductAttributeCategoryListResult) GetSuccess() *product_attribute_category.ProductAttributeCategoryListResp {
	if !p.IsSetSuccess() {
		return ProductAttributeCategoryListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductAttributeCategoryListResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_attribute_category.ProductAttributeCategoryListResp)
}

func (p *ProductAttributeCategoryListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductAttributeCategoryListResult) GetResult() interface{} {
	return p.Success
}

func productAttributeCategoryUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_attribute_category.ProductAttributeCategoryUpdateReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_attribute_category.ProductAttributeCategoryService).ProductAttributeCategoryUpdate(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductAttributeCategoryUpdateArgs:
		success, err := handler.(product_attribute_category.ProductAttributeCategoryService).ProductAttributeCategoryUpdate(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductAttributeCategoryUpdateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductAttributeCategoryUpdateArgs() interface{} {
	return &ProductAttributeCategoryUpdateArgs{}
}

func newProductAttributeCategoryUpdateResult() interface{} {
	return &ProductAttributeCategoryUpdateResult{}
}

type ProductAttributeCategoryUpdateArgs struct {
	Req *product_attribute_category.ProductAttributeCategoryUpdateReq
}

func (p *ProductAttributeCategoryUpdateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_attribute_category.ProductAttributeCategoryUpdateReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductAttributeCategoryUpdateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductAttributeCategoryUpdateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductAttributeCategoryUpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductAttributeCategoryUpdateArgs) Unmarshal(in []byte) error {
	msg := new(product_attribute_category.ProductAttributeCategoryUpdateReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductAttributeCategoryUpdateArgs_Req_DEFAULT *product_attribute_category.ProductAttributeCategoryUpdateReq

func (p *ProductAttributeCategoryUpdateArgs) GetReq() *product_attribute_category.ProductAttributeCategoryUpdateReq {
	if !p.IsSetReq() {
		return ProductAttributeCategoryUpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductAttributeCategoryUpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductAttributeCategoryUpdateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductAttributeCategoryUpdateResult struct {
	Success *product_attribute_category.ProductAttributeCategoryUpdateResp
}

var ProductAttributeCategoryUpdateResult_Success_DEFAULT *product_attribute_category.ProductAttributeCategoryUpdateResp

func (p *ProductAttributeCategoryUpdateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_attribute_category.ProductAttributeCategoryUpdateResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductAttributeCategoryUpdateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductAttributeCategoryUpdateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductAttributeCategoryUpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductAttributeCategoryUpdateResult) Unmarshal(in []byte) error {
	msg := new(product_attribute_category.ProductAttributeCategoryUpdateResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductAttributeCategoryUpdateResult) GetSuccess() *product_attribute_category.ProductAttributeCategoryUpdateResp {
	if !p.IsSetSuccess() {
		return ProductAttributeCategoryUpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductAttributeCategoryUpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_attribute_category.ProductAttributeCategoryUpdateResp)
}

func (p *ProductAttributeCategoryUpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductAttributeCategoryUpdateResult) GetResult() interface{} {
	return p.Success
}

func productAttributeCategoryDeleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product_attribute_category.ProductAttributeCategoryDeleteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product_attribute_category.ProductAttributeCategoryService).ProductAttributeCategoryDelete(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ProductAttributeCategoryDeleteArgs:
		success, err := handler.(product_attribute_category.ProductAttributeCategoryService).ProductAttributeCategoryDelete(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ProductAttributeCategoryDeleteResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newProductAttributeCategoryDeleteArgs() interface{} {
	return &ProductAttributeCategoryDeleteArgs{}
}

func newProductAttributeCategoryDeleteResult() interface{} {
	return &ProductAttributeCategoryDeleteResult{}
}

type ProductAttributeCategoryDeleteArgs struct {
	Req *product_attribute_category.ProductAttributeCategoryDeleteReq
}

func (p *ProductAttributeCategoryDeleteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product_attribute_category.ProductAttributeCategoryDeleteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ProductAttributeCategoryDeleteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ProductAttributeCategoryDeleteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ProductAttributeCategoryDeleteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ProductAttributeCategoryDeleteArgs) Unmarshal(in []byte) error {
	msg := new(product_attribute_category.ProductAttributeCategoryDeleteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ProductAttributeCategoryDeleteArgs_Req_DEFAULT *product_attribute_category.ProductAttributeCategoryDeleteReq

func (p *ProductAttributeCategoryDeleteArgs) GetReq() *product_attribute_category.ProductAttributeCategoryDeleteReq {
	if !p.IsSetReq() {
		return ProductAttributeCategoryDeleteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ProductAttributeCategoryDeleteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ProductAttributeCategoryDeleteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ProductAttributeCategoryDeleteResult struct {
	Success *product_attribute_category.ProductAttributeCategoryDeleteResp
}

var ProductAttributeCategoryDeleteResult_Success_DEFAULT *product_attribute_category.ProductAttributeCategoryDeleteResp

func (p *ProductAttributeCategoryDeleteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product_attribute_category.ProductAttributeCategoryDeleteResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ProductAttributeCategoryDeleteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ProductAttributeCategoryDeleteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ProductAttributeCategoryDeleteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ProductAttributeCategoryDeleteResult) Unmarshal(in []byte) error {
	msg := new(product_attribute_category.ProductAttributeCategoryDeleteResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ProductAttributeCategoryDeleteResult) GetSuccess() *product_attribute_category.ProductAttributeCategoryDeleteResp {
	if !p.IsSetSuccess() {
		return ProductAttributeCategoryDeleteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ProductAttributeCategoryDeleteResult) SetSuccess(x interface{}) {
	p.Success = x.(*product_attribute_category.ProductAttributeCategoryDeleteResp)
}

func (p *ProductAttributeCategoryDeleteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ProductAttributeCategoryDeleteResult) GetResult() interface{} {
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

func (p *kClient) ProductAttributeCategoryAdd(ctx context.Context, Req *product_attribute_category.ProductAttributeCategoryAddReq) (r *product_attribute_category.ProductAttributeCategoryAddResp, err error) {
	var _args ProductAttributeCategoryAddArgs
	_args.Req = Req
	var _result ProductAttributeCategoryAddResult
	if err = p.c.Call(ctx, "ProductAttributeCategoryAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductAttributeCategoryList(ctx context.Context, Req *product_attribute_category.ProductAttributeCategoryListReq) (r *product_attribute_category.ProductAttributeCategoryListResp, err error) {
	var _args ProductAttributeCategoryListArgs
	_args.Req = Req
	var _result ProductAttributeCategoryListResult
	if err = p.c.Call(ctx, "ProductAttributeCategoryList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductAttributeCategoryUpdate(ctx context.Context, Req *product_attribute_category.ProductAttributeCategoryUpdateReq) (r *product_attribute_category.ProductAttributeCategoryUpdateResp, err error) {
	var _args ProductAttributeCategoryUpdateArgs
	_args.Req = Req
	var _result ProductAttributeCategoryUpdateResult
	if err = p.c.Call(ctx, "ProductAttributeCategoryUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ProductAttributeCategoryDelete(ctx context.Context, Req *product_attribute_category.ProductAttributeCategoryDeleteReq) (r *product_attribute_category.ProductAttributeCategoryDeleteResp, err error) {
	var _args ProductAttributeCategoryDeleteArgs
	_args.Req = Req
	var _result ProductAttributeCategoryDeleteResult
	if err = p.c.Call(ctx, "ProductAttributeCategoryDelete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
