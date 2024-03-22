// Code generated by Kitex v0.9.0. DO NOT EDIT.

package productoperatelogservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	product_full_reduction "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_full_reduction"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ProductOperateLogAdd(ctx context.Context, Req *product_full_reduction.ProductOperateLogAddReq, callOptions ...callopt.Option) (r *product_full_reduction.ProductOperateLogAddResp, err error)
	ProductOperateLogList(ctx context.Context, Req *product_full_reduction.ProductOperateLogListReq, callOptions ...callopt.Option) (r *product_full_reduction.ProductOperateLogListResp, err error)
	ProductOperateLogUpdate(ctx context.Context, Req *product_full_reduction.ProductOperateLogUpdateReq, callOptions ...callopt.Option) (r *product_full_reduction.ProductOperateLogUpdateResp, err error)
	ProductOperateLogDelete(ctx context.Context, Req *product_full_reduction.ProductOperateLogDeleteReq, callOptions ...callopt.Option) (r *product_full_reduction.ProductOperateLogDeleteResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kProductOperateLogServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kProductOperateLogServiceClient struct {
	*kClient
}

func (p *kProductOperateLogServiceClient) ProductOperateLogAdd(ctx context.Context, Req *product_full_reduction.ProductOperateLogAddReq, callOptions ...callopt.Option) (r *product_full_reduction.ProductOperateLogAddResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ProductOperateLogAdd(ctx, Req)
}

func (p *kProductOperateLogServiceClient) ProductOperateLogList(ctx context.Context, Req *product_full_reduction.ProductOperateLogListReq, callOptions ...callopt.Option) (r *product_full_reduction.ProductOperateLogListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ProductOperateLogList(ctx, Req)
}

func (p *kProductOperateLogServiceClient) ProductOperateLogUpdate(ctx context.Context, Req *product_full_reduction.ProductOperateLogUpdateReq, callOptions ...callopt.Option) (r *product_full_reduction.ProductOperateLogUpdateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ProductOperateLogUpdate(ctx, Req)
}

func (p *kProductOperateLogServiceClient) ProductOperateLogDelete(ctx context.Context, Req *product_full_reduction.ProductOperateLogDeleteReq, callOptions ...callopt.Option) (r *product_full_reduction.ProductOperateLogDeleteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ProductOperateLogDelete(ctx, Req)
}
