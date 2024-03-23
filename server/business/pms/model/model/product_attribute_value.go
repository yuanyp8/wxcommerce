/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:34
 * @Desc:
 */

package model

import (
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
	"strings"
)

// PmsProductAttributeValue 存储产品参数信息的表
type PmsProductAttributeValue struct {
	gorm.Model

	// 表字段对应的结构体属性
	ProductID          int64  `gorm:"column:product_id;not null"`
	ProductAttributeID int64  `gorm:"column:product_attribute_id;not null"`
	Value              string `gorm:"column:value"`
}

// TableName 设置表名
func (PmsProductAttributeValue) TableName() string {
	return constants.TableNamePmsProductAttributeValue
}

// GetValueSlice 将值字段转换为字符串切片
func (p *PmsProductAttributeValue) GetValueSlice() []string {
	if p.Value == "" {
		return nil
	}
	return strings.Split(p.Value, ",")
}

// SetValueSlice 设置值字段为字符串切片
func (p *PmsProductAttributeValue) SetValueSlice(values []string) {
	if len(values) == 0 {
		p.Value = ""
		return
	}
	p.Value = strings.Join(values, ",")
}
