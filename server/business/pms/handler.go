package main

import (
	"context"
	album "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/album"
	sku_stock "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/sku_stock"
	product_verify_record "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_verify_record"
	product "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product"
	product_ladder "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_ladder"
	product_full_reduction "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_full_reduction"
	product_category_attribute_relation "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_category_attribute_relation"
	product_category "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_category"
	product_attribute_value "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_attribute_value"
	product_attribute_category "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_attribute_category"
	product_attribute "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/product_attribute"
	member_price "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/member_price"
	comment_reply "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/comment_reply"
	comment "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/comment"
	brand "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/brand"
	album_pic "github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/album_pic"
)

// AlbumServiceImpl implements the last service interface defined in the IDL.
type AlbumServiceImpl struct{}

// AlbumAdd implements the AlbumServiceImpl interface.
func (s *AlbumServiceImpl) AlbumAdd(ctx context.Context, req *album.AlbumAddReq) (resp *album.AlbumAddResp, err error) {
	// TODO: Your code here...
	return
}

// AlbumList implements the AlbumServiceImpl interface.
func (s *AlbumServiceImpl) AlbumList(ctx context.Context, req *album.AlbumListReq) (resp *album.AlbumListResp, err error) {
	// TODO: Your code here...
	return
}

// AlbumUpdate implements the AlbumServiceImpl interface.
func (s *AlbumServiceImpl) AlbumUpdate(ctx context.Context, req *album.AlbumUpdateReq) (resp *album.AlbumUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// AlbumDelete implements the AlbumServiceImpl interface.
func (s *AlbumServiceImpl) AlbumDelete(ctx context.Context, req *album.AlbumDeleteReq) (resp *album.AlbumDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// AlbumPicAdd implements the AlbumPicServiceImpl interface.
func (s *AlbumPicServiceImpl) AlbumPicAdd(ctx context.Context, req *album_pic.AlbumPicAddReq) (resp *album_pic.AlbumPicAddResp, err error) {
	// TODO: Your code here...
	return
}

// AlbumPicList implements the AlbumPicServiceImpl interface.
func (s *AlbumPicServiceImpl) AlbumPicList(ctx context.Context, req *album_pic.AlbumPicListReq) (resp *album_pic.AlbumPicListResp, err error) {
	// TODO: Your code here...
	return
}

// AlbumPicUpdate implements the AlbumPicServiceImpl interface.
func (s *AlbumPicServiceImpl) AlbumPicUpdate(ctx context.Context, req *album_pic.AlbumPicUpdateReq) (resp *album_pic.AlbumPicUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// AlbumPicDelete implements the AlbumPicServiceImpl interface.
func (s *AlbumPicServiceImpl) AlbumPicDelete(ctx context.Context, req *album_pic.AlbumPicDeleteReq) (resp *album_pic.AlbumPicDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// BrandAdd implements the BrandServiceImpl interface.
func (s *BrandServiceImpl) BrandAdd(ctx context.Context, req *brand.BrandAddReq) (resp *brand.BrandAddResp, err error) {
	// TODO: Your code here...
	return
}

// BrandList implements the BrandServiceImpl interface.
func (s *BrandServiceImpl) BrandList(ctx context.Context, req *brand.BrandListReq) (resp *brand.BrandListResp, err error) {
	// TODO: Your code here...
	return
}

// BrandListByIds implements the BrandServiceImpl interface.
func (s *BrandServiceImpl) BrandListByIds(ctx context.Context, req *brand.BrandListByIdsReq) (resp *brand.BrandListResp, err error) {
	// TODO: Your code here...
	return
}

// BrandUpdate implements the BrandServiceImpl interface.
func (s *BrandServiceImpl) BrandUpdate(ctx context.Context, req *brand.BrandUpdateReq) (resp *brand.BrandUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// BrandDelete implements the BrandServiceImpl interface.
func (s *BrandServiceImpl) BrandDelete(ctx context.Context, req *brand.BrandDeleteReq) (resp *brand.BrandDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// CommentAdd implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAdd(ctx context.Context, req *comment.CommentAddReq) (resp *comment.CommentAddResp, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListReq) (resp *comment.CommentListResp, err error) {
	// TODO: Your code here...
	return
}

// CommentUpdate implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentUpdate(ctx context.Context, req *comment.CommentUpdateReq) (resp *comment.CommentUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// CommentDelete implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentDelete(ctx context.Context, req *comment.CommentDeleteReq) (resp *comment.CommentDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// CommentReplayAdd implements the CommentReplayServiceImpl interface.
func (s *CommentReplayServiceImpl) CommentReplayAdd(ctx context.Context, req *comment_reply.CommentReplayAddReq) (resp *comment_reply.CommentReplayAddResp, err error) {
	// TODO: Your code here...
	return
}

// CommentReplayList implements the CommentReplayServiceImpl interface.
func (s *CommentReplayServiceImpl) CommentReplayList(ctx context.Context, req *comment_reply.CommentReplayListReq) (resp *comment_reply.CommentReplayListResp, err error) {
	// TODO: Your code here...
	return
}

// CommentReplayUpdate implements the CommentReplayServiceImpl interface.
func (s *CommentReplayServiceImpl) CommentReplayUpdate(ctx context.Context, req *comment_reply.CommentReplayUpdateReq) (resp *comment_reply.CommentReplayUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// CommentReplayDelete implements the CommentReplayServiceImpl interface.
func (s *CommentReplayServiceImpl) CommentReplayDelete(ctx context.Context, req *comment_reply.CommentReplayDeleteReq) (resp *comment_reply.CommentReplayDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// FreightTemplateAdd implements the FreightTemplateServiceImpl interface.
func (s *FreightTemplateServiceImpl) FreightTemplateAdd(ctx context.Context, req *comment_reply.FreightTemplateAddReq) (resp *comment_reply.FreightTemplateAddResp, err error) {
	// TODO: Your code here...
	return
}

// FreightTemplateList implements the FreightTemplateServiceImpl interface.
func (s *FreightTemplateServiceImpl) FreightTemplateList(ctx context.Context, req *comment_reply.FreightTemplateListReq) (resp *comment_reply.FreightTemplateListResp, err error) {
	// TODO: Your code here...
	return
}

// FreightTemplateUpdate implements the FreightTemplateServiceImpl interface.
func (s *FreightTemplateServiceImpl) FreightTemplateUpdate(ctx context.Context, req *comment_reply.FreightTemplateUpdateReq) (resp *comment_reply.FreightTemplateUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// FreightTemplateDelete implements the FreightTemplateServiceImpl interface.
func (s *FreightTemplateServiceImpl) FreightTemplateDelete(ctx context.Context, req *comment_reply.FreightTemplateDeleteReq) (resp *comment_reply.FreightTemplateDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// MemberPriceAdd implements the MemberPriceServiceImpl interface.
func (s *MemberPriceServiceImpl) MemberPriceAdd(ctx context.Context, req *member_price.MemberPriceAddReq) (resp *member_price.MemberPriceAddResp, err error) {
	// TODO: Your code here...
	return
}

// MemberPriceList implements the MemberPriceServiceImpl interface.
func (s *MemberPriceServiceImpl) MemberPriceList(ctx context.Context, req *member_price.MemberPriceListReq) (resp *member_price.MemberPriceListResp, err error) {
	// TODO: Your code here...
	return
}

// MemberPriceUpdate implements the MemberPriceServiceImpl interface.
func (s *MemberPriceServiceImpl) MemberPriceUpdate(ctx context.Context, req *member_price.MemberPriceUpdateReq) (resp *member_price.MemberPriceUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// MemberPriceDelete implements the MemberPriceServiceImpl interface.
func (s *MemberPriceServiceImpl) MemberPriceDelete(ctx context.Context, req *member_price.MemberPriceDeleteReq) (resp *member_price.MemberPriceDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeAdd implements the ProductAttributeServiceImpl interface.
func (s *ProductAttributeServiceImpl) ProductAttributeAdd(ctx context.Context, req *product_attribute.ProductAttributeAddReq) (resp *product_attribute.ProductAttributeAddResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeList implements the ProductAttributeServiceImpl interface.
func (s *ProductAttributeServiceImpl) ProductAttributeList(ctx context.Context, req *product_attribute.ProductAttributeListReq) (resp *product_attribute.ProductAttributeListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeUpdate implements the ProductAttributeServiceImpl interface.
func (s *ProductAttributeServiceImpl) ProductAttributeUpdate(ctx context.Context, req *product_attribute.ProductAttributeUpdateReq) (resp *product_attribute.ProductAttributeUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeDelete implements the ProductAttributeServiceImpl interface.
func (s *ProductAttributeServiceImpl) ProductAttributeDelete(ctx context.Context, req *product_attribute.ProductAttributeDeleteReq) (resp *product_attribute.ProductAttributeDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeCategoryAdd implements the ProductAttributeCategoryServiceImpl interface.
func (s *ProductAttributeCategoryServiceImpl) ProductAttributeCategoryAdd(ctx context.Context, req *product_attribute_category.ProductAttributeCategoryAddReq) (resp *product_attribute_category.ProductAttributeCategoryAddResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeCategoryList implements the ProductAttributeCategoryServiceImpl interface.
func (s *ProductAttributeCategoryServiceImpl) ProductAttributeCategoryList(ctx context.Context, req *product_attribute_category.ProductAttributeCategoryListReq) (resp *product_attribute_category.ProductAttributeCategoryListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeCategoryUpdate implements the ProductAttributeCategoryServiceImpl interface.
func (s *ProductAttributeCategoryServiceImpl) ProductAttributeCategoryUpdate(ctx context.Context, req *product_attribute_category.ProductAttributeCategoryUpdateReq) (resp *product_attribute_category.ProductAttributeCategoryUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeCategoryDelete implements the ProductAttributeCategoryServiceImpl interface.
func (s *ProductAttributeCategoryServiceImpl) ProductAttributeCategoryDelete(ctx context.Context, req *product_attribute_category.ProductAttributeCategoryDeleteReq) (resp *product_attribute_category.ProductAttributeCategoryDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeValueAdd implements the ProductAttributeValueServiceImpl interface.
func (s *ProductAttributeValueServiceImpl) ProductAttributeValueAdd(ctx context.Context, req *product_attribute_value.ProductAttributeValueAddReq) (resp *product_attribute_value.ProductAttributeValueAddResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeValueList implements the ProductAttributeValueServiceImpl interface.
func (s *ProductAttributeValueServiceImpl) ProductAttributeValueList(ctx context.Context, req *product_attribute_value.ProductAttributeValueListReq) (resp *product_attribute_value.ProductAttributeValueListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeValueUpdate implements the ProductAttributeValueServiceImpl interface.
func (s *ProductAttributeValueServiceImpl) ProductAttributeValueUpdate(ctx context.Context, req *product_attribute_value.ProductAttributeValueUpdateReq) (resp *product_attribute_value.ProductAttributeValueUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAttributeValueDelete implements the ProductAttributeValueServiceImpl interface.
func (s *ProductAttributeValueServiceImpl) ProductAttributeValueDelete(ctx context.Context, req *product_attribute_value.ProductAttributeValueDeleteReq) (resp *product_attribute_value.ProductAttributeValueDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// ProductCategoryAdd implements the ProductCategoryServiceImpl interface.
func (s *ProductCategoryServiceImpl) ProductCategoryAdd(ctx context.Context, req *product_category.ProductCategoryAddReq) (resp *product_category.ProductCategoryAddResp, err error) {
	// TODO: Your code here...
	return
}

// ProductCategoryList implements the ProductCategoryServiceImpl interface.
func (s *ProductCategoryServiceImpl) ProductCategoryList(ctx context.Context, req *product_category.ProductCategoryListReq) (resp *product_category.ProductCategoryListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductCategoryUpdate implements the ProductCategoryServiceImpl interface.
func (s *ProductCategoryServiceImpl) ProductCategoryUpdate(ctx context.Context, req *product_category.ProductCategoryUpdateReq) (resp *product_category.ProductCategoryUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ProductCategoryDelete implements the ProductCategoryServiceImpl interface.
func (s *ProductCategoryServiceImpl) ProductCategoryDelete(ctx context.Context, req *product_category.ProductCategoryDeleteReq) (resp *product_category.ProductCategoryDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// QueryProductCategoryList implements the ProductCategoryServiceImpl interface.
func (s *ProductCategoryServiceImpl) QueryProductCategoryList(ctx context.Context, req *product_category.QueryProductCategoryListReq) (resp *product_category.QueryProductCategoryListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductCategoryAttributeRelationAdd implements the ProductCategoryAttributeRelationServiceImpl interface.
func (s *ProductCategoryAttributeRelationServiceImpl) ProductCategoryAttributeRelationAdd(ctx context.Context, req *product_category_attribute_relation.ProductCategoryAttributeRelationAddReq) (resp *product_category_attribute_relation.ProductCategoryAttributeRelationAddResp, err error) {
	// TODO: Your code here...
	return
}

// ProductCategoryAttributeRelationList implements the ProductCategoryAttributeRelationServiceImpl interface.
func (s *ProductCategoryAttributeRelationServiceImpl) ProductCategoryAttributeRelationList(ctx context.Context, req *product_category_attribute_relation.ProductCategoryAttributeRelationListReq) (resp *product_category_attribute_relation.ProductCategoryAttributeRelationListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductCategoryAttributeRelationUpdate implements the ProductCategoryAttributeRelationServiceImpl interface.
func (s *ProductCategoryAttributeRelationServiceImpl) ProductCategoryAttributeRelationUpdate(ctx context.Context, req *product_category_attribute_relation.ProductCategoryAttributeRelationUpdateReq) (resp *product_category_attribute_relation.ProductCategoryAttributeRelationUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ProductCategoryAttributeRelationDelete implements the ProductCategoryAttributeRelationServiceImpl interface.
func (s *ProductCategoryAttributeRelationServiceImpl) ProductCategoryAttributeRelationDelete(ctx context.Context, req *product_category_attribute_relation.ProductCategoryAttributeRelationDeleteReq) (resp *product_category_attribute_relation.ProductCategoryAttributeRelationDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// ProductFullReductionAdd implements the ProductFullReductionServiceImpl interface.
func (s *ProductFullReductionServiceImpl) ProductFullReductionAdd(ctx context.Context, req *product_full_reduction.ProductFullReductionAddReq) (resp *product_full_reduction.ProductFullReductionAddResp, err error) {
	// TODO: Your code here...
	return
}

// ProductFullReductionList implements the ProductFullReductionServiceImpl interface.
func (s *ProductFullReductionServiceImpl) ProductFullReductionList(ctx context.Context, req *product_full_reduction.ProductFullReductionListReq) (resp *product_full_reduction.ProductFullReductionListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductFullReductionUpdate implements the ProductFullReductionServiceImpl interface.
func (s *ProductFullReductionServiceImpl) ProductFullReductionUpdate(ctx context.Context, req *product_full_reduction.ProductFullReductionUpdateReq) (resp *product_full_reduction.ProductFullReductionUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ProductFullReductionDelete implements the ProductFullReductionServiceImpl interface.
func (s *ProductFullReductionServiceImpl) ProductFullReductionDelete(ctx context.Context, req *product_full_reduction.ProductFullReductionDeleteReq) (resp *product_full_reduction.ProductFullReductionDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// ProductLadderAdd implements the ProductLadderServiceImpl interface.
func (s *ProductLadderServiceImpl) ProductLadderAdd(ctx context.Context, req *product_ladder.ProductLadderAddReq) (resp *product_ladder.ProductLadderAddResp, err error) {
	// TODO: Your code here...
	return
}

// ProductLadderList implements the ProductLadderServiceImpl interface.
func (s *ProductLadderServiceImpl) ProductLadderList(ctx context.Context, req *product_ladder.ProductLadderListReq) (resp *product_ladder.ProductLadderListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductLadderUpdate implements the ProductLadderServiceImpl interface.
func (s *ProductLadderServiceImpl) ProductLadderUpdate(ctx context.Context, req *product_ladder.ProductLadderUpdateReq) (resp *product_ladder.ProductLadderUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ProductLadderDelete implements the ProductLadderServiceImpl interface.
func (s *ProductLadderServiceImpl) ProductLadderDelete(ctx context.Context, req *product_ladder.ProductLadderDeleteReq) (resp *product_ladder.ProductLadderDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// ProductOperateLogAdd implements the ProductOperateLogServiceImpl interface.
func (s *ProductOperateLogServiceImpl) ProductOperateLogAdd(ctx context.Context, req *product_full_reduction.ProductOperateLogAddReq) (resp *product_full_reduction.ProductOperateLogAddResp, err error) {
	// TODO: Your code here...
	return
}

// ProductOperateLogList implements the ProductOperateLogServiceImpl interface.
func (s *ProductOperateLogServiceImpl) ProductOperateLogList(ctx context.Context, req *product_full_reduction.ProductOperateLogListReq) (resp *product_full_reduction.ProductOperateLogListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductOperateLogUpdate implements the ProductOperateLogServiceImpl interface.
func (s *ProductOperateLogServiceImpl) ProductOperateLogUpdate(ctx context.Context, req *product_full_reduction.ProductOperateLogUpdateReq) (resp *product_full_reduction.ProductOperateLogUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ProductOperateLogDelete implements the ProductOperateLogServiceImpl interface.
func (s *ProductOperateLogServiceImpl) ProductOperateLogDelete(ctx context.Context, req *product_full_reduction.ProductOperateLogDeleteReq) (resp *product_full_reduction.ProductOperateLogDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// ProductAdd implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductAdd(ctx context.Context, req *product.ProductAddReq) (resp *product.ProductAddResp, err error) {
	// TODO: Your code here...
	return
}

// ProductList implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductList(ctx context.Context, req *product.ProductListReq) (resp *product.ProductListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductListByIds implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductListByIds(ctx context.Context, req *product.ProductByIdsReq) (resp *product.ProductListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductUpdate implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductUpdate(ctx context.Context, req *product.ProductUpdateReq) (resp *product.ProductUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ProductDelete implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductDelete(ctx context.Context, req *product.ProductDeleteReq) (resp *product.ProductDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// ProductDetailById implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductDetailById(ctx context.Context, req *product.ProductDetailByIdReq) (resp *product.ProductDetailByIdResp, err error) {
	// TODO: Your code here...
	return
}

// ProductVerifyRecordAdd implements the ProductVerifyRecordServiceImpl interface.
func (s *ProductVerifyRecordServiceImpl) ProductVerifyRecordAdd(ctx context.Context, req *product_verify_record.ProductVerifyRecordAddReq) (resp *product_verify_record.ProductVerifyRecordAddResp, err error) {
	// TODO: Your code here...
	return
}

// ProductVerifyRecordList implements the ProductVerifyRecordServiceImpl interface.
func (s *ProductVerifyRecordServiceImpl) ProductVerifyRecordList(ctx context.Context, req *product_verify_record.ProductVerifyRecordListReq) (resp *product_verify_record.ProductVerifyRecordListResp, err error) {
	// TODO: Your code here...
	return
}

// ProductVerifyRecordUpdate implements the ProductVerifyRecordServiceImpl interface.
func (s *ProductVerifyRecordServiceImpl) ProductVerifyRecordUpdate(ctx context.Context, req *product_verify_record.ProductVerifyRecordUpdateReq) (resp *product_verify_record.ProductVerifyRecordUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ProductVerifyRecordDelete implements the ProductVerifyRecordServiceImpl interface.
func (s *ProductVerifyRecordServiceImpl) ProductVerifyRecordDelete(ctx context.Context, req *product_verify_record.ProductVerifyRecordDeleteReq) (resp *product_verify_record.ProductVerifyRecordDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// SkuStockAdd implements the SkuStockServiceImpl interface.
func (s *SkuStockServiceImpl) SkuStockAdd(ctx context.Context, req *sku_stock.SkuStockAddReq) (resp *sku_stock.SkuStockAddResp, err error) {
	// TODO: Your code here...
	return
}

// SkuStockList implements the SkuStockServiceImpl interface.
func (s *SkuStockServiceImpl) SkuStockList(ctx context.Context, req *sku_stock.SkuStockListReq) (resp *sku_stock.SkuStockListResp, err error) {
	// TODO: Your code here...
	return
}

// SkuStockUpdate implements the SkuStockServiceImpl interface.
func (s *SkuStockServiceImpl) SkuStockUpdate(ctx context.Context, req *sku_stock.SkuStockUpdateReq) (resp *sku_stock.SkuStockUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// SkuStockDelete implements the SkuStockServiceImpl interface.
func (s *SkuStockServiceImpl) SkuStockDelete(ctx context.Context, req *sku_stock.SkuStockDeleteReq) (resp *sku_stock.SkuStockDeleteResp, err error) {
	// TODO: Your code here...
	return
}

// ReleaseSkuStockLock implements the SkuStockServiceImpl interface.
func (s *SkuStockServiceImpl) ReleaseSkuStockLock(ctx context.Context, req *sku_stock.ReleaseSkuStockLockReq) (resp *sku_stock.ReleaseSkuStockLockResp, err error) {
	// TODO: Your code here...
	return
}

// LockSkuStockLock implements the SkuStockServiceImpl interface.
func (s *SkuStockServiceImpl) LockSkuStockLock(ctx context.Context, req *sku_stock.LockSkuStockLockReq) (resp *sku_stock.LockSkuStockLockResp, err error) {
	// TODO: Your code here...
	return
}

// QuerySkuStockByProductSkuId implements the SkuStockServiceImpl interface.
func (s *SkuStockServiceImpl) QuerySkuStockByProductSkuId(ctx context.Context, req *sku_stock.QuerySkuStockByProductSkuIdReq) (resp *sku_stock.BaseSkuStock, err error) {
	// TODO: Your code here...
	return
}
