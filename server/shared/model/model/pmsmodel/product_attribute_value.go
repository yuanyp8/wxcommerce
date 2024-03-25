/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:34
 * @Desc:
 */

package pmsmodel

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
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
func (*PmsProductAttributeValue) TableName() string {
	return constants.TableNamePmsProductAttributeValue
}

// GetValueSlice 将值字段转换为字符串切片
func (m *PmsProductAttributeValue) GetValueSlice() []string {
	if m.Value == "" {
		return nil
	}
	return strings.Split(m.Value, ",")
}

// SetValueSlice 设置值字段为字符串切片
func (m *PmsProductAttributeValue) SetValueSlice(values []string) {
	if len(values) == 0 {
		m.Value = ""
		return
	}
	m.Value = strings.Join(values, ",")
}

func (m *PmsProductAttributeValue) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if m.ID != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(constants.PmsProductAttributeValueSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	m.ID = uint(sf.Generate().Int64())
	return nil
}
