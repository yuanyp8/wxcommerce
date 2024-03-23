/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:33
 * @Desc:
 */

package model

import (
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
	"strings"
)

// PmsProductAttribute 商品属性参数表
type PmsProductAttribute struct {
	gorm.Model

	// 表字段对应的结构体属性
	ProductAttributeCategoryId int64  `gorm:"column:product_attribute_category_id;not null"`
	Name                       string `gorm:"column:name;not_null;size:64"`
	SelectType                 int8   `gorm:"column:select_type;not null;comment:'属性选择类型：0->唯一；1->单选；2->多选'"`
	InputType                  int8   `gorm:"column:input_type;not null;comment:'属性录入方式：0->手工录入；1->从列表中选取'"`
	InputList                  string `gorm:"column:input_list;not null;size:255;comment:'可选值列表，以逗号隔开'"`
	Sort                       int    `gorm:"column:sort;not null;comment:'排序字段：最高的可以单独上传图片'"`
	FilterType                 int8   `gorm:"column:filter_type;not null;comment:'分类筛选样式：1->普通；1->颜色'"`
	SearchType                 int8   `gorm:"column:search_type;not null;comment:'检索类型；0->不需要进行检索；1->关键字检索；2->范围检索'"`
	RelatedStatus              int8   `gorm:"column:related_status;not null;comment:'相同属性产品是否关联；0->不关联；1->关联'"`
	HandAddStatus              int8   `gorm:"column:hand_add_status;not null;comment:'是否支持手动新增；0->不支持；1->支持'"`
	Type                       int8   `gorm:"column:type;not null;comment:'属性的类型；0->规格；1->参数'"`
}

// TableName 设置表名
func (PmsProductAttribute) TableName() string {
	return constants.TableNamePmsProductAttribute
}

// GetInputListAsSlice 将输入列表字符串转换为切片
func (p *PmsProductAttribute) GetInputListAsSlice() []string {
	return strings.Split(p.InputList, ",")
}
