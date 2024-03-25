package pack

import (
	"errors"
	"github.com/yuanyp8/wxcommerce/shared/errno"
	"github.com/yuanyp8/wxcommerce/shared/kitex_gen/base/response"
	"time"
)

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/25 11:12
 * @Desc:
 */
func baseResp(err errno.ErrNo) *response.BaseResponse {
	return &response.BaseResponse{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}

func BuildBaseResp(err error) *response.BaseResponse {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}
