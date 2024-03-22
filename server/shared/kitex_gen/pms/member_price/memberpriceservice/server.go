// Code generated by Kitex v0.9.0. DO NOT EDIT.
package memberpriceservice

import (
	server "github.com/cloudwego/kitex/server"
	member_price "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/member_price"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler member_price.MemberPriceService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler member_price.MemberPriceService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
