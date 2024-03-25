package constants

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/21 23:36
 * @Desc:
 */

const (
	// PMS Table Name
	TableNamePmsAlbum                            = "pms_album"
	TableNamePmsAlbumPic                         = "pms_album_pic"
	TableNamePmsBrand                            = "pms_brand"
	TableNamePmsComment                          = "pms_comment"
	TableNamePmsCommentReply                     = "pms_comment_reply"
	TableNamePmsFreightTemplate                  = "pms_freight_template"
	TableNamePmsMemberPrice                      = "pms_member_price"
	TableNamePmsProduct                          = "pms_product"
	TableNamePmsProductAttribute                 = "pms_product_attribute"
	TableNamePmsProductAttributeCategory         = "pms_product_attribute_category"
	TableNamePmsProductAttributeValue            = "pms_product_attribute_value"
	TableNamePmsProductCategory                  = "pms_product_category"
	TableNamePmsProductCategoryAttributeRelation = "pms_product_category_attribute_relation"
	TableNamePmsProductCollect                   = "pms_product_collect"
	TableNamePmsProductFullReduction             = "pms_product_full_reduction"
	TableNamePmsProductLadder                    = "pms_product_ladder"
	TableNamePmsProductOperateLog
	TableNamePmsProductVerifyRecord = "pms_product_verify_record"
	TableNamePmsSkuStock            = "pms_sku_stock"

	PMSConfigPath       = "./pmsmodel/config/config.yaml"
	PMSProductTableName = "pms_product"
	DatabaseName        = "wx_commerce"

	MySqlDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	KlogFilePath = "./tmp/klog/logs/"
)

// 雪花id的node
const (
	PmsAlbumSnowflakeNode                            = 2
	PmsAlbumPicSnowflakeNode                         = 3
	PmsBrandSnowflakeNode                            = 4
	PmsCommentSnowflakeNode                          = 5
	PmsCommentReplySnowflakeNode                     = 6
	PmsFreightTemplateSnowflakeNode                  = 7
	PmsMemberPriceSnowflakeNode                      = 8
	PmsProductSnowflakeNode                          = 9
	PmsProductAttributeSnowflakeNode                 = 10
	PmsProductAttributeCategorySnowflakeNode         = 11
	PmsProductAttributeValueSnowflakeNode            = 12
	PmsProductCategorySnowflakeNode                  = 13
	PmsProductCategoryAttributeRelationSnowflakeNode = 14
	PmsProductCollectSnowflakeNode                   = 15
	PmsProductFullReductionSnowflakeNode             = 16
	PmsProductLadderSnowflakeNode                    = 17
	PmsProductOperateLogSnowflakeNode                = 18
	PmsProductVerifyRecordSnowflakeNode              = 19
	PmsSkuStockSnowflakeNode                         = 20
)
