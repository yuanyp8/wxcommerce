// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: pms/album.proto

package album

import (
	context "context"
	response "github.com/yuanyp8/wxcommerce/shared/kitex_gen/base/response"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`               // 相册名称
	CoverPic    string `protobuf:"bytes,3,opt,name=CoverPic,proto3" json:"CoverPic,omitempty"`       // 相册封面图片的地址
	PicCount    int64  `protobuf:"varint,4,opt,name=PicCount,proto3" json:"PicCount,omitempty"`      // 相册中图片的数量
	Sort        int64  `protobuf:"varint,5,opt,name=Sort,proto3" json:"Sort,omitempty"`              // 相册的排序编号
	Description string `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"` // 相册的描述信息
}

func (x *Album) Reset() {
	*x = Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pms_album_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_pms_album_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Album.ProtoReflect.Descriptor instead.
func (*Album) Descriptor() ([]byte, []int) {
	return file_pms_album_proto_rawDescGZIP(), []int{0}
}

func (x *Album) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Album) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Album) GetCoverPic() string {
	if x != nil {
		return x.CoverPic
	}
	return ""
}

func (x *Album) GetPicCount() int64 {
	if x != nil {
		return x.PicCount
	}
	return 0
}

func (x *Album) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Album) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AlbumAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *Album `protobuf:"bytes,1,opt,name=Album,proto3" json:"Album,omitempty"`
}

func (x *AlbumAddReq) Reset() {
	*x = AlbumAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pms_album_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumAddReq) ProtoMessage() {}

func (x *AlbumAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_pms_album_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumAddReq.ProtoReflect.Descriptor instead.
func (*AlbumAddReq) Descriptor() ([]byte, []int) {
	return file_pms_album_proto_rawDescGZIP(), []int{1}
}

func (x *AlbumAddReq) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

type AlbumAddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *response.BaseResponse `protobuf:"bytes,1,opt,name=BaseResponse,proto3" json:"BaseResponse,omitempty"`
}

func (x *AlbumAddResp) Reset() {
	*x = AlbumAddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pms_album_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumAddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumAddResp) ProtoMessage() {}

func (x *AlbumAddResp) ProtoReflect() protoreflect.Message {
	mi := &file_pms_album_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumAddResp.ProtoReflect.Descriptor instead.
func (*AlbumAddResp) Descriptor() ([]byte, []int) {
	return file_pms_album_proto_rawDescGZIP(), []int{2}
}

func (x *AlbumAddResp) GetBaseResponse() *response.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

type AlbumListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current  int64 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`   // 当前页号
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"` // 每页展示几个内容
}

func (x *AlbumListReq) Reset() {
	*x = AlbumListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pms_album_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumListReq) ProtoMessage() {}

