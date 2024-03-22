package db

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/21 17:22
 * @Desc:
 */

import (
	"context"
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	ID                         uint64    `gorm:"primaryKey;column:id"`
	BrandID                    uint64    `gorm:"column:brand_id;not null"`
	ProductCategoryId          uint64    `gorm:"column:product_category_id;not null"`               // 商品分类id
	ProductCategoryIdArray     string    `gorm:"column:product_category_id_Array;size:64;not null"` // 商品分类id字符串
	FreightTemplateId          uint64    `gorm:"column:freight_template_id;not null"`
	ProductAttributeCategoryId uint64    `gorm:"column:product_attribute_category_id;not null"`
	Name                       string    `gorm:"column:name;size:64;not null"`
	Pic                        string    `gorm:"column:pic;size:255;not null"`
	ProductSn                  string    `gorm:"column:product_sn;size:64;not null"`
	DeleteStatus               int       `gorm:"column:delete_status;not null"`
	PublishStatus              int       `gorm:"column:publish_status;not null"`
	NewStatus                  int       `gorm:"column:new_status;not null"`
	RecommendStatus            int       `gorm:"column:recommend_status;not null"`
	VerifyStatus               int       `gorm:"column:verify_status;not null"`
	PreviewStatus              int       `gorm:"column:preview_status;not null"` // 是否为预告商品：0->不是；1->是
	Sort                       int       `gorm:"column:sort;not null"`
	Sale                       int       `gorm:"column:sale;not null"`
	Price                      float64   `gorm:"column:price;type:decimal(10,2);not null"`
	PromotionPrice             float64   `gorm:"column:promotion_price;type:decimal(10,2);not null"`
	GiftGrowth                 uint      `gorm:"column:gift_growth;default:0;not null"`
	GiftPoint                  uint      `gorm:"column:gift_point;default:0;not null"`
	UsePointLimit              int       `gorm:"column:use_point_limit;not null"` // 使用的积分上限
	SubTitle                   string    `gorm:"column:sub_title;size:255;not null"`
	Description                string    `gorm:"column:description;text;not null"`
	OriginalPrice              float64   `gorm:"column:original_price;type:decimal(10,2);not null"`
	Stock                      int       `gorm:"column:stock;not null"`
	LowStock                   int       `gorm:"column:low_stock;not null"`
	Unit                       string    `gorm:"column:unit;size:16;not null"`
	Weight                     float64   `gorm:"column:weight;type:decimal(10,2);not null"`
	ServiceIds                 string    `gorm:"column:service_ids;size:64;not null"`
	Keywords                   string    `gorm:"column:keywords;size:255;not null"`
	Note                       string    `gorm:"column:note;size:255;not null"`
	AlbumPics                  string    `gorm:"column:album_pics;size:255;not null"`
	DetailTitle                string    `gorm:"column:detail_title;size:255;not null"`
	DetailDesc                 string    `gorm:"column:detail_desc;text;not null"`
	DetailHtml                 string    `gorm:"column:detail_html;text;not null"`
	DetailMobileHtml           string    `gorm:"column:detail_mobile_html;text;not null"`
	PromotionStartTime         time.Time `gorm:"column:promotion_start_time;not null"`
	PromotionEndTime           time.Time `gorm:"column:promotion_end_time;not null"`
	PromotionPerLimit          int       `gorm:"column:promotion_per_limit;not null"`
	PromotionType              int       `gorm:"column:promotion_type;not null"`
	BrandName                  string    `gorm:"column:brand_name;size:255;not null"`
	ProductCategoryName        string    `gorm:"column:product_category_name;size:255;not null"`
}

func (Product) TableName() string {
	return constants.PMSProductTableName
}

func ProductAdd(ctx context.Context, data *Product) error {
	//return db.WithContext(ctx).Create(data).Error
	return nil
}
