// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: product_attribute_category.proto

package product_attribute_category

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

type BaseProductAttributeCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	AttributeCount int64  `protobuf:"varint,3,opt,name=AttributeCount,proto3" json:"AttributeCount,omitempty"` // 属性数量
	ParamCount     int64  `protobuf:"varint,4,opt,name=ParamCount,proto3" json:"ParamCount,omitempty"`         // 参数数量
}

func (x *BaseProductAttributeCategory) Reset() {
	*x = BaseProductAttributeCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_attribute_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseProductAttributeCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseProductAttributeCategory) ProtoMessage() {}

func (x *BaseProductAttributeCategory) ProtoReflect() protoreflect.Message {
	mi := &file_product_attribute_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseProductAttributeCategory.ProtoReflect.Descriptor instead.
func (*BaseProductAttributeCategory) Descriptor() ([]byte, []int) {
	return file_product_attribute_category_proto_rawDescGZIP(), []int{0}
}

func (x *BaseProductAttributeCategory) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseProductAttributeCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BaseProductAttributeCategory) GetAttributeCount() int64 {
	if x != nil {
		return x.AttributeCount
	}
	return 0
}

func (x *BaseProductAttributeCategory) GetParamCount() int64 {
	if x != nil {
		return x.ParamCount
	}
	return 0
}

type ProductAttributeCategoryAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *ProductAttributeCategoryAddReq) Reset() {
	*x = ProductAttributeCategoryAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_attribute_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductAttributeCategoryAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAttributeCategoryAddReq) ProtoMessage() {}