func (x *AlbumListReq) ProtoReflect() protoreflect.Message {
	mi := &file_pms_album_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumListReq.ProtoReflect.Descriptor instead.
func (*AlbumListReq) Descriptor() ([]byte, []int) {
	return file_pms_album_proto_rawDescGZIP(), []int{3}
}

func (x *AlbumListReq) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *AlbumListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type AlbumListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total        int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List         []*Album               `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	BaseResponse *response.BaseResponse `protobuf:"bytes,3,opt,name=BaseResponse,proto3" json:"BaseResponse,omitempty"`
}

func (x *AlbumListResp) Reset() {
	*x = AlbumListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pms_album_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumListResp) ProtoMessage() {}

func (x *AlbumListResp) ProtoReflect() protoreflect.Message {
	mi := &file_pms_album_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumListResp.ProtoReflect.Descriptor instead.
func (*AlbumListResp) Descriptor() ([]byte, []int) {
	return file_pms_album_proto_rawDescGZIP(), []int{4}
}

func (x *AlbumListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *AlbumListResp) GetList() []*Album {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *AlbumListResp) GetBaseResponse() *response.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

type AlbumUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *Album `protobuf:"bytes,1,opt,name=Album,proto3" json:"Album,omitempty"`
}

func (x *AlbumUpdateReq) Reset() {
	*x = AlbumUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pms_album_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumUpdateReq) ProtoMessage() {}

func (x *AlbumUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_pms_album_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumUpdateReq.ProtoReflect.Descriptor instead.
func (*AlbumUpdateReq) Descriptor() ([]byte, []int) {
	return file_pms_album_proto_rawDescGZIP(), []int{5}
}

func (x *AlbumUpdateReq) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

type AlbumUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *response.BaseResponse `protobuf:"bytes,1,opt,name=BaseResponse,proto3" json:"BaseResponse,omitempty"`
}

func (x *AlbumUpdateResp) Reset() {
	*x = AlbumUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pms_album_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumUpdateResp) ProtoMessage() {}

func (x *AlbumUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_pms_album_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumUpdateResp.ProtoReflect.Descriptor instead.
func (*AlbumUpdateResp) Descriptor() ([]byte, []int) {
	return file_pms_album_proto_rawDescGZIP(), []int{6}
}

func (x *AlbumUpdateResp) GetBaseResponse() *response.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

type AlbumDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *AlbumDeleteReq) Reset() {
	*x = AlbumDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pms_album_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumDeleteReq) ProtoMessage() {}

func (x *AlbumDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_pms_album_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumDeleteReq.ProtoReflect.Descriptor instead.
func (*AlbumDeleteReq) Descriptor() ([]byte, []int) {
	return file_pms_album_proto_rawDescGZIP(), []int{7}
}

func (x *AlbumDeleteReq) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type AlbumDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *response.BaseResponse `protobuf:"bytes,1,opt,name=BaseResponse,proto3" json:"BaseResponse,omitempty"`
}

func (x *AlbumDeleteResp) Reset() {
	*x = AlbumDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pms_album_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumDeleteResp) ProtoMessage() {}

func (x *AlbumDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_pms_album_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumDeleteResp.ProtoReflect.Descriptor instead.
func (*AlbumDeleteResp) Descriptor() ([]byte, []int) {
	return file_pms_album_proto_rawDescGZIP(), []int{8}
}

func (x *AlbumDeleteResp) GetBaseResponse() *response.BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

var File_pms_album_proto protoreflect.FileDescriptor

var file_pms_album_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6d, 0x73, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x70, 0x6d, 0x73, 0x1a, 0x13, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x05,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x50, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x50, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x69, 0x63, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x69, 0x63, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2f, 0x0a, 0x0b, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x22, 0x4a, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x0d, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32,
	0x0a, 0x0e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x20, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x22, 0x4d, 0x0a, 0x0f, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x22, 0x0a, 0x0e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x4d, 0x0a, 0x0f, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe8, 0x01, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x41, 0x64,
	0x64, 0x12, 0x10, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x09, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0b, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x6d, 0x73, 0x2e,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x4d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x6d, 0x73, 0x2e,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75,
	0x61, 0x6e, 0x79, 0x70, 0x38, 0x2f, 0x77, 0x78, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x6d, 0x73, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pms_album_proto_rawDescOnce sync.Once
	file_pms_album_proto_rawDescData = file_pms_album_proto_rawDesc
)

func file_pms_album_proto_rawDescGZIP() []byte {
	file_pms_album_proto_rawDescOnce.Do(func() {
		file_pms_album_proto_rawDescData = protoimpl.X.CompressGZIP(file_pms_album_proto_rawDescData)
	})
	return file_pms_album_proto_rawDescData
}

var file_pms_album_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pms_album_proto_goTypes = []interface{}{
	(*Album)(nil),                 // 0: pms.Album
	(*AlbumAddReq)(nil),           // 1: pms.AlbumAddReq
	(*AlbumAddResp)(nil),          // 2: pms.AlbumAddResp
	(*AlbumListReq)(nil),          // 3: pms.AlbumListReq
	(*AlbumListResp)(nil),         // 4: pms.AlbumListResp
	(*AlbumUpdateReq)(nil),        // 5: pms.AlbumUpdateReq
	(*AlbumUpdateResp)(nil),       // 6: pms.AlbumUpdateResp
	(*AlbumDeleteReq)(nil),        // 7: pms.AlbumDeleteReq
	(*AlbumDeleteResp)(nil),       // 8: pms.AlbumDeleteResp
	(*response.BaseResponse)(nil), // 9: response.BaseResponse
}
var file_pms_album_proto_depIdxs = []int32{
	0,  // 0: pms.AlbumAddReq.Album:type_name -> pms.Album
	9,  // 1: pms.AlbumAddResp.BaseResponse:type_name -> response.BaseResponse
	0,  // 2: pms.AlbumListResp.list:type_name -> pms.Album
	9,  // 3: pms.AlbumListResp.BaseResponse:type_name -> response.BaseResponse
	0,  // 4: pms.AlbumUpdateReq.Album:type_name -> pms.Album
	9,  // 5: pms.AlbumUpdateResp.BaseResponse:type_name -> response.BaseResponse
	9,  // 6: pms.AlbumDeleteResp.BaseResponse:type_name -> response.BaseResponse
	1,  // 7: pms.AlbumService.AlbumAdd:input_type -> pms.AlbumAddReq
	3,  // 8: pms.AlbumService.AlbumList:input_type -> pms.AlbumListReq
	5,  // 9: pms.AlbumService.AlbumUpdate:input_type -> pms.AlbumUpdateReq
	7,  // 10: pms.AlbumService.AlbumMDelete:input_type -> pms.AlbumDeleteReq
	2,  // 11: pms.AlbumService.AlbumAdd:output_type -> pms.AlbumAddResp
	4,  // 12: pms.AlbumService.AlbumList:output_type -> pms.AlbumListResp
	6,  // 13: pms.AlbumService.AlbumUpdate:output_type -> pms.AlbumUpdateResp
	8,  // 14: pms.AlbumService.AlbumMDelete:output_type -> pms.AlbumDeleteResp
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_pms_album_proto_init() }
func file_pms_album_proto_init() {
	if File_pms_album_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pms_album_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Album); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pms_album_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumAddReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pms_album_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumAddResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pms_album_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pms_album_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumListResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pms_album_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumUpdateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pms_album_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumUpdateResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pms_album_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumDeleteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pms_album_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumDeleteResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pms_album_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pms_album_proto_goTypes,
		DependencyIndexes: file_pms_album_proto_depIdxs,
		MessageInfos:      file_pms_album_proto_msgTypes,
	}.Build()
	File_pms_album_proto = out.File
	file_pms_album_proto_rawDesc = nil
	file_pms_album_proto_goTypes = nil
	file_pms_album_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.0. DO NOT EDIT.

type AlbumService interface {
	AlbumAdd(ctx context.Context, req *AlbumAddReq) (res *AlbumAddResp, err error)
	AlbumList(ctx context.Context, req *AlbumListReq) (res *AlbumListResp, err error)
	AlbumUpdate(ctx context.Context, req *AlbumUpdateReq) (res *AlbumUpdateResp, err error)
	AlbumMDelete(ctx context.Context, req *AlbumDeleteReq) (res *AlbumDeleteResp, err error)
}
