// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package product_full_reduction

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *BaseProductOperateLog) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 11:
		offset, err = x.fastReadField11(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 12:
		offset, err = x.fastReadField12(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseProductOperateLog[number], err)
}

func (x *BaseProductOperateLog) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ProductId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PriceOld, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.PriceNew, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.SalePriceOld, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.SalePriceNew, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.GiftPointOld, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.GiftPointNew, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.UsePointLimitOld, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	x.UsePointLimitNew, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	x.OperateMan, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) fastReadField12(buf []byte, _type int8) (offset int, err error) {
	x.CreateTime, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ProductOperateLogAddReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductOperateLogAddReq[number], err)
}

func (x *ProductOperateLogAddReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseProductOperateLog
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.ProductOperateLog = &v
	return offset, nil
}

func (x *ProductOperateLogAddResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductOperateLogAddResp[number], err)
}

func (x *ProductOperateLogAddResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Pong, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ProductOperateLogListReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductOperateLogListReq[number], err)
}

func (x *ProductOperateLogListReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Current, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ProductOperateLogListReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ProductOperateLogListResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductOperateLogListResp[number], err)
}

func (x *ProductOperateLogListResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ProductOperateLogListResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v BaseProductOperateLog
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.List = append(x.List, &v)
	return offset, nil
}

func (x *ProductOperateLogUpdateReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductOperateLogUpdateReq[number], err)
}

func (x *ProductOperateLogUpdateReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseProductOperateLog
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.ProductOperateLog = &v
	return offset, nil
}

func (x *ProductOperateLogUpdateResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductOperateLogUpdateResp[number], err)
}

func (x *ProductOperateLogUpdateResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Pong, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ProductOperateLogDeleteReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductOperateLogDeleteReq[number], err)
}

func (x *ProductOperateLogDeleteReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *ProductOperateLogDeleteResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductOperateLogDeleteResp[number], err)
}

func (x *ProductOperateLogDeleteResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Pong, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseProductOperateLog) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	offset += x.fastWriteField12(buf[offset:])
	return offset
}

func (x *BaseProductOperateLog) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField2(buf []byte) (offset int) {
	if x.ProductId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetProductId())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField3(buf []byte) (offset int) {
	if x.PriceOld == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetPriceOld())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField4(buf []byte) (offset int) {
	if x.PriceNew == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetPriceNew())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField5(buf []byte) (offset int) {
	if x.SalePriceOld == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetSalePriceOld())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField6(buf []byte) (offset int) {
	if x.SalePriceNew == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.GetSalePriceNew())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField7(buf []byte) (offset int) {
	if x.GiftPointOld == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.GetGiftPointOld())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField8(buf []byte) (offset int) {
	if x.GiftPointNew == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 8, x.GetGiftPointNew())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField9(buf []byte) (offset int) {
	if x.UsePointLimitOld == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 9, x.GetUsePointLimitOld())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField10(buf []byte) (offset int) {
	if x.UsePointLimitNew == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 10, x.GetUsePointLimitNew())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField11(buf []byte) (offset int) {
	if x.OperateMan == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 11, x.GetOperateMan())
	return offset
}

func (x *BaseProductOperateLog) fastWriteField12(buf []byte) (offset int) {
	if x.CreateTime == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 12, x.GetCreateTime())
	return offset
}

func (x *ProductOperateLogAddReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ProductOperateLogAddReq) fastWriteField1(buf []byte) (offset int) {
	if x.ProductOperateLog == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetProductOperateLog())
	return offset
}

func (x *ProductOperateLogAddResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ProductOperateLogAddResp) fastWriteField1(buf []byte) (offset int) {
	if x.Pong == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPong())
	return offset
}

func (x *ProductOperateLogListReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ProductOperateLogListReq) fastWriteField1(buf []byte) (offset int) {
	if x.Current == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetCurrent())
	return offset
}

func (x *ProductOperateLogListReq) fastWriteField2(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetPageSize())
	return offset
}

func (x *ProductOperateLogListResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ProductOperateLogListResp) fastWriteField1(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetTotal())
	return offset
}

func (x *ProductOperateLogListResp) fastWriteField2(buf []byte) (offset int) {
	if x.List == nil {
		return offset
	}
	for i := range x.GetList() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetList()[i])
	}
	return offset
}

func (x *ProductOperateLogUpdateReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ProductOperateLogUpdateReq) fastWriteField1(buf []byte) (offset int) {
	if x.ProductOperateLog == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetProductOperateLog())
	return offset
}

func (x *ProductOperateLogUpdateResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ProductOperateLogUpdateResp) fastWriteField1(buf []byte) (offset int) {
	if x.Pong == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPong())
	return offset
}

func (x *ProductOperateLogDeleteReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ProductOperateLogDeleteReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *ProductOperateLogDeleteResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ProductOperateLogDeleteResp) fastWriteField1(buf []byte) (offset int) {
	if x.Pong == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPong())
	return offset
}

