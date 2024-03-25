package mysql

import (
	"context"
	"github.com/yuanyp8/wxcommerce/shared/errno"
	"github.com/yuanyp8/wxcommerce/shared/model/model/pmsmodel"
	"gorm.io/gorm"
)

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/24 6:27
 * @Desc:
 */

type AlbumManager struct {
	salt string
	db   *gorm.DB
}

// NewAlbumManager creates a mysql manager.
func NewAlbumManager(db *gorm.DB, salt string) *AlbumManager {
	m := db.Migrator()
	if !m.HasTable(&pmsmodel.PmsAlbum{}) {
		if err := m.CreateTable(&pmsmodel.PmsAlbum{}); err != nil {
			panic(err)
		}
	}
	return &AlbumManager{
		db:   db,
		salt: salt,
	}
}

func (m *AlbumManager) AlbumAdd(ctx context.Context, album *pmsmodel.PmsAlbum) error {
	if album.ID == 0 || album.Name == "" || album.CoverPic == "" || album.PicCount == 0 {
		return errno.PmsAlbumPicSrvErr.WithMessage("params must not null")
	}

	// 查询是否存在重复的记录, 成功查询到记录则存在图片名称记录，否则不影响
	//if err := m.db.WithContext(ctx).Where("name = ?", album.Name).First(&pmsmodel.PmsAlbum{}).Error; err == nil {
	//	return errno.PmsAlbumPicSrvErr.WithMessage("album name already exist")
	//} else if err != errno.RecordNotFound {
	//		return errno.PmsAlbumSrvErr.WithMessage(err.Error())
	//	}

	if err := m.db.WithContext(ctx).Create(album).Error; err != nil {
		return errno.RecordAlreadyExist
	}
	return nil
}

func (m *AlbumManager) AlbumList(ctx context.Context, pageNum, pageSize int64) ([]*pmsmodel.PmsAlbum, int64, error) {
	albums := make([]*pmsmodel.PmsAlbum, 0)
	offset := (pageNum - 1) * pageSize

	if err := m.db.WithContext(ctx).Limit(int(pageSize)).Offset(int(offset)).Limit(int(pageSize)).Order("sort").Find(&albums).Error; err != nil {
		if err == gorm.ErrEmptySlice {
			return nil, 0, errno.RecordNotFound
		}
		return nil, 0, err
	}
	count, err := m.Count(ctx, albums[0])
	if err != nil {
		return nil, 0, err
	}

	return albums, count, nil
}

func (m *AlbumManager) Count(ctx context.Context, album *pmsmodel.PmsAlbum) (int64, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(album).Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return 0, errno.RecordNotFound
	} else if err != nil {
		return 0, err
	}

	return count, nil
}

func (m *AlbumManager) AlbumUpdate(ctx context.Context, album *pmsmodel.PmsAlbum) error {
	err := m.db.WithContext(ctx).Model(&pmsmodel.PmsAlbum{Model: gorm.Model{ID: album.ID}}).Updates(album).Error

	if err == gorm.ErrRecordNotFound {
		return errno.RecordNotFound
	}
	return err
}

func (m *AlbumManager) AlbumMDelete(ctx context.Context, ids []interface{}) error {
	err := m.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&pmsmodel.PmsAlbum{}).Error
	return err
}
