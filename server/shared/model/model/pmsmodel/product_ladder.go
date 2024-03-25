/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:42
 * @Desc:
 */

package pmsmodel

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
)

// PmsProductLadder 产品阶梯价格表
type PmsProductLadder struct {
	gorm.Model

	// 表字段对应的结构体属性
	ProductId uint64  `gorm:"column:product_id;not null"`
	Count     int     `gorm:"column:count;not null;comment:'满足购买数量'"`
	Discount  float64 `gorm:"column:discount;type:decimal(10,2);not null;comment:'折扣比例'"`
	Price     float64 `gorm:"column:price;type:decimal(10,2);not null;comment:'达到指定数量后的折后单价（如果为空，则代表直接打折）'"`
}

// TableName 设置表名
func (*PmsProductLadder) TableName() string {
	return constants.TableNamePmsProductLadder
}

// GetDiscountedPrice 获取折后价格
func (m *PmsProductLadder) GetDiscountedPrice(originalPrice float64) float64 {
	if m.Price > 0 {
		return m.Price
	}
	discountedPrice := originalPrice * (1 - m.Discount/100)
	return discountedPrice
}

func (m *PmsProductLadder) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if m.ID != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(constants.PmsProductLadderSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	m.ID = uint(sf.Generate().Int64())
	return nil
}
