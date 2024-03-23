/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:37
 * @Desc:
 */

package model

import (
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
)

// PmsProductCategoryAttributeRelation 产品的分类和属性的关系表
type PmsProductCategoryAttributeRelation struct {
	gorm.Model

	// 表字段对应的结构体属性
	ProductCategoryId  int64 `gorm:"column:product_category_id;not null"`
	ProductAttributeId int64 `gorm:"column:product_attribute_id;not null"`
}

// TableName 设置表名
func (PmsProductCategoryAttributeRelation) TableName() string {
	return constants.TableNamePmsProductCategoryAttributeRelation
}
