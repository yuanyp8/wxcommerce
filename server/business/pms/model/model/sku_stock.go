/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:46
 * @Desc:
 */

package model

import (
	"encoding/json"
	"fmt"
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
)

// PmsSkuStock SKU的库存
type PmsSkuStock struct {
	gorm.Model

	// 表字段对应的结构体属性
	ProductID      uint64  `gorm:"column:product_id;not null"`
	SkuCode        string  `gorm:"column:sku_code;size:64;not null;comment:'SKU编码'"`
	Price          float64 `gorm:"column:price;type:decimal(10,2);not null"`
	Stock          int     `gorm:"column:stock;default:0;not null;comment:'库存'"`
	LowStock       int     `gorm:"column:low_stock;not null;comment:'预警库存'"`
	Pic            string  `gorm:"column:pic;size:255;not null;comment:'展示图片'"`
	Sale           int     `gorm:"column:sale;not null;comment:'销量'"`
	PromotionPrice float64 `gorm:"column:promotion_price;type:decimal(10,2);not null;comment:'单品促销价格'"`
	LockStock      int     `gorm:"column:lock_stock;default:0;not null;comment:'锁定库存'"`
	SpData         string  `gorm:"column:sp_data;size:500;not null;comment:'商品销售属性，JSON格式'"`
}

// TableName 设置表名
func (PmsSkuStock) TableName() string {
	return constants.TableNamePmsSkuStock
}

// GetSpDataMap 将SpData字符串转换为map[string]string格式
func (p *PmsSkuStock) GetSpDataMap() (map[string]string, error) {
	if p.SpData == "" {
		return nil, nil
	}

	spDataMap := make(map[string]string)
	err := json.Unmarshal([]byte(p.SpData), &spDataMap)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spData: %w", err)
	}

	return spDataMap, nil
}

// SetSpDataMap 将map[string]string格式的数据转换为SpData字符串
func (p *PmsSkuStock) SetSpDataMap(spDataMap map[string]string) error {
	if spDataMap == nil {
		p.SpData = ""
		return nil
	}

	spData, err := json.Marshal(spDataMap)
	if err != nil {
		return fmt.Errorf("failed to marshal spData: %w", err)
	}

	p.SpData = string(spData)
	return nil
}

// GetSaleQuantity 获取销售数量（销量）
func (p *PmsSkuStock) GetSaleQuantity() int {
	return p.Sale
}

// SetSaleQuantity 设置销售数量（销量）
func (p *PmsSkuStock) SetSaleQuantity(quantity int) {
	p.Sale = quantity
}

// GetStock 获取库存数量
func (p *PmsSkuStock) GetStock() int {
	return p.Stock
}

// SetStock 设置库存数量
func (p *PmsSkuStock) SetStock(stock int) {
	p.Stock = stock
}

// GetLockStock 获取锁定库存数量
func (p *PmsSkuStock) GetLockStock() int {
	return p.LockStock
}

// SetLockStock 设置锁定库存数量
func (p *PmsSkuStock) SetLockStock(lockStock int) {
	p.LockStock = lockStock
}

// GetPrice 获取价格
func (p *PmsSkuStock) GetPrice() float64 {
	return p.Price
}

// SetPrice 设置价格
func (p *PmsSkuStock) SetPrice(price float64) {
	p.Price = price
}

// GetPromotionPrice 获取促销价格
func (p *PmsSkuStock) GetPromotionPrice() float64 {
	return p.PromotionPrice
}

// SetPromotionPrice 设置促销价格
func (p *PmsSkuStock) SetPromotionPrice(promotionPrice float64) {
	p.PromotionPrice = promotionPrice
}

// GetSkuCode 获取SKU编码
func (p *PmsSkuStock) GetSkuCode() string {
	return p.SkuCode
}

// SetSkuCode 设置SKU编码
func (p *PmsSkuStock) SetSkuCode(skuCode string) {
	p.SkuCode = skuCode
}

// GetPic 获取展示图片
func (p *PmsSkuStock) GetPic() string {
	return p.Pic
}

// SetPic 设置展示图片
func (p *PmsSkuStock) SetPic(pic string) {
	p.Pic = pic
}

// GetLowStock 获取预警库存
func (p *PmsSkuStock) GetLowStock() int {
	return p.LowStock
}

// SetLowStock 设置预警库存
func (p *PmsSkuStock) SetLowStock(lowStock int) {
	p.LowStock = lowStock
}