func (x *ProductAttributeCategoryAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_attribute_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAttributeCategoryAddReq.ProtoReflect.Descriptor instead.
func (*ProductAttributeCategoryAddReq) Descriptor() ([]byte, []int) {
	return file_product_attribute_category_proto_rawDescGZIP(), []int{1}
}

func (x *ProductAttributeCategoryAddReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProductAttributeCategoryAddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *ProductAttributeCategoryAddResp) Reset() {
	*x = ProductAttributeCategoryAddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_attribute_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductAttributeCategoryAddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAttributeCategoryAddResp) ProtoMessage() {}

func (x *ProductAttributeCategoryAddResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_attribute_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAttributeCategoryAddResp.ProtoReflect.Descriptor instead.
func (*ProductAttributeCategoryAddResp) Descriptor() ([]byte, []int) {
	return file_product_attribute_category_proto_rawDescGZIP(), []int{2}
}

func (x *ProductAttributeCategoryAddResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type ProductAttributeCategoryListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current  int64  `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *ProductAttributeCategoryListReq) Reset() {
	*x = ProductAttributeCategoryListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_attribute_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductAttributeCategoryListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAttributeCategoryListReq) ProtoMessage() {}

func (x *ProductAttributeCategoryListReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_attribute_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAttributeCategoryListReq.ProtoReflect.Descriptor instead.
func (*ProductAttributeCategoryListReq) Descriptor() ([]byte, []int) {
	return file_product_attribute_category_proto_rawDescGZIP(), []int{3}
}

func (x *ProductAttributeCategoryListReq) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *ProductAttributeCategoryListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ProductAttributeCategoryListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProductAttributeCategoryListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                           `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*BaseProductAttributeCategory `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ProductAttributeCategoryListResp) Reset() {
	*x = ProductAttributeCategoryListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_attribute_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductAttributeCategoryListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAttributeCategoryListResp) ProtoMessage() {}

func (x *ProductAttributeCategoryListResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_attribute_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAttributeCategoryListResp.ProtoReflect.Descriptor instead.
func (*ProductAttributeCategoryListResp) Descriptor() ([]byte, []int) {
	return file_product_attribute_category_proto_rawDescGZIP(), []int{4}
}

func (x *ProductAttributeCategoryListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ProductAttributeCategoryListResp) GetList() []*BaseProductAttributeCategory {
	if x != nil {
		return x.List
	}
	return nil
}

type ProductAttributeCategoryUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductAttributeCategory *BaseProductAttributeCategory `protobuf:"bytes,1,opt,name=ProductAttributeCategory,proto3" json:"ProductAttributeCategory,omitempty"`
}

func (x *ProductAttributeCategoryUpdateReq) Reset() {
	*x = ProductAttributeCategoryUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_attribute_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductAttributeCategoryUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAttributeCategoryUpdateReq) ProtoMessage() {}

func (x *ProductAttributeCategoryUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_attribute_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAttributeCategoryUpdateReq.ProtoReflect.Descriptor instead.
func (*ProductAttributeCategoryUpdateReq) Descriptor() ([]byte, []int) {
	return file_product_attribute_category_proto_rawDescGZIP(), []int{5}
}

func (x *ProductAttributeCategoryUpdateReq) GetProductAttributeCategory() *BaseProductAttributeCategory {
	if x != nil {
		return x.ProductAttributeCategory
	}
	return nil
}

type ProductAttributeCategoryUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *ProductAttributeCategoryUpdateResp) Reset() {
	*x = ProductAttributeCategoryUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_attribute_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductAttributeCategoryUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAttributeCategoryUpdateResp) ProtoMessage() {}

func (x *ProductAttributeCategoryUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_attribute_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAttributeCategoryUpdateResp.ProtoReflect.Descriptor instead.
func (*ProductAttributeCategoryUpdateResp) Descriptor() ([]byte, []int) {
	return file_product_attribute_category_proto_rawDescGZIP(), []int{6}
}

func (x *ProductAttributeCategoryUpdateResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type ProductAttributeCategoryDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ProductAttributeCategoryDeleteReq) Reset() {
	*x = ProductAttributeCategoryDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_attribute_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductAttributeCategoryDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAttributeCategoryDeleteReq) ProtoMessage() {}

func (x *ProductAttributeCategoryDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_attribute_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAttributeCategoryDeleteReq.ProtoReflect.Descriptor instead.
func (*ProductAttributeCategoryDeleteReq) Descriptor() ([]byte, []int) {
	return file_product_attribute_category_proto_rawDescGZIP(), []int{7}
}

func (x *ProductAttributeCategoryDeleteReq) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ProductAttributeCategoryDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *ProductAttributeCategoryDeleteResp) Reset() {
	*x = ProductAttributeCategoryDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_attribute_category_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductAttributeCategoryDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAttributeCategoryDeleteResp) ProtoMessage() {}

func (x *ProductAttributeCategoryDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_attribute_category_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAttributeCategoryDeleteResp.ProtoReflect.Descriptor instead.
func (*ProductAttributeCategoryDeleteResp) Descriptor() ([]byte, []int) {
	return file_product_attribute_category_proto_rawDescGZIP(), []int{8}
}

func (x *ProductAttributeCategoryDeleteResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

var File_product_attribute_category_proto protoreflect.FileDescriptor

var file_product_attribute_category_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x6d, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x1c, 0x42, 0x61, 0x73, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x34, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x1f, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x22, 0x6b, 0x0a, 0x1f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6f,
	0x0a, 0x20, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22,
	0x82, 0x01, 0x0a, 0x21, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x5d, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x18, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x22, 0x38, 0x0a, 0x22, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x35,
	0x0a, 0x21, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x38, 0x0a, 0x22, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x32,
	0xde, 0x03, 0x0a, 0x1f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x1b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41,
	0x64, 0x64, 0x12, 0x23, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x6b, 0x0a,
	0x1c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e,
	0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x71, 0x0a, 0x1e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x70,
	0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x71, 0x0a,
	0x1e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x26, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x75, 0x61, 0x6e, 0x79, 0x70, 0x38, 0x2f, 0x77, 0x78, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_attribute_category_proto_rawDescOnce sync.Once
	file_product_attribute_category_proto_rawDescData = file_product_attribute_category_proto_rawDesc
)

func file_product_attribute_category_proto_rawDescGZIP() []byte {
	file_product_attribute_category_proto_rawDescOnce.Do(func() {
		file_product_attribute_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_attribute_category_proto_rawDescData)
	})
	return file_product_attribute_category_proto_rawDescData
}

var file_product_attribute_category_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_product_attribute_category_proto_goTypes = []interface{}{
	(*BaseProductAttributeCategory)(nil),       // 0: pms.BaseProductAttributeCategory
	(*ProductAttributeCategoryAddReq)(nil),     // 1: pms.ProductAttributeCategoryAddReq
	(*ProductAttributeCategoryAddResp)(nil),    // 2: pms.ProductAttributeCategoryAddResp
	(*ProductAttributeCategoryListReq)(nil),    // 3: pms.ProductAttributeCategoryListReq
	(*ProductAttributeCategoryListResp)(nil),   // 4: pms.ProductAttributeCategoryListResp
	(*ProductAttributeCategoryUpdateReq)(nil),  // 5: pms.ProductAttributeCategoryUpdateReq
	(*ProductAttributeCategoryUpdateResp)(nil), // 6: pms.ProductAttributeCategoryUpdateResp
	(*ProductAttributeCategoryDeleteReq)(nil),  // 7: pms.ProductAttributeCategoryDeleteReq
	(*ProductAttributeCategoryDeleteResp)(nil), // 8: pms.ProductAttributeCategoryDeleteResp
}
var file_product_attribute_category_proto_depIdxs = []int32{
	0, // 0: pms.ProductAttributeCategoryListResp.list:type_name -> pms.BaseProductAttributeCategory
	0, // 1: pms.ProductAttributeCategoryUpdateReq.ProductAttributeCategory:type_name -> pms.BaseProductAttributeCategory
	1, // 2: pms.ProductAttributeCategoryService.ProductAttributeCategoryAdd:input_type -> pms.ProductAttributeCategoryAddReq
	3, // 3: pms.ProductAttributeCategoryService.ProductAttributeCategoryList:input_type -> pms.ProductAttributeCategoryListReq
	5, // 4: pms.ProductAttributeCategoryService.ProductAttributeCategoryUpdate:input_type -> pms.ProductAttributeCategoryUpdateReq
	7, // 5: pms.ProductAttributeCategoryService.ProductAttributeCategoryDelete:input_type -> pms.ProductAttributeCategoryDeleteReq
	2, // 6: pms.ProductAttributeCategoryService.ProductAttributeCategoryAdd:output_type -> pms.ProductAttributeCategoryAddResp
	4, // 7: pms.ProductAttributeCategoryService.ProductAttributeCategoryList:output_type -> pms.ProductAttributeCategoryListResp
	6, // 8: pms.ProductAttributeCategoryService.ProductAttributeCategoryUpdate:output_type -> pms.ProductAttributeCategoryUpdateResp
	8, // 9: pms.ProductAttributeCategoryService.ProductAttributeCategoryDelete:output_type -> pms.ProductAttributeCategoryDeleteResp
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_product_attribute_category_proto_init() }
func file_product_attribute_category_proto_init() {
	if File_product_attribute_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_attribute_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseProductAttributeCategory); i {
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
		file_product_attribute_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductAttributeCategoryAddReq); i {
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
		file_product_attribute_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductAttributeCategoryAddResp); i {
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
		file_product_attribute_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductAttributeCategoryListReq); i {
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
		file_product_attribute_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductAttributeCategoryListResp); i {
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
		file_product_attribute_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductAttributeCategoryUpdateReq); i {
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
		file_product_attribute_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductAttributeCategoryUpdateResp); i {
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
		file_product_attribute_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductAttributeCategoryDeleteReq); i {
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
		file_product_attribute_category_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductAttributeCategoryDeleteResp); i {
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
			RawDescriptor: file_product_attribute_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_attribute_category_proto_goTypes,
		DependencyIndexes: file_product_attribute_category_proto_depIdxs,
		MessageInfos:      file_product_attribute_category_proto_msgTypes,
	}.Build()
	File_product_attribute_category_proto = out.File
	file_product_attribute_category_proto_rawDesc = nil
	file_product_attribute_category_proto_goTypes = nil
	file_product_attribute_category_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.0. DO NOT EDIT.

type ProductAttributeCategoryService interface {
	ProductAttributeCategoryAdd(ctx context.Context, req *ProductAttributeCategoryAddReq) (res *ProductAttributeCategoryAddResp, err error)
	ProductAttributeCategoryList(ctx context.Context, req *ProductAttributeCategoryListReq) (res *ProductAttributeCategoryListResp, err error)
	ProductAttributeCategoryUpdate(ctx context.Context, req *ProductAttributeCategoryUpdateReq) (res *ProductAttributeCategoryUpdateResp, err error)
	ProductAttributeCategoryDelete(ctx context.Context, req *ProductAttributeCategoryDeleteReq) (res *ProductAttributeCategoryDeleteResp, err error)
}
