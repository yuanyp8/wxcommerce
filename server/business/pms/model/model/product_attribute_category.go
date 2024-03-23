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
)

// PmsProductAttributeCategory 产品属性分类表
type PmsProductAttributeCategory struct {
	gorm.Model

	// 表字段对应的结构体属性
	Name           string `gorm:"column:name;not null;size:64"`
	AttributeCount int    `gorm:"column:attribute_count;default:0;not null;comment:'属性数量'"`
	ParamCount     int    `gorm:"column:param_count;default:0;not null;comment:'参数数量'"`
}

// TableName 设置表名
func (PmsProductAttributeCategory) TableName() string {
	return constants.TableNamePmsProductAttributeCategory
}
