// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3/enum.proto

package proto3

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type Enum int32

const (
	Enum_ZERO Enum = 0
	Enum_ONE  Enum = 1
	Enum_TWO  Enum = 2
)

// Deprecated: Use Enum.Type.Values instead.
var Enum_name = map[int32]string{
	0: "ZERO",
	1: "ONE",
	2: "TWO",
}

// Deprecated: Use Enum.Type.Values instead.
var Enum_value = map[string]int32{
	"ZERO": 0,
	"ONE":  1,
	"TWO":  2,
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

func (Enum) Type() protoreflect.EnumType {
	return xxx_File_proto3_enum_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Type instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_proto3_enum_proto_rawDescGZIP(), []int{0}
}

var File_proto3_enum_proto protoreflect.FileDescriptor

var xxx_File_proto3_enum_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2a, 0x22, 0x0a, 0x04, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x42, 0x41,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f,
	0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	xxx_File_proto3_enum_proto_rawDesc_once sync.Once
	xxx_File_proto3_enum_proto_rawDesc_data = xxx_File_proto3_enum_proto_rawDesc
)

func xxx_File_proto3_enum_proto_rawDescGZIP() []byte {
	xxx_File_proto3_enum_proto_rawDesc_once.Do(func() {
		xxx_File_proto3_enum_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_proto3_enum_proto_rawDesc_data)
	})
	return xxx_File_proto3_enum_proto_rawDesc_data
}

var xxx_File_proto3_enum_proto_enumTypes = make([]protoreflect.EnumType, 1)
var xxx_File_proto3_enum_proto_goTypes = []interface{}{
	(Enum)(0), // 0: goproto.protoc.proto3.Enum
}
var xxx_File_proto3_enum_proto_depIdxs = []int32{}

func init() { xxx_File_proto3_enum_proto_init() }
func xxx_File_proto3_enum_proto_init() {
	if File_proto3_enum_proto != nil {
		return
	}
	File_proto3_enum_proto = protoimpl.FileBuilder{
		RawDescriptor:     xxx_File_proto3_enum_proto_rawDesc,
		GoTypes:           xxx_File_proto3_enum_proto_goTypes,
		DependencyIndexes: xxx_File_proto3_enum_proto_depIdxs,
		EnumOutputTypes:   xxx_File_proto3_enum_proto_enumTypes,
		FilesRegistry:     protoregistry.GlobalFiles,
		TypesRegistry:     protoregistry.GlobalTypes,
	}.Init()
	xxx_File_proto3_enum_proto_rawDesc = nil
	xxx_File_proto3_enum_proto_goTypes = nil
	xxx_File_proto3_enum_proto_depIdxs = nil
}
