// Code generated by Kitex v0.9.0. DO NOT EDIT.

package albumpicservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	album_pic "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/album_pic"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AlbumPicAdd(ctx context.Context, Req *album_pic.AlbumPicAddReq, callOptions ...callopt.Option) (r *album_pic.AlbumPicAddResp, err error)
	AlbumPicList(ctx context.Context, Req *album_pic.AlbumPicListReq, callOptions ...callopt.Option) (r *album_pic.AlbumPicListResp, err error)
	AlbumPicUpdate(ctx context.Context, Req *album_pic.AlbumPicUpdateReq, callOptions ...callopt.Option) (r *album_pic.AlbumPicUpdateResp, err error)
	AlbumPicDelete(ctx context.Context, Req *album_pic.AlbumPicDeleteReq, callOptions ...callopt.Option) (r *album_pic.AlbumPicDeleteResp, err error)
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
	return &kAlbumPicServiceClient{
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

type kAlbumPicServiceClient struct {
	*kClient
}

func (p *kAlbumPicServiceClient) AlbumPicAdd(ctx context.Context, Req *album_pic.AlbumPicAddReq, callOptions ...callopt.Option) (r *album_pic.AlbumPicAddResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AlbumPicAdd(ctx, Req)
}

func (p *kAlbumPicServiceClient) AlbumPicList(ctx context.Context, Req *album_pic.AlbumPicListReq, callOptions ...callopt.Option) (r *album_pic.AlbumPicListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AlbumPicList(ctx, Req)
}

func (p *kAlbumPicServiceClient) AlbumPicUpdate(ctx context.Context, Req *album_pic.AlbumPicUpdateReq, callOptions ...callopt.Option) (r *album_pic.AlbumPicUpdateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AlbumPicUpdate(ctx, Req)
}

func (p *kAlbumPicServiceClient) AlbumPicDelete(ctx context.Context, Req *album_pic.AlbumPicDeleteReq, callOptions ...callopt.Option) (r *album_pic.AlbumPicDeleteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AlbumPicDelete(ctx, Req)
}
