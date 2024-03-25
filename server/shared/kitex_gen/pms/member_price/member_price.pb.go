// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: member_price.proto

package member_price

import (
	context "context"
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

type BaseMemberPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ProductId       uint64 `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`            // 商品id
	MemberLevelId   uint64 `protobuf:"varint,3,opt,name=MemberLevelId,proto3" json:"MemberLevelId,omitempty"`    // 会员等级id
	MemberPrice     int64  `protobuf:"varint,4,opt,name=MemberPrice,proto3" json:"MemberPrice,omitempty"`        // 会员价格
	MemberLevelName string `protobuf:"bytes,5,opt,name=MemberLevelName,proto3" json:"MemberLevelName,omitempty"` // 会员等级名称
}

func (x *BaseMemberPrice) Reset() {
	*x = BaseMemberPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseMemberPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseMemberPrice) ProtoMessage() {}

func (x *BaseMemberPrice) ProtoReflect() protoreflect.Message {
	mi := &file_member_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseMemberPrice.ProtoReflect.Descriptor instead.
func (*BaseMemberPrice) Descriptor() ([]byte, []int) {
	return file_member_price_proto_rawDescGZIP(), []int{0}
}

func (x *BaseMemberPrice) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseMemberPrice) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *BaseMemberPrice) GetMemberLevelId() uint64 {
	if x != nil {
		return x.MemberLevelId
	}
	return 0
}

func (x *BaseMemberPrice) GetMemberPrice() int64 {
	if x != nil {
		return x.MemberPrice
	}
	return 0
}

func (x *BaseMemberPrice) GetMemberLevelName() string {
	if x != nil {
		return x.MemberLevelName
	}
	return ""
}

type MemberPriceAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberPrice *BaseMemberPrice `protobuf:"bytes,1,opt,name=MemberPrice,proto3" json:"MemberPrice,omitempty"`
}

func (x *MemberPriceAddReq) Reset() {
	*x = MemberPriceAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_price_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceAddReq) ProtoMessage() {}

func (x *MemberPriceAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_member_price_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceAddReq.ProtoReflect.Descriptor instead.
func (*MemberPriceAddReq) Descriptor() ([]byte, []int) {
	return file_member_price_proto_rawDescGZIP(), []int{1}
}

func (x *MemberPriceAddReq) GetMemberPrice() *BaseMemberPrice {
	if x != nil {
		return x.MemberPrice
	}
	return nil
}

type MemberPriceAddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *MemberPriceAddResp) Reset() {
	*x = MemberPriceAddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_price_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceAddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceAddResp) ProtoMessage() {}

