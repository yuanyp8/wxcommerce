// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: errno.proto

package errno

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

type Err int32

const (
	Err_Success                                      Err = 0
	Err_NoRoute                                      Err = 1
	Err_NoMethod                                     Err = 2
	Err_BadRequest                                   Err = 10000
	Err_ParamsErr                                    Err = 10001
	Err_AuthorizeFail                                Err = 10002
	Err_TooManyRequest                               Err = 10003
	Err_ServiceErr                                   Err = 20000
	Err_RPCPmsAlbumSrvErr                            Err = 30000
	Err_PmsAlbumSrvErr                               Err = 30001
	Err_RPCPmsAlbumPicSrvErr                         Err = 30002
	Err_PmsAlbumPicSrvErr                            Err = 30003
	Err_RPCPmsBrandSrvErr                            Err = 30004
	Err_PmsBrandSrvErr                               Err = 30005
	Err_RPCPmsCommentSrvErr                          Err = 30006
	Err_PmsCommentSrvErr                             Err = 30007
	Err_RPCPmsCommentReplySrvErr                     Err = 30008
	Err_PmsCommentReplySrvErr                        Err = 30009
	Err_RPCPmsMemberPriceSrvErr                      Err = 30010
	Err_PmsMemberPriceSrvErr                         Err = 30011
	Err_RPCPmsProductSrvErr                          Err = 30012
	Err_PmsProductSrvErr                             Err = 30013
	Err_RPCPmsProductAttrSrvErr                      Err = 30014
	Err_PmsProductAttrSrvErr                         Err = 30015
	Err_RPCPmsProductAttributeCategorySrvErr         Err = 30016
	Err_PmsProductAttributeCategorySrvErr            Err = 30017
	Err_RPCPmsProductAttributeValueSrvErr            Err = 30018
	Err_PmsProductAttributeValueSrvErr               Err = 30019
	Err_RPCPmsProductCategorySrvErr                  Err = 30020
	Err_PmsProductCategorySrvErr                     Err = 30021
	Err_RPCPmsProductCategoryAttributeRelationSrvErr Err = 30022
	Err_PmsProductCategoryAttributeRelationSrvErr    Err = 30023
	Err_RPCPmsProductFullReductionSrvErr             Err = 30024
	Err_PmsProductFullReductionSrvErr                Err = 30025
	Err_RPCPmsProductLadderSrvErr                    Err = 30026
	Err_PmsProductLadderSrvErr                       Err = 30027
	Err_RPCPmsProductVerifyRecordSrvErr              Err = 30028
	Err_PmsProductVerifyRecordSrvErr                 Err = 30029
	Err_RPCPmsSkuStockSrvErr                         Err = 30030
	Err_PmsSkuStockSrvErr                            Err = 30031
	Err_RecordNotFound                               Err = 80000
	Err_RecordAlreadyExist                           Err = 80001
	Err_DirtyData                                    Err = 80003
)

