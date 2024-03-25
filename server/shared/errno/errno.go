package errno

import (
	"fmt"
	"github.com/yuanyp8/wxcommerce/shared/kitex_gen/base/errno"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

// NewErrNo return ErrNo
func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success        = NewErrNo(int64(errno.Err_Success), "success")
	NoRoute        = NewErrNo(int64(errno.Err_NoRoute), "no route")
	NoMethod       = NewErrNo(int64(errno.Err_NoMethod), "no method")
	BadRequest     = NewErrNo(int64(errno.Err_BadRequest), "bad request")
	ParamsErr      = NewErrNo(int64(errno.Err_ParamsErr), "params error")
	AuthorizeFail  = NewErrNo(int64(errno.Err_AuthorizeFail), "authorize failed")
	TooManyReqeust = NewErrNo(int64(errno.Err_TooManyRequest), "too many requests")
	ServiceErr     = NewErrNo(int64(errno.Err_ServiceErr), "service error")

	RecordNotFound     = NewErrNo(int64(errno.Err_RecordNotFound), "record not found")
	RecordAlreadyExist = NewErrNo(int64(errno.Err_RecordAlreadyExist), "record already exist")
	DirtyData          = NewErrNo(int64(errno.Err_DirtyData), "dirty data")

	RPCPmsAlbumSrvErr = NewErrNo(int64(errno.Err_RPCPmsAlbumSrvErr), "rpc pms album service error")
	PmsAlbumSrvErr    = NewErrNo(int64(errno.Err_PmsAlbumSrvErr), "pms album service error")

	RPCPmsAlbumPicSrvErr = NewErrNo(int64(errno.Err_RPCPmsAlbumPicSrvErr), "rpc pms album pic service error")
	PmsAlbumPicSrvErr    = NewErrNo(int64(errno.Err_PmsAlbumPicSrvErr), "pms album pic service error")

	RPCPmsBrandSrvErr = NewErrNo(int64(errno.Err_RPCPmsBrandSrvErr), "rpc pms brand service error")
	PmsBrandSrvErr    = NewErrNo(int64(errno.Err_PmsBrandSrvErr), "pms brand service error")

	RPCPmsCommentSrvErr = NewErrNo(int64(errno.Err_RPCPmsCommentSrvErr), "rpc pms comment service error")
	PmsCommentSrvErr    = NewErrNo(int64(errno.Err_PmsCommentSrvErr), "pms comment service error")

	RPCPmsCommentReplySrvErr = NewErrNo(int64(errno.Err_RPCPmsCommentReplySrvErr), "rpc pms comment reply service error")
	PmsCommentReplySrvErr    = NewErrNo(int64(errno.Err_PmsCommentReplySrvErr), "pms comment reply service error")

	RPCPmsMemberPriceSrvErr = NewErrNo(int64(errno.Err_RPCPmsMemberPriceSrvErr), "rpc pms member price service error")
	PmsMemberPriceSrvErr    = NewErrNo(int64(errno.Err_PmsMemberPriceSrvErr), "pms member price service error")

	RPCPmsProductSrvErr = NewErrNo(int64(errno.Err_RPCPmsProductSrvErr), "rpc pms product service error")
	PmsProductSrvErr    = NewErrNo(int64(errno.Err_PmsProductSrvErr), "pms product service error")

	RPCPmsProductAttrSrvErr = NewErrNo(int64(errno.Err_RPCPmsProductAttrSrvErr), "rpc pms product attr service error")
	PmsProductAttrSrvErr    = NewErrNo(int64(errno.Err_PmsProductAttrSrvErr), "pms product attr service error")

	RPCPmsProductAttributeCategorySrvErr = NewErrNo(int64(errno.Err_RPCPmsProductAttributeCategorySrvErr), "rpc pms product attribute category service error")
	PmsProductAttributeCategorySrvErr    = NewErrNo(int64(errno.Err_PmsProductAttributeCategorySrvErr), "pms product attribute category service error")

	RPCPmsProductAttributeValueSrvErr = NewErrNo(int64(errno.Err_RPCPmsProductAttributeValueSrvErr), "rpc pms product attribute value service error")
	PmsProductAttributeValueSrvErr    = NewErrNo(int64(errno.Err_PmsProductAttributeValueSrvErr), "pms product attribute value service error")

	RPCPmsProductCategorySrvErr = NewErrNo(int64(errno.Err_RPCPmsProductCategorySrvErr), "rpc pms product category service error")
	PmsProductCategorySrvErr    = NewErrNo(int64(errno.Err_PmsProductCategorySrvErr), "pms product category service error")

	RPCPmsProductCategoryAttributeRelationSrvErr = NewErrNo(int64(errno.Err_RPCPmsProductCategoryAttributeRelationSrvErr), "rpc pms product category attribute relation service error")
	PmsProductCategoryAttributeRelationSrvErr    = NewErrNo(int64(errno.Err_PmsProductCategoryAttributeRelationSrvErr), "pms product category attribute relation service error")

	RPCPmsProductFullReductionSrvErr = NewErrNo(int64(errno.Err_RPCPmsProductFullReductionSrvErr), "rpc pms product full reduction service error")
	PmsProductFullReductionSrvErr    = NewErrNo(int64(errno.Err_PmsProductFullReductionSrvErr), "pms product full reduction service error")

	RPCPmsProductLadderSrvErr = NewErrNo(int64(errno.Err_RPCPmsProductLadderSrvErr), "rpc pms product ladder service error")
	PmsProductLadderSrvErr    = NewErrNo(int64(errno.Err_PmsProductLadderSrvErr), "pms product ladder service error")

	RPCPmsProductVerifyRecordSrvErr = NewErrNo(int64(errno.Err_RPCPmsProductVerifyRecordSrvErr), "rpc pms product verify record service error")
	PmsProductVerifyRecordSrvErr    = NewErrNo(int64(errno.Err_PmsProductVerifyRecordSrvErr), "pms product verify record service error")

	RPCPmsSkuStockSrvErr = NewErrNo(int64(errno.Err_RPCPmsSkuStockSrvErr), "rpc pms sku stock service error")
	PmsSkuStockSrvErr    = NewErrNo(int64(errno.Err_PmsSkuStockSrvErr), "pms sku stock service error")
)
