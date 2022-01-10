// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: embedfi/finance/v1/validate.proto

package finance

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Rules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScaledAmount *ScaledAmountRules `protobuf:"bytes,1,opt,name=scaled_amount,json=scaledAmount,proto3" json:"scaled_amount,omitempty"`
}

func (x *Rules) Reset() {
	*x = Rules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_embedfi_finance_v1_validate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rules) ProtoMessage() {}

func (x *Rules) ProtoReflect() protoreflect.Message {
	mi := &file_embedfi_finance_v1_validate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rules.ProtoReflect.Descriptor instead.
func (*Rules) Descriptor() ([]byte, []int) {
	return file_embedfi_finance_v1_validate_proto_rawDescGZIP(), []int{0}
}

func (x *Rules) GetScaledAmount() *ScaledAmountRules {
	if x != nil {
		return x.ScaledAmount
	}
	return nil
}

type ScaledAmountRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lt         *ScaledAmount          `protobuf:"bytes,2,opt,name=lt,proto3" json:"lt,omitempty"`
	Lte        *ScaledAmount          `protobuf:"bytes,3,opt,name=lte,proto3" json:"lte,omitempty"`
	Gt         *ScaledAmount          `protobuf:"bytes,4,opt,name=gt,proto3" json:"gt,omitempty"`
	Gte        *ScaledAmount          `protobuf:"bytes,5,opt,name=gte,proto3" json:"gte,omitempty"`
	MaxScale   *wrapperspb.Int32Value `protobuf:"bytes,6,opt,name=max_scale,json=maxScale,proto3" json:"max_scale,omitempty"`
	MinScale   *wrapperspb.Int32Value `protobuf:"bytes,7,opt,name=min_scale,json=minScale,proto3" json:"min_scale,omitempty"`
	ExactScale *wrapperspb.Int32Value `protobuf:"bytes,8,opt,name=exact_scale,json=exactScale,proto3" json:"exact_scale,omitempty"`
}

func (x *ScaledAmountRules) Reset() {
	*x = ScaledAmountRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_embedfi_finance_v1_validate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScaledAmountRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScaledAmountRules) ProtoMessage() {}

func (x *ScaledAmountRules) ProtoReflect() protoreflect.Message {
	mi := &file_embedfi_finance_v1_validate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScaledAmountRules.ProtoReflect.Descriptor instead.
func (*ScaledAmountRules) Descriptor() ([]byte, []int) {
	return file_embedfi_finance_v1_validate_proto_rawDescGZIP(), []int{1}
}

func (x *ScaledAmountRules) GetLt() *ScaledAmount {
	if x != nil {
		return x.Lt
	}
	return nil
}

func (x *ScaledAmountRules) GetLte() *ScaledAmount {
	if x != nil {
		return x.Lte
	}
	return nil
}

func (x *ScaledAmountRules) GetGt() *ScaledAmount {
	if x != nil {
		return x.Gt
	}
	return nil
}

func (x *ScaledAmountRules) GetGte() *ScaledAmount {
	if x != nil {
		return x.Gte
	}
	return nil
}

func (x *ScaledAmountRules) GetMaxScale() *wrapperspb.Int32Value {
	if x != nil {
		return x.MaxScale
	}
	return nil
}

func (x *ScaledAmountRules) GetMinScale() *wrapperspb.Int32Value {
	if x != nil {
		return x.MinScale
	}
	return nil
}

func (x *ScaledAmountRules) GetExactScale() *wrapperspb.Int32Value {
	if x != nil {
		return x.ExactScale
	}
	return nil
}

var file_embedfi_finance_v1_validate_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Rules)(nil),
		Field:         92650872,
		Name:          "embedfi.finance.v1.rules",
		Tag:           "bytes,92650872,opt,name=rules",
		Filename:      "embedfi/finance/v1/validate.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional embedfi.finance.v1.Rules rules = 92650872;
	E_Rules = &file_embedfi_finance_v1_validate_proto_extTypes[0]
)

var File_embedfi_finance_v1_validate_proto protoreflect.FileDescriptor

