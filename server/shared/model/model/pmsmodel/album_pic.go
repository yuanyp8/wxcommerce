package pmsmodel

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
)

// PmsAlbumPic 画册图片表
type PmsAlbumPic struct {
	gorm.Model

	// 表字段对应的结构体属性
	AlbumId int64  `gorm:"column:album_id;not null;comment:'画册ID'"`      // BIGINT
	Pic     string `gorm:"column:pic;not null;size:1000;comment:'图片地址'"` // VARCHAR(1000)
}

// TableName 设置表名
func (*PmsAlbumPic) TableName() string {
	return constants.TableNamePmsAlbumPic
}

func (m *PmsAlbumPic) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if m.ID != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(constants.PmsAlbumPicSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	m.ID = uint(sf.Generate().Int64())
	return nil
}
