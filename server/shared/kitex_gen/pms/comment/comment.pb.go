// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: comment.proto

package comment

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

type BaseComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: 去掉自增id
	Id               uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ProductId        uint64 `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	MemberNickName   string `protobuf:"bytes,3,opt,name=MemberNickName,proto3" json:"MemberNickName,omitempty"`     // 用户昵称
	ProductName      string `protobuf:"bytes,4,opt,name=ProductName,proto3" json:"ProductName,omitempty"`           // 商品名称
	Star             int64  `protobuf:"varint,5,opt,name=Star,proto3" json:"Star,omitempty"`                        // 评价星数：0->5
	MemberIp         string `protobuf:"bytes,6,opt,name=MemberIp,proto3" json:"MemberIp,omitempty"`                 // 用户ip
	CreateTime       string `protobuf:"bytes,7,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`             // 评价时间
	ShowStatus       int64  `protobuf:"varint,8,opt,name=ShowStatus,proto3" json:"ShowStatus,omitempty"`            // 是否显示：0->不显示；1->显示
	ProductAttribute string `protobuf:"bytes,9,opt,name=ProductAttribute,proto3" json:"ProductAttribute,omitempty"` // 购买时的商品属性
	CollectCount     int64  `protobuf:"varint,10,opt,name=CollectCount,proto3" json:"CollectCount,omitempty"`       // 收藏数
	ReadCount        int64  `protobuf:"varint,11,opt,name=ReadCount,proto3" json:"ReadCount,omitempty"`             // 阅读数
	Content          string `protobuf:"bytes,12,opt,name=Content,proto3" json:"Content,omitempty"`                  // 评价内容
	Pics             string `protobuf:"bytes,13,opt,name=Pics,proto3" json:"Pics,omitempty"`                        // 评论图片/视频/音频的链接
	MemberIcon       string `protobuf:"bytes,14,opt,name=MemberIcon,proto3" json:"MemberIcon,omitempty"`            // 用户头像
	ReplayCount      int64  `protobuf:"varint,15,opt,name=ReplayCount,proto3" json:"ReplayCount,omitempty"`         // 回复数
}

func (x *BaseComment) Reset() {
	*x = BaseComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseComment) ProtoMessage() {}

func (x *BaseComment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseComment.ProtoReflect.Descriptor instead.
func (*BaseComment) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{0}
}

func (x *BaseComment) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseComment) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *BaseComment) GetMemberNickName() string {
	if x != nil {
		return x.MemberNickName
	}
	return ""
}

func (x *BaseComment) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *BaseComment) GetStar() int64 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *BaseComment) GetMemberIp() string {
	if x != nil {
		return x.MemberIp
	}
	return ""
}

func (x *BaseComment) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *BaseComment) GetShowStatus() int64 {
	if x != nil {
		return x.ShowStatus
	}
	return 0
}

func (x *BaseComment) GetProductAttribute() string {
	if x != nil {
		return x.ProductAttribute
	}
	return ""
}

func (x *BaseComment) GetCollectCount() int64 {
	if x != nil {
		return x.CollectCount
	}
	return 0
}

func (x *BaseComment) GetReadCount() int64 {
	if x != nil {
		return x.ReadCount
	}
	return 0
}

func (x *BaseComment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *BaseComment) GetPics() string {
	if x != nil {
		return x.Pics
	}
	return ""
}

func (x *BaseComment) GetMemberIcon() string {
	if x != nil {
		return x.MemberIcon
	}
	return ""
}

func (x *BaseComment) GetReplayCount() int64 {
	if x != nil {
		return x.ReplayCount
	}
	return 0
}

type CommentAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *BaseComment `protobuf:"bytes,1,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *CommentAddReq) Reset() {
	*x = CommentAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentAddReq) ProtoMessage() {}

func (x *CommentAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentAddReq.ProtoReflect.Descriptor instead.
func (*CommentAddReq) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentAddReq) GetComment() *BaseComment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CommentAddResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *CommentAddResp) Reset() {
	*x = CommentAddResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentAddResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentAddResp) ProtoMessage() {}

func (x *CommentAddResp) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentAddResp.ProtoReflect.Descriptor instead.
func (*CommentAddResp) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CommentAddResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type CommentListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current  int64 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *CommentListReq) Reset() {
	*x = CommentListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListReq) ProtoMessage() {}

func (x *CommentListReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListReq.ProtoReflect.Descriptor instead.
func (*CommentListReq) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{3}
}

func (x *CommentListReq) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *CommentListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type CommentListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*BaseComment `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *CommentListResp) Reset() {
	*x = CommentListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListResp) ProtoMessage() {}

func (x *CommentListResp) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListResp.ProtoReflect.Descriptor instead.
func (*CommentListResp) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{4}
}

func (x *CommentListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CommentListResp) GetList() []*BaseComment {
	if x != nil {
		return x.List
	}
	return nil
}

type CommentUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *BaseComment `protobuf:"bytes,1,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *CommentUpdateReq) Reset() {
	*x = CommentUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentUpdateReq) ProtoMessage() {}

func (x *CommentUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentUpdateReq.ProtoReflect.Descriptor instead.
func (*CommentUpdateReq) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{5}
}

func (x *CommentUpdateReq) GetComment() *BaseComment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CommentUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *CommentUpdateResp) Reset() {
	*x = CommentUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentUpdateResp) ProtoMessage() {}

func (x *CommentUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentUpdateResp.ProtoReflect.Descriptor instead.
func (*CommentUpdateResp) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{6}
}

func (x *CommentUpdateResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type CommentDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *CommentDeleteReq) Reset() {
	*x = CommentDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentDeleteReq) ProtoMessage() {}

func (x *CommentDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentDeleteReq.ProtoReflect.Descriptor instead.
func (*CommentDeleteReq) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{7}
}

func (x *CommentDeleteReq) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CommentDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *CommentDeleteResp) Reset() {
	*x = CommentDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentDeleteResp) ProtoMessage() {}

func (x *CommentDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentDeleteResp.ProtoReflect.Descriptor instead.
func (*CommentDeleteResp) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{8}
}

func (x *CommentDeleteResp) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

var File_comment_proto protoreflect.FileDescriptor

var file_comment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x70, 0x6d, 0x73, 0x22, 0xd3, 0x03, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x53, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x74, 0x61, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x70, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x69, 0x63, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x50, 0x69, 0x63, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x52,
	0x65, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x0d, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x6d, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x46, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4d, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x24,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x6d, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6d, 0x73, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x24, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x32, 0x81, 0x02, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x35, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x12, 0x12, 0x2e,
	0x70, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x13, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x6d, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x6d, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x75, 0x61, 0x6e, 0x79, 0x70, 0x38, 0x2f, 0x77, 0x78, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_proto_rawDescOnce sync.Once
	file_comment_proto_rawDescData = file_comment_proto_rawDesc
)

func file_comment_proto_rawDescGZIP() []byte {
	file_comment_proto_rawDescOnce.Do(func() {
		file_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_proto_rawDescData)
	})
	return file_comment_proto_rawDescData
}

var file_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_comment_proto_goTypes = []interface{}{
	(*BaseComment)(nil),       // 0: pms.BaseComment
	(*CommentAddReq)(nil),     // 1: pms.CommentAddReq
	(*CommentAddResp)(nil),    // 2: pms.CommentAddResp
	(*CommentListReq)(nil),    // 3: pms.CommentListReq
	(*CommentListResp)(nil),   // 4: pms.CommentListResp
	(*CommentUpdateReq)(nil),  // 5: pms.CommentUpdateReq
	(*CommentUpdateResp)(nil), // 6: pms.CommentUpdateResp
	(*CommentDeleteReq)(nil),  // 7: pms.CommentDeleteReq
	(*CommentDeleteResp)(nil), // 8: pms.CommentDeleteResp
}
var file_comment_proto_depIdxs = []int32{
	0, // 0: pms.CommentAddReq.Comment:type_name -> pms.BaseComment
	0, // 1: pms.CommentListResp.list:type_name -> pms.BaseComment
	0, // 2: pms.CommentUpdateReq.Comment:type_name -> pms.BaseComment
	1, // 3: pms.CommentService.CommentAdd:input_type -> pms.CommentAddReq
	3, // 4: pms.CommentService.CommentList:input_type -> pms.CommentListReq
	5, // 5: pms.CommentService.CommentUpdate:input_type -> pms.CommentUpdateReq
	7, // 6: pms.CommentService.CommentDelete:input_type -> pms.CommentDeleteReq
	2, // 7: pms.CommentService.CommentAdd:output_type -> pms.CommentAddResp
	4, // 8: pms.CommentService.CommentList:output_type -> pms.CommentListResp
	6, // 9: pms.CommentService.CommentUpdate:output_type -> pms.CommentUpdateResp
	8, // 10: pms.CommentService.CommentDelete:output_type -> pms.CommentDeleteResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_comment_proto_init() }
func file_comment_proto_init() {
	if File_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseComment); i {
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
		file_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentAddReq); i {
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
		file_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentAddResp); i {
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
		file_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListReq); i {
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
		file_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListResp); i {
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
		file_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentUpdateReq); i {
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
		file_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentUpdateResp); i {
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
		file_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentDeleteReq); i {
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
		file_comment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentDeleteResp); i {
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
			RawDescriptor: file_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_proto_goTypes,
		DependencyIndexes: file_comment_proto_depIdxs,
		MessageInfos:      file_comment_proto_msgTypes,
	}.Build()
	File_comment_proto = out.File
	file_comment_proto_rawDesc = nil
	file_comment_proto_goTypes = nil
	file_comment_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.0. DO NOT EDIT.

type CommentService interface {
	CommentAdd(ctx context.Context, req *CommentAddReq) (res *CommentAddResp, err error)
	CommentList(ctx context.Context, req *CommentListReq) (res *CommentListResp, err error)
	CommentUpdate(ctx context.Context, req *CommentUpdateReq) (res *CommentUpdateResp, err error)
	CommentDelete(ctx context.Context, req *CommentDeleteReq) (res *CommentDeleteResp, err error)
}
