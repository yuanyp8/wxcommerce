/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:43
 * @Desc:
 */

package pmsmodel

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
	"time"
)

// PmsProductVerifyRecord 商品审核记录表
type PmsProductVerifyRecord struct {
	gorm.Model

	// 表字段对应的结构体属性
	ProductID  uint64    `gorm:"column:product_id;not null"`
	CreateTime time.Time `gorm:"column:create_time;not null;comment:'创建时间'"`
	VerifyMan  string    `gorm:"column:verify_man;size:64;not null;comment:'审核人'"`
	Status     int8      `gorm:"column:status;not null;comment:'审核状态'"`
	Detail     string    `gorm:"column:detail;not null;comment:'审核反馈详情'"`
}

// TableName 设置表名
func (*PmsProductVerifyRecord) TableName() string {
	return constants.TableNamePmsProductVerifyRecord
}

// GetStatusText 获取审核状态的文本描述
func (m *PmsProductVerifyRecord) GetStatusText() string {
	switch m.Status {
	case 0:
		return "审核通过"
	case 1:
		return "审核拒绝"
	default:
		return "未知状态"
	}
}

func (m *PmsProductVerifyRecord) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if m.ID != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(constants.PmsProductVerifyRecordSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	m.ID = uint(sf.Generate().Int64())
	return nil
}
