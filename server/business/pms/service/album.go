package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yuanyp8/wxcommerce/business/pms/dal/mysql"
	"github.com/yuanyp8/wxcommerce/business/pms/pack"
	"github.com/yuanyp8/wxcommerce/shared/errno"
	"github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/album"
	"github.com/yuanyp8/wxcommerce/shared/model/model/pmsmodel"
)

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/25 10:20
 * @Desc:
 */

// AlbumServiceImpl implements the last service interface defined in the IDL.
type AlbumServiceImpl struct{
	AlbumMysqlManager
}

type AlbumMysqlManager interface {
	AlbumAdd(ctx context.Context, album *pmsmodel.PmsAlbum) error
	AlbumList(ctx context.Context, pageNum, pageSize int64) ([]*pmsmodel.PmsAlbum, int64, error)
	AlbumUpdate(ctx context.Context, album *pmsmodel.PmsAlbum) error
	AlbumMDelete(ctx context.Context, ids []interface{})
}

// AlbumAdd implements the AlbumServiceImpl interface.
func (s *AlbumServiceImpl) AlbumAdd(ctx context.Context, req *album.AlbumAddReq) (resp *album.AlbumAddResp, err error) {
	resp = new(album.AlbumAddResp)

	if req.Album.Name == "" || req.Album.CoverPic == "" || req.Album.PicCount == 0 {
		resp.BaseResponse = pack.BuildBaseResp(errno.ParamsErr)
		return resp, nil
	}

	album := &pmsmodel.PmsAlbum{
		Name:        req.Album.Name,
		CoverPic:    req.Album.CoverPic,
		PicCount:    int64(req.Album.PicCount),
		Description: req.Album.Description,
		Sort:        req.Album.Sort,
		Status:      req.Album.Status,
	}


	err := s.AlbumMysqlManager.AlbumAdd()

	_, err = s..CreateUser(&mysql.User{
		ID:           req.AccountId,
		PhoneNumber:  req.PhoneNumber,
		AvatarBlobId: req.AvatarBlobId,
		Username:     req.Username,
		OpenID:       req.OpenId,
	})
	if err != nil {
		if err == errno.RecordAlreadyExist {
			klog.Error("add user error", err)
			resp.BaseResp = tools.BuildBaseResp(errno.RecordAlreadyExist)
			return resp, nil
		}
		klog.Error("add user error", err)
		resp.BaseResp = tools.BuildBaseResp(errno.UserSrvErr)
		return resp, nil
	}
	resp.BaseResp = tools.BuildBaseResp(nil)
	return resp, nil
}

	return
}

// AlbumList implements the AlbumServiceImpl interface.
func (s *AlbumServiceImpl) AlbumList(ctx context.Context, req *album.AlbumListReq) (resp *album.AlbumListResp, err error) {
	// TODO: Your code here...
	return
}

// AlbumUpdate implements the AlbumServiceImpl interface.
func (s *AlbumServiceImpl) AlbumUpdate(ctx context.Context, req *album.AlbumUpdateReq) (resp *album.AlbumUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// AlbumDelete implements the AlbumServiceImpl interface.
func (s *AlbumServiceImpl) AlbumDelete(ctx context.Context, req *album.AlbumDeleteReq) (resp *album.AlbumDeleteResp, err error) {
	// TODO: Your code here...
	return
}
