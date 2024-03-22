// Code generated by Kitex v0.9.0. DO NOT EDIT.

package albumservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	album "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/album"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"AlbumAdd": kitex.NewMethodInfo(
		albumAddHandler,
		newAlbumAddArgs,
		newAlbumAddResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AlbumList": kitex.NewMethodInfo(
		albumListHandler,
		newAlbumListArgs,
		newAlbumListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AlbumUpdate": kitex.NewMethodInfo(
		albumUpdateHandler,
		newAlbumUpdateArgs,
		newAlbumUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AlbumDelete": kitex.NewMethodInfo(
		albumDeleteHandler,
		newAlbumDeleteArgs,
		newAlbumDeleteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	albumServiceServiceInfo                = NewServiceInfo()
	albumServiceServiceInfoForClient       = NewServiceInfoForClient()
	albumServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return albumServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return albumServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return albumServiceServiceInfoForClient
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
	serviceName := "AlbumService"
	handlerType := (*album.AlbumService)(nil)
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

func albumAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(album.AlbumAddReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(album.AlbumService).AlbumAdd(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AlbumAddArgs:
		success, err := handler.(album.AlbumService).AlbumAdd(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AlbumAddResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAlbumAddArgs() interface{} {
	return &AlbumAddArgs{}
}

func newAlbumAddResult() interface{} {
	return &AlbumAddResult{}
}

type AlbumAddArgs struct {
	Req *album.AlbumAddReq
}

func (p *AlbumAddArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(album.AlbumAddReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AlbumAddArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AlbumAddArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AlbumAddArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AlbumAddArgs) Unmarshal(in []byte) error {
	msg := new(album.AlbumAddReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AlbumAddArgs_Req_DEFAULT *album.AlbumAddReq

func (p *AlbumAddArgs) GetReq() *album.AlbumAddReq {
	if !p.IsSetReq() {
		return AlbumAddArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AlbumAddArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AlbumAddArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AlbumAddResult struct {
	Success *album.AlbumAddResp
}

var AlbumAddResult_Success_DEFAULT *album.AlbumAddResp

func (p *AlbumAddResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(album.AlbumAddResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AlbumAddResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AlbumAddResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AlbumAddResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AlbumAddResult) Unmarshal(in []byte) error {
	msg := new(album.AlbumAddResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AlbumAddResult) GetSuccess() *album.AlbumAddResp {
	if !p.IsSetSuccess() {
		return AlbumAddResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AlbumAddResult) SetSuccess(x interface{}) {
	p.Success = x.(*album.AlbumAddResp)
}

func (p *AlbumAddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AlbumAddResult) GetResult() interface{} {
	return p.Success
}

func albumListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(album.AlbumListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(album.AlbumService).AlbumList(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AlbumListArgs:
		success, err := handler.(album.AlbumService).AlbumList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AlbumListResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAlbumListArgs() interface{} {
	return &AlbumListArgs{}
}

func newAlbumListResult() interface{} {
	return &AlbumListResult{}
}

type AlbumListArgs struct {
	Req *album.AlbumListReq
}

func (p *AlbumListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(album.AlbumListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AlbumListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AlbumListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AlbumListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AlbumListArgs) Unmarshal(in []byte) error {
	msg := new(album.AlbumListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AlbumListArgs_Req_DEFAULT *album.AlbumListReq

func (p *AlbumListArgs) GetReq() *album.AlbumListReq {
	if !p.IsSetReq() {
		return AlbumListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AlbumListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AlbumListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AlbumListResult struct {
	Success *album.AlbumListResp
}

var AlbumListResult_Success_DEFAULT *album.AlbumListResp

func (p *AlbumListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(album.AlbumListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AlbumListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AlbumListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AlbumListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AlbumListResult) Unmarshal(in []byte) error {
	msg := new(album.AlbumListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AlbumListResult) GetSuccess() *album.AlbumListResp {
	if !p.IsSetSuccess() {
		return AlbumListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AlbumListResult) SetSuccess(x interface{}) {
	p.Success = x.(*album.AlbumListResp)
}

func (p *AlbumListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AlbumListResult) GetResult() interface{} {
	return p.Success
}

func albumUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(album.AlbumUpdateReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(album.AlbumService).AlbumUpdate(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AlbumUpdateArgs:
		success, err := handler.(album.AlbumService).AlbumUpdate(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AlbumUpdateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAlbumUpdateArgs() interface{} {
	return &AlbumUpdateArgs{}
}

func newAlbumUpdateResult() interface{} {
	return &AlbumUpdateResult{}
}

type AlbumUpdateArgs struct {
	Req *album.AlbumUpdateReq
}

func (p *AlbumUpdateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(album.AlbumUpdateReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AlbumUpdateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AlbumUpdateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AlbumUpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AlbumUpdateArgs) Unmarshal(in []byte) error {
	msg := new(album.AlbumUpdateReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AlbumUpdateArgs_Req_DEFAULT *album.AlbumUpdateReq

func (p *AlbumUpdateArgs) GetReq() *album.AlbumUpdateReq {
	if !p.IsSetReq() {
		return AlbumUpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AlbumUpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AlbumUpdateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AlbumUpdateResult struct {
	Success *album.AlbumUpdateResp
}

var AlbumUpdateResult_Success_DEFAULT *album.AlbumUpdateResp

func (p *AlbumUpdateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(album.AlbumUpdateResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AlbumUpdateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AlbumUpdateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AlbumUpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AlbumUpdateResult) Unmarshal(in []byte) error {
	msg := new(album.AlbumUpdateResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AlbumUpdateResult) GetSuccess() *album.AlbumUpdateResp {
	if !p.IsSetSuccess() {
		return AlbumUpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AlbumUpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*album.AlbumUpdateResp)
}

func (p *AlbumUpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AlbumUpdateResult) GetResult() interface{} {
	return p.Success
}

func albumDeleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(album.AlbumDeleteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(album.AlbumService).AlbumDelete(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AlbumDeleteArgs:
		success, err := handler.(album.AlbumService).AlbumDelete(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AlbumDeleteResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAlbumDeleteArgs() interface{} {
	return &AlbumDeleteArgs{}
}

func newAlbumDeleteResult() interface{} {
	return &AlbumDeleteResult{}
}

type AlbumDeleteArgs struct {
	Req *album.AlbumDeleteReq
}

func (p *AlbumDeleteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(album.AlbumDeleteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AlbumDeleteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AlbumDeleteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AlbumDeleteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AlbumDeleteArgs) Unmarshal(in []byte) error {
	msg := new(album.AlbumDeleteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AlbumDeleteArgs_Req_DEFAULT *album.AlbumDeleteReq

func (p *AlbumDeleteArgs) GetReq() *album.AlbumDeleteReq {
	if !p.IsSetReq() {
		return AlbumDeleteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AlbumDeleteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AlbumDeleteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AlbumDeleteResult struct {
	Success *album.AlbumDeleteResp
}

var AlbumDeleteResult_Success_DEFAULT *album.AlbumDeleteResp

func (p *AlbumDeleteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(album.AlbumDeleteResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AlbumDeleteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AlbumDeleteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AlbumDeleteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AlbumDeleteResult) Unmarshal(in []byte) error {
	msg := new(album.AlbumDeleteResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AlbumDeleteResult) GetSuccess() *album.AlbumDeleteResp {
	if !p.IsSetSuccess() {
		return AlbumDeleteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AlbumDeleteResult) SetSuccess(x interface{}) {
	p.Success = x.(*album.AlbumDeleteResp)
}

func (p *AlbumDeleteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AlbumDeleteResult) GetResult() interface{} {
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

func (p *kClient) AlbumAdd(ctx context.Context, Req *album.AlbumAddReq) (r *album.AlbumAddResp, err error) {
	var _args AlbumAddArgs
	_args.Req = Req
	var _result AlbumAddResult
	if err = p.c.Call(ctx, "AlbumAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AlbumList(ctx context.Context, Req *album.AlbumListReq) (r *album.AlbumListResp, err error) {
	var _args AlbumListArgs
	_args.Req = Req
	var _result AlbumListResult
	if err = p.c.Call(ctx, "AlbumList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AlbumUpdate(ctx context.Context, Req *album.AlbumUpdateReq) (r *album.AlbumUpdateResp, err error) {
	var _args AlbumUpdateArgs
	_args.Req = Req
	var _result AlbumUpdateResult
	if err = p.c.Call(ctx, "AlbumUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AlbumDelete(ctx context.Context, Req *album.AlbumDeleteReq) (r *album.AlbumDeleteResp, err error) {
	var _args AlbumDeleteArgs
	_args.Req = Req
	var _result AlbumDeleteResult
	if err = p.c.Call(ctx, "AlbumDelete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
