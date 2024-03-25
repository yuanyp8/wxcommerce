package pmsmodel

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
)

// PmsAlbum 相册表
type PmsAlbum struct {
	gorm.Model

	// 表字段对应的结构体属性
	Name        string `gorm:"not null;comment:'相册名称'"`   // varchar(64)
	CoverPic    string `gorm:"not null;comment:'封面图片地址'"` // varchar(1000)
	PicCount    int    `gorm:"not null;comment:'图片数量'"`   // int
	Sort        int    `gorm:"not null;comment:'排序'"`     // int
	Description string `gorm:"not null;comment:'描述'"`     // varchar(1000)
}

// TableName 设置表名
func (*PmsAlbum) TableName() string {
	return constants.TableNamePmsAlbum
}

func (m *PmsAlbum) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if m.ID != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(constants.PmsAlbumSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	m.ID = uint(sf.Generate().Int64())
	return nil
}
