// Code generated by Kitex v0.9.0. DO NOT EDIT.
package productcategoryattributerelationservice

import (
	server "github.com/cloudwego/kitex/server"
	product_category_attribute_relation "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_category_attribute_relation"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler product_category_attribute_relation.ProductCategoryAttributeRelationService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler product_category_attribute_relation.ProductCategoryAttributeRelationService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
