// Code generated by Kitex v0.9.0. DO NOT EDIT.

package memberpriceservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	member_price "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/member_price"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MemberPriceAdd(ctx context.Context, Req *member_price.MemberPriceAddReq, callOptions ...callopt.Option) (r *member_price.MemberPriceAddResp, err error)
	MemberPriceList(ctx context.Context, Req *member_price.MemberPriceListReq, callOptions ...callopt.Option) (r *member_price.MemberPriceListResp, err error)
	MemberPriceUpdate(ctx context.Context, Req *member_price.MemberPriceUpdateReq, callOptions ...callopt.Option) (r *member_price.MemberPriceUpdateResp, err error)
	MemberPriceDelete(ctx context.Context, Req *member_price.MemberPriceDeleteReq, callOptions ...callopt.Option) (r *member_price.MemberPriceDeleteResp, err error)
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
	return &kMemberPriceServiceClient{
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

type kMemberPriceServiceClient struct {
	*kClient
}

func (p *kMemberPriceServiceClient) MemberPriceAdd(ctx context.Context, Req *member_price.MemberPriceAddReq, callOptions ...callopt.Option) (r *member_price.MemberPriceAddResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MemberPriceAdd(ctx, Req)
}

func (p *kMemberPriceServiceClient) MemberPriceList(ctx context.Context, Req *member_price.MemberPriceListReq, callOptions ...callopt.Option) (r *member_price.MemberPriceListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MemberPriceList(ctx, Req)
}

func (p *kMemberPriceServiceClient) MemberPriceUpdate(ctx context.Context, Req *member_price.MemberPriceUpdateReq, callOptions ...callopt.Option) (r *member_price.MemberPriceUpdateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MemberPriceUpdate(ctx, Req)
}

func (p *kMemberPriceServiceClient) MemberPriceDelete(ctx context.Context, Req *member_price.MemberPriceDeleteReq, callOptions ...callopt.Option) (r *member_price.MemberPriceDeleteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MemberPriceDelete(ctx, Req)
}
