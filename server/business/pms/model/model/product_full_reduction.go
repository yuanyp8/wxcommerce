/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:40
 * @Desc:
 */

package model

import (
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
	"math"
	"strconv"
)

// PmsProductFullReduction 产品满减表
type PmsProductFullReduction struct {
	gorm.Model

	// 表字段对应的结构体属性
	ProductId   uint64  `gorm:"column:product_id;not null"`
	FullPrice   float64 `gorm:"column:full_price;type:decimal(10,2);not null;comment:'满多少金额'"`
	ReducePrice float64 `gorm:"column:reduce_price;type:decimal(10,2);not null;comment:'减少多少金额'"`
}

// TableName 设置表名
func (PmsProductFullReduction) TableName() string {
	return constants.TableNamePmsProductFullReduction
}

// GetDiscountPercentage 获取满减折扣百分比（四舍五入到两位小数）
func (p *PmsProductFullReduction) GetDiscountPercentage() float64 {
	if p.FullPrice == 0 || p.ReducePrice == 0 {
		return 0.00
	}
	discount := p.ReducePrice / p.FullPrice * 100
	return roundToTwoDecimals(discount)
}

// roundToTwoDecimals 四舍五入保留两位小数
func roundToTwoDecimals(f float64) float64 {
	return math.Round(f*100) / 100
}

// 字符串转浮点数（由于 Go 没有内置函数可以直接实现，此处假设 toFloat64 是一个自定义的将字符串转为浮点数的函数）
func toFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
