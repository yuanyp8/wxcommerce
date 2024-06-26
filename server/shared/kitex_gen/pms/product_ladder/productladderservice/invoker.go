// Code generated by Kitex v0.9.0. DO NOT EDIT.

package productladderservice

import (
	server "github.com/cloudwego/kitex/server"
	product_ladder "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_ladder"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler product_ladder.ProductLadderService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
