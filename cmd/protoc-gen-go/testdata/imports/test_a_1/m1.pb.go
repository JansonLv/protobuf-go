// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_a_1/m1.proto

package test_a_1

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type E1 int32

const (
	E1_E1_ZERO E1 = 0
)

// Deprecated: Use E1.Type.Values instead.
var E1_name = map[int32]string{
	0: "E1_ZERO",
}

// Deprecated: Use E1.Type.Values instead.
var E1_value = map[string]int32{
	"E1_ZERO": 0,
}

func (x E1) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

func (E1) Type() protoreflect.EnumType {
	return xxx_File_imports_test_a_1_m1_proto_enumTypes[0]
}

func (x E1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use E1.Type instead.
func (E1) EnumDescriptor() ([]byte, []int) {
	return xxx_File_imports_test_a_1_m1_proto_rawDescGZIP(), []int{0}
}

type M1 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *M1) Reset() {
	*x = M1{}
}

func (x *M1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M1) ProtoMessage() {}

func (x *M1) ProtoReflect() protoreflect.Message {
	return xxx_File_imports_test_a_1_m1_proto_messageTypes[0].MessageOf(x)
}

// Deprecated: Use M1.ProtoReflect.Type instead.
func (*M1) Descriptor() ([]byte, []int) {
	return xxx_File_imports_test_a_1_m1_proto_rawDescGZIP(), []int{0}
}

type M1_1 struct {
	M1                   *M1      `protobuf:"bytes,1,opt,name=m1,proto3" json:"m1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *M1_1) Reset() {
	*x = M1_1{}
}

func (x *M1_1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M1_1) ProtoMessage() {}

func (x *M1_1) ProtoReflect() protoreflect.Message {
	return xxx_File_imports_test_a_1_m1_proto_messageTypes[1].MessageOf(x)
}

// Deprecated: Use M1_1.ProtoReflect.Type instead.
func (*M1_1) Descriptor() ([]byte, []int) {
	return xxx_File_imports_test_a_1_m1_proto_rawDescGZIP(), []int{1}
}

func (x *M1_1) GetM1() *M1 {
	if x != nil {
		return x.M1
	}
	return nil
}

var File_imports_test_a_1_m1_proto protoreflect.FileDescriptor

var xxx_File_imports_test_a_1_m1_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x5f, 0x31, 0x2f, 0x6d, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x22, 0x04, 0x0a, 0x02, 0x4d, 0x31, 0x22, 0x22, 0x0a, 0x04, 0x4d, 0x31, 0x5f,
	0x31, 0x12, 0x1a, 0x0a, 0x02, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x2e, 0x4d, 0x31, 0x52, 0x02, 0x6d, 0x31, 0x2a, 0x11, 0x0a,
	0x02, 0x45, 0x31, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x31, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00,
	0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	xxx_File_imports_test_a_1_m1_proto_rawDesc_once sync.Once
	xxx_File_imports_test_a_1_m1_proto_rawDesc_data = xxx_File_imports_test_a_1_m1_proto_rawDesc
)

func xxx_File_imports_test_a_1_m1_proto_rawDescGZIP() []byte {
	xxx_File_imports_test_a_1_m1_proto_rawDesc_once.Do(func() {
		xxx_File_imports_test_a_1_m1_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_imports_test_a_1_m1_proto_rawDesc_data)
	})
	return xxx_File_imports_test_a_1_m1_proto_rawDesc_data
}

var xxx_File_imports_test_a_1_m1_proto_enumTypes = make([]protoreflect.EnumType, 1)
var xxx_File_imports_test_a_1_m1_proto_messageTypes = make([]protoimpl.MessageType, 2)
var xxx_File_imports_test_a_1_m1_proto_goTypes = []interface{}{
	(E1)(0),      // 0: test.a.E1
	(*M1)(nil),   // 1: test.a.M1
	(*M1_1)(nil), // 2: test.a.M1_1
}
var xxx_File_imports_test_a_1_m1_proto_depIdxs = []int32{
	1, // test.a.M1_1.m1:type_name -> test.a.M1
}

func init() { xxx_File_imports_test_a_1_m1_proto_init() }
func xxx_File_imports_test_a_1_m1_proto_init() {
	if File_imports_test_a_1_m1_proto != nil {
		return
	}
	File_imports_test_a_1_m1_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_imports_test_a_1_m1_proto_rawDesc,
		GoTypes:            xxx_File_imports_test_a_1_m1_proto_goTypes,
		DependencyIndexes:  xxx_File_imports_test_a_1_m1_proto_depIdxs,
		EnumOutputTypes:    xxx_File_imports_test_a_1_m1_proto_enumTypes,
		MessageOutputTypes: xxx_File_imports_test_a_1_m1_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_imports_test_a_1_m1_proto_rawDesc = nil
	xxx_File_imports_test_a_1_m1_proto_goTypes = nil
	xxx_File_imports_test_a_1_m1_proto_depIdxs = nil
}
