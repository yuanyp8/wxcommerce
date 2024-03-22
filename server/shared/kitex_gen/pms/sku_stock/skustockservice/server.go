// Code generated by Kitex v0.9.0. DO NOT EDIT.
package skustockservice

import (
	server "github.com/cloudwego/kitex/server"
	sku_stock "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/sku_stock"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler sku_stock.SkuStockService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler sku_stock.SkuStockService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}