// Enum value maps for Err.
var (
	Err_name = map[int32]string{
		0:     "Success",
		1:     "NoRoute",
		2:     "NoMethod",
		10000: "BadRequest",
		10001: "ParamsErr",
		10002: "AuthorizeFail",
		10003: "TooManyRequest",
		20000: "ServiceErr",
		30000: "RPCPmsAlbumSrvErr",
		30001: "PmsAlbumSrvErr",
		30002: "RPCPmsAlbumPicSrvErr",
		30003: "PmsAlbumPicSrvErr",
		30004: "RPCPmsBrandSrvErr",
		30005: "PmsBrandSrvErr",
		30006: "RPCPmsCommentSrvErr",
		30007: "PmsCommentSrvErr",
		30008: "RPCPmsCommentReplySrvErr",
		30009: "PmsCommentReplySrvErr",
		30010: "RPCPmsMemberPriceSrvErr",
		30011: "PmsMemberPriceSrvErr",
		30012: "RPCPmsProductSrvErr",
		30013: "PmsProductSrvErr",
		30014: "RPCPmsProductAttrSrvErr",
		30015: "PmsProductAttrSrvErr",
		30016: "RPCPmsProductAttributeCategorySrvErr",
		30017: "PmsProductAttributeCategorySrvErr",
		30018: "RPCPmsProductAttributeValueSrvErr",
		30019: "PmsProductAttributeValueSrvErr",
		30020: "RPCPmsProductCategorySrvErr",
		30021: "PmsProductCategorySrvErr",
		30022: "RPCPmsProductCategoryAttributeRelationSrvErr",
		30023: "PmsProductCategoryAttributeRelationSrvErr",
		30024: "RPCPmsProductFullReductionSrvErr",
		30025: "PmsProductFullReductionSrvErr",
		30026: "RPCPmsProductLadderSrvErr",
		30027: "PmsProductLadderSrvErr",
		30028: "RPCPmsProductVerifyRecordSrvErr",
		30029: "PmsProductVerifyRecordSrvErr",
		30030: "RPCPmsSkuStockSrvErr",
		30031: "PmsSkuStockSrvErr",
		80000: "RecordNotFound",
		80001: "RecordAlreadyExist",
		80003: "DirtyData",
	}
	Err_value = map[string]int32{
		"Success":                                      0,
		"NoRoute":                                      1,
		"NoMethod":                                     2,
		"BadRequest":                                   10000,
		"ParamsErr":                                    10001,
		"AuthorizeFail":                                10002,
		"TooManyRequest":                               10003,
		"ServiceErr":                                   20000,
		"RPCPmsAlbumSrvErr":                            30000,
		"PmsAlbumSrvErr":                               30001,
		"RPCPmsAlbumPicSrvErr":                         30002,
		"PmsAlbumPicSrvErr":                            30003,
		"RPCPmsBrandSrvErr":                            30004,
		"PmsBrandSrvErr":                               30005,
		"RPCPmsCommentSrvErr":                          30006,
		"PmsCommentSrvErr":                             30007,
		"RPCPmsCommentReplySrvErr":                     30008,
		"PmsCommentReplySrvErr":                        30009,
		"RPCPmsMemberPriceSrvErr":                      30010,
		"PmsMemberPriceSrvErr":                         30011,
		"RPCPmsProductSrvErr":                          30012,
		"PmsProductSrvErr":                             30013,
		"RPCPmsProductAttrSrvErr":                      30014,
		"PmsProductAttrSrvErr":                         30015,
		"RPCPmsProductAttributeCategorySrvErr":         30016,
		"PmsProductAttributeCategorySrvErr":            30017,
		"RPCPmsProductAttributeValueSrvErr":            30018,
		"PmsProductAttributeValueSrvErr":               30019,
		"RPCPmsProductCategorySrvErr":                  30020,
		"PmsProductCategorySrvErr":                     30021,
		"RPCPmsProductCategoryAttributeRelationSrvErr": 30022,
		"PmsProductCategoryAttributeRelationSrvErr":    30023,
		"RPCPmsProductFullReductionSrvErr":             30024,
		"PmsProductFullReductionSrvErr":                30025,
		"RPCPmsProductLadderSrvErr":                    30026,
		"PmsProductLadderSrvErr":                       30027,
		"RPCPmsProductVerifyRecordSrvErr":              30028,
		"PmsProductVerifyRecordSrvErr":                 30029,
		"RPCPmsSkuStockSrvErr":                         30030,
		"PmsSkuStockSrvErr":                            30031,
		"RecordNotFound":                               80000,
		"RecordAlreadyExist":                           80001,
		"DirtyData":                                    80003,
	}
)

func (x Err) Enum() *Err {
	p := new(Err)
	*p = x
	return p
}

func (x Err) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Err) Descriptor() protoreflect.EnumDescriptor {
	return file_errno_proto_enumTypes[0].Descriptor()
}

func (Err) Type() protoreflect.EnumType {
	return &file_errno_proto_enumTypes[0]
}

func (x Err) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Err.Descriptor instead.
func (Err) EnumDescriptor() ([]byte, []int) {
	return file_errno_proto_rawDescGZIP(), []int{0}
}

var File_errno_proto protoreflect.FileDescriptor

