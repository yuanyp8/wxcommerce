/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:38
 * @Desc:
 */

package model

import (
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
	"time"
)

// PmsProductCollect 收藏表
type PmsProductCollect struct {
	gorm.Model

	// 用户表的用户ID
	UserId uint `gorm:"column:user_id;not null;default:0;comment:'用户表的用户ID'"`
	// 如果collect_type=0，则是商品ID；如果collect_type=1，则是专题ID
	ValueId     uint  `gorm:"column:value_id;not null;default:0;comment:'如果type=0，则是商品ID；如果type=1，则是专题ID'"`
	CollectType uint8 `gorm:"column:collect_type;not null;default:0;comment:'收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID'"`
	// 创建时间
	AddTime time.Time `gorm:"column:add_time;not null;comment:'创建时间'"`
	// 更新时间
	UpdateTime *time.Time `gorm:"column:update_time;null;comment:'更新时间'"`
	// 逻辑删除
	Deleted bool `gorm:"column:deleted;not null;default:false;comment:'逻辑删除'"`
}

// TableName 设置表名
func (PmsProductCollect) TableName() string {
	return constants.TableNamePmsProductCollect
}
