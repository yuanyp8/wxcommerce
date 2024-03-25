/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:23
 * @Desc:
 */

package pmsmodel

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
)

// PmsFreightTemplate 运费模版表
type PmsFreightTemplate struct {
	gorm.Model

	// 表字段对应的结构体属性
	Name           string  `gorm:"column:name;not null;size:64;comment:'运费模版名称'"`
	ChargeType     int     `gorm:"column:charge_type;not null;comment:'计费类型：0->按重量；1->按件数'"`
	FirstWeight    float64 `gorm:"column:first_weight;not null;precision:10;scale:2;comment:'首重kg'"`
	FirstFee       float64 `gorm:"column:first_fee;not null;precision:10;scale:2;comment:'首费（元）'"`
	ContinueWeight float64 `gorm:"column:continue_weight;not null;precision:10;scale:2"`
	ContinueFee    float64 `gorm:"column:continue_fee;not null;precision:10;scale:2"`
	Dest           string  `gorm:"column:dest;not null;size:255;comment:'目的地（省、市）'"`
}

// TableName 设置表名
func (*PmsFreightTemplate) TableName() string {
	return constants.TableNamePmsFreightTemplate
}

func (m *PmsFreightTemplate) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if m.ID != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(constants.PmsFreightTemplateSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	m.ID = uint(sf.Generate().Int64())
	return nil
}