var file_errno_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x72, 0x72, 0x6e, 0x6f, 0x2a, 0xce, 0x09, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x6f, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x6f, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0a, 0x42, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x10, 0x90, 0x4e, 0x12, 0x0e, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x72, 0x72, 0x10, 0x91, 0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x92, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x54, 0x6f, 0x6f,
	0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x93, 0x4e, 0x12, 0x10,
	0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x10, 0xa0, 0x9c, 0x01,
	0x12, 0x17, 0x0a, 0x11, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x53,
	0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xb0, 0xea, 0x01, 0x12, 0x14, 0x0a, 0x0e, 0x50, 0x6d, 0x73,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xb1, 0xea, 0x01, 0x12,
	0x1a, 0x0a, 0x14, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x50, 0x69,
	0x63, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xb2, 0xea, 0x01, 0x12, 0x17, 0x0a, 0x11, 0x50,
	0x6d, 0x73, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x50, 0x69, 0x63, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72,
	0x10, 0xb3, 0xea, 0x01, 0x12, 0x17, 0x0a, 0x11, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xb4, 0xea, 0x01, 0x12, 0x14, 0x0a,
	0x0e, 0x50, 0x6d, 0x73, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10,
	0xb5, 0xea, 0x01, 0x12, 0x19, 0x0a, 0x13, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xb6, 0xea, 0x01, 0x12, 0x16,
	0x0a, 0x10, 0x50, 0x6d, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x72, 0x76, 0x45,
	0x72, 0x72, 0x10, 0xb7, 0xea, 0x01, 0x12, 0x1e, 0x0a, 0x18, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x72, 0x76, 0x45,
	0x72, 0x72, 0x10, 0xb8, 0xea, 0x01, 0x12, 0x1b, 0x0a, 0x15, 0x50, 0x6d, 0x73, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10,
	0xb9, 0xea, 0x01, 0x12, 0x1d, 0x0a, 0x17, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xba,
	0xea, 0x01, 0x12, 0x1a, 0x0a, 0x14, 0x50, 0x6d, 0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xbb, 0xea, 0x01, 0x12, 0x19,
	0x0a, 0x13, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xbc, 0xea, 0x01, 0x12, 0x16, 0x0a, 0x10, 0x50, 0x6d, 0x73,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xbd, 0xea,
	0x01, 0x12, 0x1d, 0x0a, 0x17, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xbe, 0xea, 0x01,
	0x12, 0x1a, 0x0a, 0x14, 0x50, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xbf, 0xea, 0x01, 0x12, 0x2a, 0x0a, 0x24,
	0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x72,
	0x76, 0x45, 0x72, 0x72, 0x10, 0xc0, 0xea, 0x01, 0x12, 0x27, 0x0a, 0x21, 0x50, 0x6d, 0x73, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xc1, 0xea,
	0x01, 0x12, 0x27, 0x0a, 0x21, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xc2, 0xea, 0x01, 0x12, 0x24, 0x0a, 0x1e, 0x50, 0x6d,
	0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xc3, 0xea, 0x01,
	0x12, 0x21, 0x0a, 0x1b, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10,
	0xc4, 0xea, 0x01, 0x12, 0x1e, 0x0a, 0x18, 0x50, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10,
	0xc5, 0xea, 0x01, 0x12, 0x32, 0x0a, 0x2c, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x72, 0x76,
	0x45, 0x72, 0x72, 0x10, 0xc6, 0xea, 0x01, 0x12, 0x2f, 0x0a, 0x29, 0x50, 0x6d, 0x73, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x72,
	0x76, 0x45, 0x72, 0x72, 0x10, 0xc7, 0xea, 0x01, 0x12, 0x26, 0x0a, 0x20, 0x52, 0x50, 0x43, 0x50,
	0x6d, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xc8, 0xea, 0x01,
	0x12, 0x23, 0x0a, 0x1d, 0x50, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x75,
	0x6c, 0x6c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x72, 0x76, 0x45, 0x72,
	0x72, 0x10, 0xc9, 0xea, 0x01, 0x12, 0x1f, 0x0a, 0x19, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x61, 0x64, 0x64, 0x65, 0x72, 0x53, 0x72, 0x76, 0x45,
	0x72, 0x72, 0x10, 0xca, 0xea, 0x01, 0x12, 0x1c, 0x0a, 0x16, 0x50, 0x6d, 0x73, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x61, 0x64, 0x64, 0x65, 0x72, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72,
	0x10, 0xcb, 0xea, 0x01, 0x12, 0x25, 0x0a, 0x1f, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xcc, 0xea, 0x01, 0x12, 0x22, 0x0a, 0x1c, 0x50,
	0x6d, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xcd, 0xea, 0x01, 0x12,
	0x1a, 0x0a, 0x14, 0x52, 0x50, 0x43, 0x50, 0x6d, 0x73, 0x53, 0x6b, 0x75, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72, 0x10, 0xce, 0xea, 0x01, 0x12, 0x17, 0x0a, 0x11, 0x50,
	0x6d, 0x73, 0x53, 0x6b, 0x75, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x72, 0x76, 0x45, 0x72, 0x72,
	0x10, 0xcf, 0xea, 0x01, 0x12, 0x14, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f,
	0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x80, 0xf1, 0x04, 0x12, 0x18, 0x0a, 0x12, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x10, 0x81, 0xf1, 0x04, 0x12, 0x0f, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x74, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x10, 0x83, 0xf1, 0x04, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x61, 0x6e, 0x79, 0x70, 0x38, 0x2f, 0x77, 0x78, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6b, 0x69,
	0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x65, 0x72, 0x72,
	0x6e, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errno_proto_rawDescOnce sync.Once
	file_errno_proto_rawDescData = file_errno_proto_rawDesc
)

func file_errno_proto_rawDescGZIP() []byte {
	file_errno_proto_rawDescOnce.Do(func() {
		file_errno_proto_rawDescData = protoimpl.X.CompressGZIP(file_errno_proto_rawDescData)
	})
	return file_errno_proto_rawDescData
}

var file_errno_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errno_proto_goTypes = []interface{}{
	(Err)(0), // 0: errno.Err
}
var file_errno_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errno_proto_init() }
func file_errno_proto_init() {
	if File_errno_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errno_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errno_proto_goTypes,
		DependencyIndexes: file_errno_proto_depIdxs,
		EnumInfos:         file_errno_proto_enumTypes,
	}.Build()
	File_errno_proto = out.File
	file_errno_proto_rawDesc = nil
	file_errno_proto_goTypes = nil
	file_errno_proto_depIdxs = nil
}

var _ context.Context