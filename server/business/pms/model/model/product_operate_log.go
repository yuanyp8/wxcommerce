/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:43
 * @Desc:
 */

package model

import (
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
	"time"
)

// PmsProductOperateLog 产品操作日志表
type PmsProductOperateLog struct {
	gorm.Model

	// 表字段对应的结构体属性
	ProductId        uint64    `gorm:"column:product_id;not null"`
	PriceOld         float64   `gorm:"column:price_old;type:decimal(10,2);not null;comment:'原价'"`
	PriceNew         float64   `gorm:"column:price_new;type:decimal(10,2);not null;comment:'新价'"`
	SalePriceOld     float64   `gorm:"column:sale_price_old;type:decimal(10,2);not null;comment:'原促销价'"`
	SalePriceNew     float64   `gorm:"column:sale_price_new;type:decimal(10,2);not null;comment:'新促销价'"`
	GiftPointOld     uint      `gorm:"column:gift_point_old;not null;comment:'原赠送积分'"`
	GiftPointNew     uint      `gorm:"column:gift_point_new;not null;comment:'新赠送积分'"`
	UsePointLimitOld uint      `gorm:"column:use_point_limit_old;not null;comment:'原使用积分限制'"`
	UsePointLimitNew uint      `gorm:"column:use_point_limit_new;not null;comment:'新使用积分限制'"`
	OperateMan       string    `gorm:"column:operate_man;size:64;not null;comment:'操作人'"`
	CreateTime       time.Time `gorm:"column:create_time;not null;comment:'操作时间';autoCreateTime"`
}

// TableName 设置表名
func (PmsProductOperateLog) TableName() string {
	return constants.TableNamePmsProductOperateLog
}
