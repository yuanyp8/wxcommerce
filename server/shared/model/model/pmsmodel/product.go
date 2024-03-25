/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:26
 * @Desc:
 */

package pmsmodel

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

// PmsProduct 商品信息表
type PmsProduct struct {
	gorm.Model

	// 表字段对应的结构体属性
	BrandId                    int64     `gorm:"column:brand_id;not null;comment:'品牌id'"`
	ProductCategoryId          int64     `gorm:"column:product_category_id;not null;comment:'商品分类id'"`
	FreightTemplateId          int64     `gorm:"column:freight_template_id;not null;comment:'商品运费模板id'"`
	ProductAttributeCategoryId int64     `gorm:"column:product_attribute_category_id;not null;comment:'商品属性分类id'"`
	Name                       string    `gorm:"column:name;not null;size:64;comment:'商品名称'"`
	Pic                        string    `gorm:"column:pic;not null;size:255;comment:'商品图片'"`
	ProductSn                  string    `gorm:"column:product_sn;not null;size:64;comment:'货号'"`
	DeleteStatus               int8      `gorm:"column:delete_status;not null;comment:'删除状态：0->未删除；1->已删除'"`
	PublishStatus              int8      `gorm:"column:publish_status;not null;comment:'上架状态：0->下架；1->上架'"`
	NewStatus                  int8      `gorm:"column:new_status;not null;comment:'新品状态:0->不是新品；1->新品'"`
	RecommendStatus            int8      `gorm:"column:recommend_status;not null;comment:'推荐状态；0->不推荐；1->推荐'"`
	VerifyStatus               int8      `gorm:"column:verify_status;not null;comment:'审核状态：0->未审核；1->审核通过'"`
	Sort                       int       `gorm:"column:sort;not null;comment:'排序'"`
	Sale                       int       `gorm:"column:sale;not null;comment:'销量'"`
	Price                      float64   `gorm:"column:price;not null;type:decimal(10, 2);comment:'商品价格'"`
	PromotionPrice             float64   `gorm:"column:promotion_price;not null;type:decimal(10, 2);comment:'促销价格'"`
	GiftGrowth                 int       `gorm:"column:gift_growth;default:0;not null;comment:'赠送的成长值'"`
	GiftPoint                  int       `gorm:"column:gift_point;default:0;not null;comment:'赠送的积分'"`
	UsePointLimit              int       `gorm:"column:use_point_limit;not null;comment:'限制使用的积分数'"`
	SubTitle                   string    `gorm:"column:sub_title;not null;size:255;comment:'副标题'"`
	Description                string    `gorm:"column:description;text;not null;comment:'商品描述'"`
	OriginalPrice              float64   `gorm:"column:original_price;not null;type:decimal(10, 2);comment:'市场价'"`
	Stock                      int       `gorm:"column:stock;not null;comment:'库存'"`
	LowStock                   int       `gorm:"column:low_stock;not null;comment:'库存预警值'"`
	Unit                       string    `gorm:"column:unit;not null;size:16;comment:'单位'"`
	Weight                     float64   `gorm:"column:weight;not null;type:decimal(10, 2);comment:'商品重量，默认为克'"`
	PreviewStatus              int8      `gorm:"column:preview_status;not null;comment:'是否为预告商品：0->不是；1->是'"`
	ServiceIds                 string    `gorm:"column:service_ids;not null;size:64;comment:'以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮'"`
	Keywords                   string    `gorm:"column:keywords;not null;size:255;comment:'搜索关键字'"`
	Note                       string    `gorm:"column:note;not null;size:255;comment:'备注'"`
	AlbumPics                  string    `gorm:"column:album_pics;not null;size:255;comment:'画册图片，连产品图片限制为5张，以逗号分割'"`
	_DetailTitle               string    `gorm:"column:detail_title;not null;size:255;comment:'详情标题'"`
	DetailDesc                 string    `gorm:"column:detail_desc;text;not null;comment:'详情描述'"`
	DetailHtml                 string    `gorm:"column:detail_html;text;not null;comment:'产品详情网页内容'"`
	DetailMobileHtml           string    `gorm:"column:detail_mobile_html;text;not null;comment:'移动端网页详情'"`
	PromotionStartTime         time.Time `gorm:"column:promotion_start_time;not null;comment:'促销开始时间'"`
	PromotionEndTime           time.Time `gorm:"column:promotion_end_time;not null;comment:'促销结束时间'"`
	PromotionPerLimit          int       `gorm:"column:promotion_per_limit;not null;comment:'活动限购数量'"`
	PromotionType              int8      `gorm:"column:promotion_type;not null;comment:'促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购'"`
	BrandName                  string    `gorm:"column:brand_name;not null;size:255;comment:'品牌名称'"`
	ProductCategoryName        string    `gorm:"column:product_category_name;not null;size:255;comment:'商品分类名称'"`
	ProductCategoryIdArray     string    `gorm:"column:product_category_id_array;not null;size:64;comment:'商品分类id字符串'"`
}

// TableName 设置表名
func (*PmsProduct) TableName() string {
	return constants.TableNamePmsProduct
}

// GetProductCategoryIdArray 将product_category_id_array转换为整数数组
func (m *PmsProduct) GetProductCategoryIdArray() []int64 {
	ids := strings.Split(m.ProductCategoryIdArray, ",")
	categoryIds := make([]int64, len(ids))
	for i, v := range ids {
		id, _ := strconv.ParseInt(v, 10, 64)
		categoryIds[i] = id
	}
	return categoryIds
}

func (m *PmsProduct) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if m.ID != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(constants.PmsProductSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	m.ID = uint(sf.Generate().Int64())
	return nil
}