func (x *BaseProductOperateLog) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	n += x.sizeField12()
	return n
}

func (x *BaseProductOperateLog) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetId())
	return n
}

func (x *BaseProductOperateLog) sizeField2() (n int) {
	if x.ProductId == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetProductId())
	return n
}

func (x *BaseProductOperateLog) sizeField3() (n int) {
	if x.PriceOld == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetPriceOld())
	return n
}

func (x *BaseProductOperateLog) sizeField4() (n int) {
	if x.PriceNew == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetPriceNew())
	return n
}

func (x *BaseProductOperateLog) sizeField5() (n int) {
	if x.SalePriceOld == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetSalePriceOld())
	return n
}

func (x *BaseProductOperateLog) sizeField6() (n int) {
	if x.SalePriceNew == 0 {
		return n
	}
	n += fastpb.SizeInt64(6, x.GetSalePriceNew())
	return n
}

func (x *BaseProductOperateLog) sizeField7() (n int) {
	if x.GiftPointOld == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.GetGiftPointOld())
	return n
}

func (x *BaseProductOperateLog) sizeField8() (n int) {
	if x.GiftPointNew == 0 {
		return n
	}
	n += fastpb.SizeInt64(8, x.GetGiftPointNew())
	return n
}

func (x *BaseProductOperateLog) sizeField9() (n int) {
	if x.UsePointLimitOld == 0 {
		return n
	}
	n += fastpb.SizeInt64(9, x.GetUsePointLimitOld())
	return n
}

func (x *BaseProductOperateLog) sizeField10() (n int) {
	if x.UsePointLimitNew == 0 {
		return n
	}
	n += fastpb.SizeInt64(10, x.GetUsePointLimitNew())
	return n
}

func (x *BaseProductOperateLog) sizeField11() (n int) {
	if x.OperateMan == "" {
		return n
	}
	n += fastpb.SizeString(11, x.GetOperateMan())
	return n
}

func (x *BaseProductOperateLog) sizeField12() (n int) {
	if x.CreateTime == "" {
		return n
	}
	n += fastpb.SizeString(12, x.GetCreateTime())
	return n
}

func (x *ProductOperateLogAddReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ProductOperateLogAddReq) sizeField1() (n int) {
	if x.ProductOperateLog == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetProductOperateLog())
	return n
}

func (x *ProductOperateLogAddResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ProductOperateLogAddResp) sizeField1() (n int) {
	if x.Pong == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPong())
	return n
}

func (x *ProductOperateLogListReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ProductOperateLogListReq) sizeField1() (n int) {
	if x.Current == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetCurrent())
	return n
}

func (x *ProductOperateLogListReq) sizeField2() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetPageSize())
	return n
}

func (x *ProductOperateLogListResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ProductOperateLogListResp) sizeField1() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetTotal())
	return n
}

func (x *ProductOperateLogListResp) sizeField2() (n int) {
	if x.List == nil {
		return n
	}
	for i := range x.GetList() {
		n += fastpb.SizeMessage(2, x.GetList()[i])
	}
	return n
}

func (x *ProductOperateLogUpdateReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ProductOperateLogUpdateReq) sizeField1() (n int) {
	if x.ProductOperateLog == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetProductOperateLog())
	return n
}

func (x *ProductOperateLogUpdateResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ProductOperateLogUpdateResp) sizeField1() (n int) {
	if x.Pong == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPong())
	return n
}

func (x *ProductOperateLogDeleteReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ProductOperateLogDeleteReq) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetId())
	return n
}

func (x *ProductOperateLogDeleteResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ProductOperateLogDeleteResp) sizeField1() (n int) {
	if x.Pong == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPong())
	return n
}

var fieldIDToName_BaseProductOperateLog = map[int32]string{
	1:  "Id",
	2:  "ProductId",
	3:  "PriceOld",
	4:  "PriceNew",
	5:  "SalePriceOld",
	6:  "SalePriceNew",
	7:  "GiftPointOld",
	8:  "GiftPointNew",
	9:  "UsePointLimitOld",
	10: "UsePointLimitNew",
	11: "OperateMan",
	12: "CreateTime",
}

var fieldIDToName_ProductOperateLogAddReq = map[int32]string{
	1: "ProductOperateLog",
}

var fieldIDToName_ProductOperateLogAddResp = map[int32]string{
	1: "Pong",
}

var fieldIDToName_ProductOperateLogListReq = map[int32]string{
	1: "Current",
	2: "PageSize",
}

var fieldIDToName_ProductOperateLogListResp = map[int32]string{
	1: "Total",
	2: "List",
}

var fieldIDToName_ProductOperateLogUpdateReq = map[int32]string{
	1: "ProductOperateLog",
}

var fieldIDToName_ProductOperateLogUpdateResp = map[int32]string{
	1: "Pong",
}

var fieldIDToName_ProductOperateLogDeleteReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_ProductOperateLogDeleteResp = map[int32]string{
	1: "Pong",
}
