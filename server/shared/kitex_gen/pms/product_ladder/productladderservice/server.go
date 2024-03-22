// Code generated by Kitex v0.9.0. DO NOT EDIT.
package productladderservice

import (
	server "github.com/cloudwego/kitex/server"
	product_ladder "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_ladder"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler product_ladder.ProductLadderService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler product_ladder.ProductLadderService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
