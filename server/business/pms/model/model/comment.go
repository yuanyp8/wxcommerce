/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:21
 * @Desc:
 */

package model

import (
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
	"time"
)

// PmsComment 商品评价表
type PmsComment struct {
	gorm.Model

	// 表字段对应的结构体属性
	ProductId        int64     `gorm:"column:product_id;not null;comment:'商品ID'"`                     // BIGINT
	MemberNickName   string    `gorm:"column:member_nick_name;not null;size:255;comment:'会员昵称'"`      // VARCHAR(255)
	ProductName      string    `gorm:"column:product_name;not null;size:255;comment:'商品名称'"`          // VARCHAR(255)
	Star             int       `gorm:"column:star;not null;comment:'评价星数：0->5'"`                      // INT(3)
	MemberIp         string    `gorm:"column:member_ip;not null;size:64;comment:'评价的ip'"`             // VARCHAR(64)
	CreateTime       time.Time `gorm:"column:create_time;not null;comment:'创建时间'"`                    // DATETIME
	ShowStatus       int       `gorm:"column:show_status;not null;comment:'显示状态'"`                    // INT(1)
	ProductAttribute string    `gorm:"column:product_attribute;not null;size:255;comment:'购买时的商品属性'"` // VARCHAR(255)
	CollectCount     int       `gorm:"column:collect_count;not null;comment:'收藏数量'"`                  // INT
	ReadCount        int       `gorm:"column:read_count;not null;comment:'阅读数量'"`                     // INT
	Content          string    `gorm:"column:content;text;not null;comment:'评价内容'"`                   // TEXT
	Pics             string    `gorm:"column:pics;not null;size:1000;comment:'上传图片地址，以逗号隔开'"`         // VARCHAR(1000)
	MemberIcon       string    `gorm:"column:member_icon;not null;size:255;comment:'评论用户头像'"`         // VARCHAR(255)
	ReplyCount       int       `gorm:"column:reply_count;not null;comment:'回复数量'"`                    // INT
}

// TableName 设置表名
func (PmsComment) TableName() string {
	return constants.TableNamePmsComment
}
