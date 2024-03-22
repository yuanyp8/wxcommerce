// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: product_operate_log.proto

package product_full_reduction

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

type BaseProductOperateLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ProductId        int64  `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	PriceOld         int64  `protobuf:"varint,3,opt,name=PriceOld,proto3" json:"PriceOld,omitempty"`
	PriceNew         int64  `protobuf:"varint,4,opt,name=PriceNew,proto3" json:"PriceNew,omitempty"`
	SalePriceOld     int64  `protobuf:"varint,5,opt,name=SalePriceOld,proto3" json:"SalePriceOld,omitempty"`
	SalePriceNew     int64  `protobuf:"varint,6,opt,name=SalePriceNew,proto3" json:"SalePriceNew,omitempty"`
	GiftPointOld     int64  `protobuf:"varint,7,opt,name=GiftPointOld,proto3" json:"GiftPointOld,omitempty"`
	GiftPointNew     int64  `protobuf:"varint,8,opt,name=GiftPointNew,proto3" json:"GiftPointNew,omitempty"`
	UsePointLimitOld int64  `protobuf:"varint,9,opt,name=UsePointLimitOld,proto3" json:"UsePointLimitOld,omitempty"`
	UsePointLimitNew int64  `protobuf:"varint,10,opt,name=UsePointLimitNew,proto3" json:"UsePointLimitNew,omitempty"`
	OperateMan       string `protobuf:"bytes,11,opt,name=OperateMan,proto3" json:"OperateMan,omitempty"`
	CreateTime       string `protobuf:"bytes,12,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *BaseProductOperateLog) Reset() {
	*x = BaseProductOperateLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_operate_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseProductOperateLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseProductOperateLog) ProtoMessage() {}

func (x *BaseProductOperateLog) ProtoReflect() protoreflect.Message {
	mi := &file_product_operate_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseProductOperateLog.ProtoReflect.Descriptor instead.
func (*BaseProductOperateLog) Descriptor() ([]byte, []int) {
	return file_product_operate_log_proto_rawDescGZIP(), []int{0}
}

func (x *BaseProductOperateLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseProductOperateLog) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *BaseProductOperateLog) GetPriceOld() int64 {
	if x != nil {
		return x.PriceOld
	}
	return 0
}

func (x *BaseProductOperateLog) GetPriceNew() int64 {
	if x != nil {
		return x.PriceNew
	}
	return 0
}

func (x *BaseProductOperateLog) GetSalePriceOld() int64 {
	if x != nil {
		return x.SalePriceOld
	}
	return 0
}

func (x *BaseProductOperateLog) GetSalePriceNew() int64 {
	if x != nil {
		return x.SalePriceNew
	}
	return 0
}

func (x *BaseProductOperateLog) GetGiftPointOld() int64 {
	if x != nil {
		return x.GiftPointOld
	}
	return 0
}

func (x *BaseProductOperateLog) GetGiftPointNew() int64 {
	if x != nil {
		return x.GiftPointNew
	}
	return 0
}

func (x *BaseProductOperateLog) GetUsePointLimitOld() int64 {
	if x != nil {
		return x.UsePointLimitOld
	}
	return 0
}

func (x *BaseProductOperateLog) GetUsePointLimitNew() int64 {
	if x != nil {
		return x.UsePointLimitNew
	}
	return 0
}

func (x *BaseProductOperateLog) GetOperateMan() string {
	if x != nil {
		return x.OperateMan
	}
	return ""
}

func (x *BaseProductOperateLog) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

type ProductOperateLogAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductOperateLog *BaseProductOperateLog `protobuf:"bytes,1,opt,name=ProductOperateLog,proto3" json:"ProductOperateLog,omitempty"`
}

func (x *ProductOperateLogAddReq) Reset() {
	*x = ProductOperateLogAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_operate_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOperateLogAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOperateLogAddReq) ProtoMessage() {}

func (x *ProductOperateLogAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_operate_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOperateLogAddReq.ProtoReflect.Descriptor instead.
func (*ProductOperateLogAddReq) Descriptor() ([]byte, []int) {
	return file_product_operate_log_proto_rawDescGZIP(), []int{1}
}

func (x *ProductOperateLogAddReq) GetProductOperateLog() *BaseProductOperateLog {
	if x != nil {
		return x.ProductOperateLog
	}
	return nil
}

type ProductOperateLogAddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *ProductOperateLogAddResp) Reset() {
	*x = ProductOperateLogAddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_operate_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOperateLogAddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOperateLogAddResp) ProtoMessage() {}

