/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:35
 * @Desc:
 */

package model

import (
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
)

// PmsProductCategory 产品分类表
type PmsProductCategory struct {
	gorm.Model

	// 表字段对应的结构体属性
	ParentID     int64  `gorm:"column:parent_id;not null;comment:'上级分类的编号：0表示一级分类'"`
	Name         string `gorm:"column:name;not null;size:64"`
	Level        int8   `gorm:"column:level;not null;comment:'分类级别：0->1级；1->2级'"`
	ProductCount int    `gorm:"column:product_count;not null"`
	ProductUnit  string `gorm:"column:product_unit;not null;size:64"`
	NavStatus    int8   `gorm:"column:nav_status;not null;comment:'是否显示在导航栏：0->不显示；1->显示'"`
	ShowStatus   int8   `gorm:"column:show_status;not null;comment:'显示状态：0->不显示；1->显示'"`
	Sort         int    `gorm:"column:sort;not null"`
	Icon         string `gorm:"column:icon;not null;size:255;comment:'图标'"`
	Keywords     string `gorm:"column:keywords;not null;size:255"`
	Description  string `gorm:"column:description;type:text"`
}

// TableName 设置表名
func (PmsProductCategory) TableName() string {
	return constants.TableNamePmsProductCategory
}
