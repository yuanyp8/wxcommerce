package pmsmodel

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
)

// PmsBrand 品牌表
type PmsBrand struct {
	gorm.Model

	// 表字段对应的结构体属性
	Name                string `gorm:"column:name;not null;size:64;comment:'品牌名称'"`                  // VARCHAR(64)
	FirstLetter         string `gorm:"column:first_letter;not null;size:8;comment:'首字母'"`            // VARCHAR(8)
	Sort                int    `gorm:"column:sort;not null;comment:'排序'"`                            // INT
	FactoryStatus       int    `gorm:"column:factory_status;not null;comment:'是否为品牌制造商：0->不是；1->是'"` // INT(1)
	ShowStatus          int    `gorm:"column:show_status;not null;comment:'显示状态'"`                   // INT(1)
	ProductCount        int    `gorm:"column:product_count;not null;comment:'产品数量'"`                 // INT
	ProductCommentCount int    `gorm:"column:product_comment_count;not null;comment:'产品评论数量'"`       // INT
	Logo                string `gorm:"column:logo;not null;size:255;comment:'品牌logo'"`               // VARCHAR(255)
	BigPic              string `gorm:"column:big_pic;not null;size:255;comment:'专区大图'"`              // VARCHAR(255)
	BrandStory          string `gorm:"column:brand_story;text;not null;comment:'品牌故事'"`              // TEXT
}

// TableName 设置表名
func (*PmsBrand) TableName() string {
	return constants.TableNamePmsBrand
}

func (m *PmsBrand) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if m.ID != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(constants.PmsBrandSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	m.ID = uint(sf.Generate().Int64())
	return nil
}