func (x *MemberPriceAddResp) ProtoReflect() protoreflect.Message {
	mi := &file_member_price_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceAddResp.ProtoReflect.Descriptor instead.
func (*MemberPriceAddResp) Descriptor() ([]byte, []int) {
	return file_member_price_proto_rawDescGZIP(), []int{2}
}

func (x *MemberPriceAddResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type MemberPriceListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current  int64 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *MemberPriceListReq) Reset() {
	*x = MemberPriceListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_price_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceListReq) ProtoMessage() {}

func (x *MemberPriceListReq) ProtoReflect() protoreflect.Message {
	mi := &file_member_price_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceListReq.ProtoReflect.Descriptor instead.
func (*MemberPriceListReq) Descriptor() ([]byte, []int) {
	return file_member_price_proto_rawDescGZIP(), []int{3}
}

func (x *MemberPriceListReq) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *MemberPriceListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type MemberPriceListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64              `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*BaseMemberPrice `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MemberPriceListResp) Reset() {
	*x = MemberPriceListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_price_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceListResp) ProtoMessage() {}

func (x *MemberPriceListResp) ProtoReflect() protoreflect.Message {
	mi := &file_member_price_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceListResp.ProtoReflect.Descriptor instead.
func (*MemberPriceListResp) Descriptor() ([]byte, []int) {
	return file_member_price_proto_rawDescGZIP(), []int{4}
}

func (x *MemberPriceListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MemberPriceListResp) GetList() []*BaseMemberPrice {
	if x != nil {
		return x.List
	}
	return nil
}

type MemberPriceUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberPrice *BaseMemberPrice `protobuf:"bytes,1,opt,name=MemberPrice,proto3" json:"MemberPrice,omitempty"`
}

func (x *MemberPriceUpdateReq) Reset() {
	*x = MemberPriceUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_price_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceUpdateReq) ProtoMessage() {}

func (x *MemberPriceUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_member_price_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceUpdateReq.ProtoReflect.Descriptor instead.
func (*MemberPriceUpdateReq) Descriptor() ([]byte, []int) {
	return file_member_price_proto_rawDescGZIP(), []int{5}
}

func (x *MemberPriceUpdateReq) GetMemberPrice() *BaseMemberPrice {
	if x != nil {
		return x.MemberPrice
	}
	return nil
}

type MemberPriceUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *MemberPriceUpdateResp) Reset() {
	*x = MemberPriceUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_price_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceUpdateResp) ProtoMessage() {}

func (x *MemberPriceUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_member_price_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceUpdateResp.ProtoReflect.Descriptor instead.
func (*MemberPriceUpdateResp) Descriptor() ([]byte, []int) {
	return file_member_price_proto_rawDescGZIP(), []int{6}
}

func (x *MemberPriceUpdateResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type MemberPriceDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *MemberPriceDeleteReq) Reset() {
	*x = MemberPriceDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_price_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceDeleteReq) ProtoMessage() {}

func (x *MemberPriceDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_member_price_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceDeleteReq.ProtoReflect.Descriptor instead.
func (*MemberPriceDeleteReq) Descriptor() ([]byte, []int) {
	return file_member_price_proto_rawDescGZIP(), []int{7}
}

func (x *MemberPriceDeleteReq) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type MemberPriceDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *MemberPriceDeleteResp) Reset() {
	*x = MemberPriceDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_member_price_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceDeleteResp) ProtoMessage() {}

func (x *MemberPriceDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_member_price_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceDeleteResp.ProtoReflect.Descriptor instead.
func (*MemberPriceDeleteResp) Descriptor() ([]byte, []int) {
	return file_member_price_proto_rawDescGZIP(), []int{8}
}

func (x *MemberPriceDeleteResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

var File_member_price_proto protoreflect.FileDescriptor

var file_member_price_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x6d, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x0f, 0x42, 0x61,
	0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a,
	0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x36, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x0b, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x6f, 0x6e, 0x67, 0x22, 0x4a, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x55, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6d,
	0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x36, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x0b, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x2b, 0x0a, 0x15, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x6f, 0x6e, 0x67, 0x22, 0x28, 0x0a, 0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x2b,
	0x0a, 0x15, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x32, 0xb5, 0x02, 0x0a, 0x12,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70,
	0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x44, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x18, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4a, 0x0a, 0x11, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x6d,
	0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4a, 0x0a, 0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70,
	0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x75, 0x61, 0x6e, 0x79, 0x70, 0x38, 0x2f, 0x77, 0x78, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6b, 0x69, 0x74, 0x65,
	0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6d, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_member_price_proto_rawDescOnce sync.Once
	file_member_price_proto_rawDescData = file_member_price_proto_rawDesc
)

func file_member_price_proto_rawDescGZIP() []byte {
	file_member_price_proto_rawDescOnce.Do(func() {
		file_member_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_member_price_proto_rawDescData)
	})
	return file_member_price_proto_rawDescData
}

var file_member_price_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_member_price_proto_goTypes = []interface{}{
	(*BaseMemberPrice)(nil),       // 0: pms.BaseMemberPrice
	(*MemberPriceAddReq)(nil),     // 1: pms.MemberPriceAddReq
	(*MemberPriceAddResp)(nil),    // 2: pms.MemberPriceAddResp
	(*MemberPriceListReq)(nil),    // 3: pms.MemberPriceListReq
	(*MemberPriceListResp)(nil),   // 4: pms.MemberPriceListResp
	(*MemberPriceUpdateReq)(nil),  // 5: pms.MemberPriceUpdateReq
	(*MemberPriceUpdateResp)(nil), // 6: pms.MemberPriceUpdateResp
	(*MemberPriceDeleteReq)(nil),  // 7: pms.MemberPriceDeleteReq
	(*MemberPriceDeleteResp)(nil), // 8: pms.MemberPriceDeleteResp
}
var file_member_price_proto_depIdxs = []int32{
	0, // 0: pms.MemberPriceAddReq.MemberPrice:type_name -> pms.BaseMemberPrice
	0, // 1: pms.MemberPriceListResp.list:type_name -> pms.BaseMemberPrice
	0, // 2: pms.MemberPriceUpdateReq.MemberPrice:type_name -> pms.BaseMemberPrice
	1, // 3: pms.MemberPriceService.MemberPriceAdd:input_type -> pms.MemberPriceAddReq
	3, // 4: pms.MemberPriceService.MemberPriceList:input_type -> pms.MemberPriceListReq
	5, // 5: pms.MemberPriceService.MemberPriceUpdate:input_type -> pms.MemberPriceUpdateReq
	7, // 6: pms.MemberPriceService.MemberPriceDelete:input_type -> pms.MemberPriceDeleteReq
	2, // 7: pms.MemberPriceService.MemberPriceAdd:output_type -> pms.MemberPriceAddResp
	4, // 8: pms.MemberPriceService.MemberPriceList:output_type -> pms.MemberPriceListResp
	6, // 9: pms.MemberPriceService.MemberPriceUpdate:output_type -> pms.MemberPriceUpdateResp
	8, // 10: pms.MemberPriceService.MemberPriceDelete:output_type -> pms.MemberPriceDeleteResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_member_price_proto_init() }
func file_member_price_proto_init() {
	if File_member_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_member_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseMemberPrice); i {
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
		file_member_price_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceAddReq); i {
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
		file_member_price_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceAddResp); i {
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
		file_member_price_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceListReq); i {
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
		file_member_price_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceListResp); i {
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
		file_member_price_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceUpdateReq); i {
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
		file_member_price_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceUpdateResp); i {
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
		file_member_price_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceDeleteReq); i {
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
		file_member_price_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceDeleteResp); i {
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
			RawDescriptor: file_member_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_member_price_proto_goTypes,
		DependencyIndexes: file_member_price_proto_depIdxs,
		MessageInfos:      file_member_price_proto_msgTypes,
	}.Build()
	File_member_price_proto = out.File
	file_member_price_proto_rawDesc = nil
	file_member_price_proto_goTypes = nil
	file_member_price_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.0. DO NOT EDIT.

type MemberPriceService interface {
	MemberPriceAdd(ctx context.Context, req *MemberPriceAddReq) (res *MemberPriceAddResp, err error)
	MemberPriceList(ctx context.Context, req *MemberPriceListReq) (res *MemberPriceListResp, err error)
	MemberPriceUpdate(ctx context.Context, req *MemberPriceUpdateReq) (res *MemberPriceUpdateResp, err error)
	MemberPriceDelete(ctx context.Context, req *MemberPriceDeleteReq) (res *MemberPriceDeleteResp, err error)
}
