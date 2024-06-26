// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package album_pic

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *AlbumPicAddReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlbumPicAddReq[number], err)
}

func (x *AlbumPicAddReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *AlbumPicAddReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.AlbumId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *AlbumPicAddReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Pic, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlbumPicAddResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlbumPicAddResp[number], err)
}

func (x *AlbumPicAddResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Pong, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlbumPicListReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlbumPicListReq[number], err)
}

func (x *AlbumPicListReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Current, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AlbumPicListReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AlbumPicListData) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlbumPicListData[number], err)
}

func (x *AlbumPicListData) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *AlbumPicListData) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.AlbumId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *AlbumPicListData) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Pic, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlbumPicListResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlbumPicListResp[number], err)
}

func (x *AlbumPicListResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AlbumPicListResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v AlbumPicListData
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.List = append(x.List, &v)
	return offset, nil
}

func (x *AlbumPicUpdateReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlbumPicUpdateReq[number], err)
}

func (x *AlbumPicUpdateReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *AlbumPicUpdateReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.AlbumId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *AlbumPicUpdateReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Pic, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlbumPicUpdateResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlbumPicUpdateResp[number], err)
}

func (x *AlbumPicUpdateResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Pong, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlbumPicDeleteReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlbumPicDeleteReq[number], err)
}

func (x *AlbumPicDeleteReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v uint64
			v, offset, err = fastpb.ReadUint64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.Ids = append(x.Ids, v)
			return offset, err
		})
	return offset, err
}

func (x *AlbumPicDeleteResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlbumPicDeleteResp[number], err)
}

func (x *AlbumPicDeleteResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Pong, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlbumPicAddReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *AlbumPicAddReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *AlbumPicAddReq) fastWriteField2(buf []byte) (offset int) {
	if x.AlbumId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetAlbumId())
	return offset
}

func (x *AlbumPicAddReq) fastWriteField3(buf []byte) (offset int) {
	if x.Pic == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPic())
	return offset
}

func (x *AlbumPicAddResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *AlbumPicAddResp) fastWriteField1(buf []byte) (offset int) {
	if x.Pong == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPong())
	return offset
}

func (x *AlbumPicListReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *AlbumPicListReq) fastWriteField1(buf []byte) (offset int) {
	if x.Current == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetCurrent())
	return offset
}

func (x *AlbumPicListReq) fastWriteField2(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetPageSize())
	return offset
}

func (x *AlbumPicListData) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *AlbumPicListData) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *AlbumPicListData) fastWriteField2(buf []byte) (offset int) {
	if x.AlbumId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetAlbumId())
	return offset
}

func (x *AlbumPicListData) fastWriteField3(buf []byte) (offset int) {
	if x.Pic == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPic())
	return offset
}

func (x *AlbumPicListResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *AlbumPicListResp) fastWriteField1(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetTotal())
	return offset
}

func (x *AlbumPicListResp) fastWriteField2(buf []byte) (offset int) {
	if x.List == nil {
		return offset
	}
	for i := range x.GetList() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetList()[i])
	}
	return offset
}

func (x *AlbumPicUpdateReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *AlbumPicUpdateReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *AlbumPicUpdateReq) fastWriteField2(buf []byte) (offset int) {
	if x.AlbumId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetAlbumId())
	return offset
}

func (x *AlbumPicUpdateReq) fastWriteField3(buf []byte) (offset int) {
	if x.Pic == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPic())
	return offset
}

func (x *AlbumPicUpdateResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *AlbumPicUpdateResp) fastWriteField1(buf []byte) (offset int) {
	if x.Pong == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPong())
	return offset
}

func (x *AlbumPicDeleteReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *AlbumPicDeleteReq) fastWriteField1(buf []byte) (offset int) {
	if len(x.Ids) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.GetIds()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteUint64(buf[offset:], numTagOrKey, x.GetIds()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *AlbumPicDeleteResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *AlbumPicDeleteResp) fastWriteField1(buf []byte) (offset int) {
	if x.Pong == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPong())
	return offset
}

func (x *AlbumPicAddReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *AlbumPicAddReq) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetId())
	return n
}

func (x *AlbumPicAddReq) sizeField2() (n int) {
	if x.AlbumId == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetAlbumId())
	return n
}

func (x *AlbumPicAddReq) sizeField3() (n int) {
	if x.Pic == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPic())
	return n
}

func (x *AlbumPicAddResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *AlbumPicAddResp) sizeField1() (n int) {
	if x.Pong == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPong())
	return n
}

func (x *AlbumPicListReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *AlbumPicListReq) sizeField1() (n int) {
	if x.Current == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetCurrent())
	return n
}

func (x *AlbumPicListReq) sizeField2() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetPageSize())
	return n
}

func (x *AlbumPicListData) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *AlbumPicListData) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetId())
	return n
}

func (x *AlbumPicListData) sizeField2() (n int) {
	if x.AlbumId == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetAlbumId())
	return n
}

func (x *AlbumPicListData) sizeField3() (n int) {
	if x.Pic == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPic())
	return n
}

func (x *AlbumPicListResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *AlbumPicListResp) sizeField1() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetTotal())
	return n
}

func (x *AlbumPicListResp) sizeField2() (n int) {
	if x.List == nil {
		return n
	}
	for i := range x.GetList() {
		n += fastpb.SizeMessage(2, x.GetList()[i])
	}
	return n
}

func (x *AlbumPicUpdateReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *AlbumPicUpdateReq) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetId())
	return n
}

func (x *AlbumPicUpdateReq) sizeField2() (n int) {
	if x.AlbumId == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetAlbumId())
	return n
}

func (x *AlbumPicUpdateReq) sizeField3() (n int) {
	if x.Pic == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPic())
	return n
}

func (x *AlbumPicUpdateResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *AlbumPicUpdateResp) sizeField1() (n int) {
	if x.Pong == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPong())
	return n
}

func (x *AlbumPicDeleteReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *AlbumPicDeleteReq) sizeField1() (n int) {
	if len(x.Ids) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.GetIds()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeUint64(numTagOrKey, x.GetIds()[numIdxOrVal])
			return n
		})
	return n
}

func (x *AlbumPicDeleteResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *AlbumPicDeleteResp) sizeField1() (n int) {
	if x.Pong == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPong())
	return n
}

var fieldIDToName_AlbumPicAddReq = map[int32]string{
	1: "Id",
	2: "AlbumId",
	3: "Pic",
}

var fieldIDToName_AlbumPicAddResp = map[int32]string{
	1: "Pong",
}

var fieldIDToName_AlbumPicListReq = map[int32]string{
	1: "Current",
	2: "PageSize",
}

var fieldIDToName_AlbumPicListData = map[int32]string{
	1: "Id",
	2: "AlbumId",
	3: "Pic",
}

var fieldIDToName_AlbumPicListResp = map[int32]string{
	1: "Total",
	2: "List",
}

var fieldIDToName_AlbumPicUpdateReq = map[int32]string{
	1: "Id",
	2: "AlbumId",
	3: "Pic",
}

var fieldIDToName_AlbumPicUpdateResp = map[int32]string{
	1: "Pong",
}

var fieldIDToName_AlbumPicDeleteReq = map[int32]string{
	1: "Ids",
}

var fieldIDToName_AlbumPicDeleteResp = map[int32]string{
	1: "Pong",
}
