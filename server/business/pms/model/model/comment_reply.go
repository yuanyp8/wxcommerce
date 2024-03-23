/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 4:22
 * @Desc:
 */

package model

import (
	"github.com/yuanyp8/wxcommerce/shared/constants"
	"gorm.io/gorm"
	"time"
)

// PmsCommentReplay 产品评价回复表
type PmsCommentReplay struct {
	gorm.Model

	// 表字段对应的结构体属性
	CommentID      int64     `gorm:"column:comment_id;not null;"`
	MemberNickName string    `gorm:"column:member_nick_name;not null;size:255;"`
	MemberIcon     string    `gorm:"column:member_icon;not null;size:255;"`
	Content        string    `gorm:"column:content;not null;size:1000;"`
	CreateTime     time.Time `gorm:"column:create_time;not null;"`
	Type           int       `gorm:"column:type;not null;comment:'评论人员类型；0->会员；1->管理员';"`
}

// TableName 设置表名
func (PmsCommentReplay) TableName() string {
	return constants.TableNamePmsCommentReply
}