var file_embedfi_finance_v1_validate_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x66, 0x69, 0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x66, 0x69, 0x2e, 0x66, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x26, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x66, 0x69,
	0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x53, 0x0a, 0x05, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0d, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x66, 0x69, 0x2e, 0x66, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x0c, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x91, 0x03, 0x0a, 0x11, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x02,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x62, 0x65, 0x64,
	0x66, 0x69, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63,
	0x61, 0x6c, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x02, 0x6c, 0x74, 0x12, 0x32,
	0x0a, 0x03, 0x6c, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d,
	0x62, 0x65, 0x64, 0x66, 0x69, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x03, 0x6c,
	0x74, 0x65, 0x12, 0x30, 0x0a, 0x02, 0x67, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x66, 0x69, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x02, 0x67, 0x74, 0x12, 0x32, 0x0a, 0x03, 0x67, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x66, 0x69, 0x2e, 0x66, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x03, 0x67, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0b,
	0x65, 0x78, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a,
	0x65, 0x78, 0x61, 0x63, 0x74, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x3a, 0x54, 0x0a, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xf8, 0xfa, 0x96, 0x2c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6d,
	0x62, 0x65, 0x64, 0x66, 0x69, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6d, 0x62, 0x65, 0x64, 0x66, 0x69, 0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x3b, 0x66,
	0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_embedfi_finance_v1_validate_proto_rawDescOnce sync.Once
	file_embedfi_finance_v1_validate_proto_rawDescData = file_embedfi_finance_v1_validate_proto_rawDesc
)

func file_embedfi_finance_v1_validate_proto_rawDescGZIP() []byte {
	file_embedfi_finance_v1_validate_proto_rawDescOnce.Do(func() {
		file_embedfi_finance_v1_validate_proto_rawDescData = protoimpl.X.CompressGZIP(file_embedfi_finance_v1_validate_proto_rawDescData)
	})
	return file_embedfi_finance_v1_validate_proto_rawDescData
}

var file_embedfi_finance_v1_validate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_embedfi_finance_v1_validate_proto_goTypes = []interface{}{
	(*Rules)(nil),                     // 0: embedfi.finance.v1.Rules
	(*ScaledAmountRules)(nil),         // 1: embedfi.finance.v1.ScaledAmountRules
	(*ScaledAmount)(nil),              // 2: embedfi.finance.v1.ScaledAmount
	(*wrapperspb.Int32Value)(nil),     // 3: google.protobuf.Int32Value
	(*descriptorpb.FieldOptions)(nil), // 4: google.protobuf.FieldOptions
}
var file_embedfi_finance_v1_validate_proto_depIdxs = []int32{
	1,  // 0: embedfi.finance.v1.Rules.scaled_amount:type_name -> embedfi.finance.v1.ScaledAmountRules
	2,  // 1: embedfi.finance.v1.ScaledAmountRules.lt:type_name -> embedfi.finance.v1.ScaledAmount
	2,  // 2: embedfi.finance.v1.ScaledAmountRules.lte:type_name -> embedfi.finance.v1.ScaledAmount
	2,  // 3: embedfi.finance.v1.ScaledAmountRules.gt:type_name -> embedfi.finance.v1.ScaledAmount
	2,  // 4: embedfi.finance.v1.ScaledAmountRules.gte:type_name -> embedfi.finance.v1.ScaledAmount
	3,  // 5: embedfi.finance.v1.ScaledAmountRules.max_scale:type_name -> google.protobuf.Int32Value
	3,  // 6: embedfi.finance.v1.ScaledAmountRules.min_scale:type_name -> google.protobuf.Int32Value
	3,  // 7: embedfi.finance.v1.ScaledAmountRules.exact_scale:type_name -> google.protobuf.Int32Value
	4,  // 8: embedfi.finance.v1.rules:extendee -> google.protobuf.FieldOptions
	0,  // 9: embedfi.finance.v1.rules:type_name -> embedfi.finance.v1.Rules
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	9,  // [9:10] is the sub-list for extension type_name
	8,  // [8:9] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_embedfi_finance_v1_validate_proto_init() }
func file_embedfi_finance_v1_validate_proto_init() {
	if File_embedfi_finance_v1_validate_proto != nil {
		return
	}
	file_embedfi_finance_v1_scaled_amount_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_embedfi_finance_v1_validate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rules); i {
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
		file_embedfi_finance_v1_validate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScaledAmountRules); i {
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
			RawDescriptor: file_embedfi_finance_v1_validate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_embedfi_finance_v1_validate_proto_goTypes,
		DependencyIndexes: file_embedfi_finance_v1_validate_proto_depIdxs,
		MessageInfos:      file_embedfi_finance_v1_validate_proto_msgTypes,
		ExtensionInfos:    file_embedfi_finance_v1_validate_proto_extTypes,
	}.Build()
	File_embedfi_finance_v1_validate_proto = out.File
	file_embedfi_finance_v1_validate_proto_rawDesc = nil
	file_embedfi_finance_v1_validate_proto_goTypes = nil
	file_embedfi_finance_v1_validate_proto_depIdxs = nil
}
