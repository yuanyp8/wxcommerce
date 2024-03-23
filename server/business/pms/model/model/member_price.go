/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:25
 * @Desc:
 */

package model

import (
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
)

// PmsMemberPrice 商品会员价格表
type PmsMemberPrice struct {
	gorm.Model

	// 表字段对应的结构体属性
	ProductId       int64   `gorm:"column:product_id;not null;comment:'商品id'"`
	MemberLevelId   int64   `gorm:"column:member_level_id;not null;comment:'会员等级id'"`
	MemberPrice     float64 `gorm:"column:member_price;not null;type:decimal(10, 2);comment:'会员价格'"`
	MemberLevelName string  `gorm:"column:member_level_name;not null;size:100;comment:'会员等级名称'"`
}

// TableName 设置表名
func (PmsMemberPrice) TableName() string {
	return constants.TableNamePmsMemberPrice
}