func (x *ProductOperateLogAddResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_operate_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOperateLogAddResp.ProtoReflect.Descriptor instead.
func (*ProductOperateLogAddResp) Descriptor() ([]byte, []int) {
	return file_product_operate_log_proto_rawDescGZIP(), []int{2}
}

func (x *ProductOperateLogAddResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type ProductOperateLogListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current  int64 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ProductOperateLogListReq) Reset() {
	*x = ProductOperateLogListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_operate_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOperateLogListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOperateLogListReq) ProtoMessage() {}

func (x *ProductOperateLogListReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_operate_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOperateLogListReq.ProtoReflect.Descriptor instead.
func (*ProductOperateLogListReq) Descriptor() ([]byte, []int) {
	return file_product_operate_log_proto_rawDescGZIP(), []int{3}
}

func (x *ProductOperateLogListReq) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *ProductOperateLogListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ProductOperateLogListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*BaseProductOperateLog `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ProductOperateLogListResp) Reset() {
	*x = ProductOperateLogListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_operate_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOperateLogListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOperateLogListResp) ProtoMessage() {}

func (x *ProductOperateLogListResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_operate_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOperateLogListResp.ProtoReflect.Descriptor instead.
func (*ProductOperateLogListResp) Descriptor() ([]byte, []int) {
	return file_product_operate_log_proto_rawDescGZIP(), []int{4}
}

func (x *ProductOperateLogListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ProductOperateLogListResp) GetList() []*BaseProductOperateLog {
	if x != nil {
		return x.List
	}
	return nil
}

type ProductOperateLogUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductOperateLog *BaseProductOperateLog `protobuf:"bytes,1,opt,name=ProductOperateLog,proto3" json:"ProductOperateLog,omitempty"`
}

func (x *ProductOperateLogUpdateReq) Reset() {
	*x = ProductOperateLogUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_operate_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOperateLogUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOperateLogUpdateReq) ProtoMessage() {}

func (x *ProductOperateLogUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_operate_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOperateLogUpdateReq.ProtoReflect.Descriptor instead.
func (*ProductOperateLogUpdateReq) Descriptor() ([]byte, []int) {
	return file_product_operate_log_proto_rawDescGZIP(), []int{5}
}

func (x *ProductOperateLogUpdateReq) GetProductOperateLog() *BaseProductOperateLog {
	if x != nil {
		return x.ProductOperateLog
	}
	return nil
}

type ProductOperateLogUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *ProductOperateLogUpdateResp) Reset() {
	*x = ProductOperateLogUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_operate_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOperateLogUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOperateLogUpdateResp) ProtoMessage() {}

func (x *ProductOperateLogUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_operate_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOperateLogUpdateResp.ProtoReflect.Descriptor instead.
func (*ProductOperateLogUpdateResp) Descriptor() ([]byte, []int) {
	return file_product_operate_log_proto_rawDescGZIP(), []int{6}
}

func (x *ProductOperateLogUpdateResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type ProductOperateLogDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductOperateLogDeleteReq) Reset() {
	*x = ProductOperateLogDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_operate_log_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOperateLogDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOperateLogDeleteReq) ProtoMessage() {}

func (x *ProductOperateLogDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_operate_log_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOperateLogDeleteReq.ProtoReflect.Descriptor instead.
func (*ProductOperateLogDeleteReq) Descriptor() ([]byte, []int) {
	return file_product_operate_log_proto_rawDescGZIP(), []int{7}
}

func (x *ProductOperateLogDeleteReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProductOperateLogDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *ProductOperateLogDeleteResp) Reset() {
	*x = ProductOperateLogDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_operate_log_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOperateLogDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOperateLogDeleteResp) ProtoMessage() {}

func (x *ProductOperateLogDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_operate_log_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOperateLogDeleteResp.ProtoReflect.Descriptor instead.
func (*ProductOperateLogDeleteResp) Descriptor() ([]byte, []int) {
	return file_product_operate_log_proto_rawDescGZIP(), []int{8}
}

func (x *ProductOperateLogDeleteResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

var File_product_operate_log_proto protoreflect.FileDescriptor

var file_product_operate_log_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x6d, 0x73,
	0x22, 0xa5, 0x03, 0x0a, 0x15, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x4f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x4f, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4e, 0x65, 0x77,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4e, 0x65, 0x77,
	0x12, 0x22, 0x0a, 0x0c, 0x53, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4f, 0x6c, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x53, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x4f, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x4e, 0x65, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x53, 0x61, 0x6c, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x4e, 0x65, 0x77, 0x12, 0x22, 0x0a, 0x0c, 0x47, 0x69, 0x66, 0x74,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x47, 0x69, 0x66, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x47, 0x69, 0x66, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4e, 0x65, 0x77, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x47, 0x69, 0x66, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4e, 0x65, 0x77,
	0x12, 0x2a, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x4f, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x55, 0x73, 0x65, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x10,
	0x55, 0x73, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4e, 0x65, 0x77,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x55, 0x73, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4e, 0x65, 0x77, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x48, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x11, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x22, 0x2e, 0x0a,
	0x18, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x67, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x50, 0x0a,
	0x18, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x61, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x66, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x48, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6d,
	0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x22, 0x31, 0x0a, 0x1b, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x2c, 0x0a,
	0x1a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x1b, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x32, 0x83,
	0x03, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x14, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x41, 0x64, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x1d, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x56, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x6d, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5c, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5c, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x1f, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x20, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x61, 0x6e, 0x79, 0x70, 0x38, 0x2f, 0x77, 0x78, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6b, 0x69, 0x74,
	0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_operate_log_proto_rawDescOnce sync.Once
	file_product_operate_log_proto_rawDescData = file_product_operate_log_proto_rawDesc
)

func file_product_operate_log_proto_rawDescGZIP() []byte {
	file_product_operate_log_proto_rawDescOnce.Do(func() {
		file_product_operate_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_operate_log_proto_rawDescData)
	})
	return file_product_operate_log_proto_rawDescData
}

var file_product_operate_log_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_product_operate_log_proto_goTypes = []interface{}{
	(*BaseProductOperateLog)(nil),       // 0: pms.BaseProductOperateLog
	(*ProductOperateLogAddReq)(nil),     // 1: pms.ProductOperateLogAddReq
	(*ProductOperateLogAddResp)(nil),    // 2: pms.ProductOperateLogAddResp
	(*ProductOperateLogListReq)(nil),    // 3: pms.ProductOperateLogListReq
	(*ProductOperateLogListResp)(nil),   // 4: pms.ProductOperateLogListResp
	(*ProductOperateLogUpdateReq)(nil),  // 5: pms.ProductOperateLogUpdateReq
	(*ProductOperateLogUpdateResp)(nil), // 6: pms.ProductOperateLogUpdateResp
	(*ProductOperateLogDeleteReq)(nil),  // 7: pms.ProductOperateLogDeleteReq
	(*ProductOperateLogDeleteResp)(nil), // 8: pms.ProductOperateLogDeleteResp
}
var file_product_operate_log_proto_depIdxs = []int32{
	0, // 0: pms.ProductOperateLogAddReq.ProductOperateLog:type_name -> pms.BaseProductOperateLog
	0, // 1: pms.ProductOperateLogListResp.list:type_name -> pms.BaseProductOperateLog
	0, // 2: pms.ProductOperateLogUpdateReq.ProductOperateLog:type_name -> pms.BaseProductOperateLog
	1, // 3: pms.ProductOperateLogService.ProductOperateLogAdd:input_type -> pms.ProductOperateLogAddReq
	3, // 4: pms.ProductOperateLogService.ProductOperateLogList:input_type -> pms.ProductOperateLogListReq
	5, // 5: pms.ProductOperateLogService.ProductOperateLogUpdate:input_type -> pms.ProductOperateLogUpdateReq
	7, // 6: pms.ProductOperateLogService.ProductOperateLogDelete:input_type -> pms.ProductOperateLogDeleteReq
	2, // 7: pms.ProductOperateLogService.ProductOperateLogAdd:output_type -> pms.ProductOperateLogAddResp
	4, // 8: pms.ProductOperateLogService.ProductOperateLogList:output_type -> pms.ProductOperateLogListResp
	6, // 9: pms.ProductOperateLogService.ProductOperateLogUpdate:output_type -> pms.ProductOperateLogUpdateResp
	8, // 10: pms.ProductOperateLogService.ProductOperateLogDelete:output_type -> pms.ProductOperateLogDeleteResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_product_operate_log_proto_init() }
func file_product_operate_log_proto_init() {
	if File_product_operate_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_operate_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseProductOperateLog); i {
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
		file_product_operate_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOperateLogAddReq); i {
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
		file_product_operate_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOperateLogAddResp); i {
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
		file_product_operate_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOperateLogListReq); i {
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
		file_product_operate_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOperateLogListResp); i {
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
		file_product_operate_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOperateLogUpdateReq); i {
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
		file_product_operate_log_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOperateLogUpdateResp); i {
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
		file_product_operate_log_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOperateLogDeleteReq); i {
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
		file_product_operate_log_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductOperateLogDeleteResp); i {
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
			RawDescriptor: file_product_operate_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_operate_log_proto_goTypes,
		DependencyIndexes: file_product_operate_log_proto_depIdxs,
		MessageInfos:      file_product_operate_log_proto_msgTypes,
	}.Build()
	File_product_operate_log_proto = out.File
	file_product_operate_log_proto_rawDesc = nil
	file_product_operate_log_proto_goTypes = nil
	file_product_operate_log_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.0. DO NOT EDIT.

type ProductOperateLogService interface {
	ProductOperateLogAdd(ctx context.Context, req *ProductOperateLogAddReq) (res *ProductOperateLogAddResp, err error)
	ProductOperateLogList(ctx context.Context, req *ProductOperateLogListReq) (res *ProductOperateLogListResp, err error)
	ProductOperateLogUpdate(ctx context.Context, req *ProductOperateLogUpdateReq) (res *ProductOperateLogUpdateResp, err error)
	ProductOperateLogDelete(ctx context.Context, req *ProductOperateLogDeleteReq) (res *ProductOperateLogDeleteResp, err error)
}