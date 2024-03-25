package pack

import (
	"github.com/yuanyp8/wxcommerce/shared/kitex_gen/pms/album"
	"github.com/yuanyp8/wxcommerce/shared/model/model/pmsmodel"
)

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/25 9:34
 * @Desc:
 */

func Album(a *pmsmodel.PmsAlbum) *album.BaseAlbum {
	if a == nil {
		return nil
	}
	return &album.BaseAlbum{
		Id:          uint64(a.ID),
		Name:        a.Name,
		CoverPic:    a.CoverPic,
		PicCount:    int64(a.PicCount),
		Sort:        int64(a.Sort),
		Description: a.Description,
	}
}

func Albums(as []*pmsmodel.PmsAlbum) []*album.BaseAlbum {
	if as == nil {
		return nil
	}
	albums := make([]*album.BaseAlbum, 0, len(as))
	for _, a := range as {
		if a2 := Album(a); a2 != nil {
			albums = append(albums, Album(a))
		}
	}
	return albums
}
