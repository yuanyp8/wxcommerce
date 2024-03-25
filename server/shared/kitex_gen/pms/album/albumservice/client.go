// Code generated by Kitex v0.9.0. DO NOT EDIT.

package albumservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	album "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/album"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AlbumAdd(ctx context.Context, Req *album.AlbumAddReq, callOptions ...callopt.Option) (r *album.AlbumAddResp, err error)
	AlbumList(ctx context.Context, Req *album.AlbumListReq, callOptions ...callopt.Option) (r *album.AlbumListResp, err error)
	AlbumUpdate(ctx context.Context, Req *album.AlbumUpdateReq, callOptions ...callopt.Option) (r *album.AlbumUpdateResp, err error)
	AlbumMDelete(ctx context.Context, Req *album.AlbumDeleteReq, callOptions ...callopt.Option) (r *album.AlbumDeleteResp, err error)
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
	return &kAlbumServiceClient{
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

type kAlbumServiceClient struct {
	*kClient
}

func (p *kAlbumServiceClient) AlbumAdd(ctx context.Context, Req *album.AlbumAddReq, callOptions ...callopt.Option) (r *album.AlbumAddResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AlbumAdd(ctx, Req)
}

func (p *kAlbumServiceClient) AlbumList(ctx context.Context, Req *album.AlbumListReq, callOptions ...callopt.Option) (r *album.AlbumListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AlbumList(ctx, Req)
}

func (p *kAlbumServiceClient) AlbumUpdate(ctx context.Context, Req *album.AlbumUpdateReq, callOptions ...callopt.Option) (r *album.AlbumUpdateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AlbumUpdate(ctx, Req)
}

func (p *kAlbumServiceClient) AlbumMDelete(ctx context.Context, Req *album.AlbumDeleteReq, callOptions ...callopt.Option) (r *album.AlbumDeleteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AlbumMDelete(ctx, Req)
}
