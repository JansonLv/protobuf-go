// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/proto2.proto

package proto2

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type Message struct {
	I32                  *int32   `protobuf:"varint,1,opt,name=i32" json:"i32,omitempty"`
	M                    *Message `protobuf:"bytes,2,opt,name=m" json:"m,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *Message) Reset() {
	*x = Message{}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	return xxx_File_proto2_proto2_proto_messageTypes[0].MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Type instead.
func (*Message) Descriptor() ([]byte, []int) {
	return xxx_File_proto2_proto2_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetI32() int32 {
	if x != nil && x.I32 != nil {
		return *x.I32
	}
	return 0
}

func (x *Message) GetM() *Message {
	if x != nil {
		return x.M
	}
	return nil
}

var File_proto2_proto2_proto protoreflect.FileDescriptor

var xxx_File_proto2_proto2_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x22, 0x49, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x33, 0x32, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x33, 0x32, 0x12, 0x2c, 0x0a, 0x01, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x01, 0x6d, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
}

var (
	xxx_File_proto2_proto2_proto_rawDesc_once sync.Once
	xxx_File_proto2_proto2_proto_rawDesc_data = xxx_File_proto2_proto2_proto_rawDesc
)

func xxx_File_proto2_proto2_proto_rawDescGZIP() []byte {
	xxx_File_proto2_proto2_proto_rawDesc_once.Do(func() {
		xxx_File_proto2_proto2_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_proto2_proto2_proto_rawDesc_data)
	})
	return xxx_File_proto2_proto2_proto_rawDesc_data
}

var xxx_File_proto2_proto2_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_proto2_proto2_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: goproto.protoc.proto2.Message
}
var xxx_File_proto2_proto2_proto_depIdxs = []int32{
	0, // goproto.protoc.proto2.Message.m:type_name -> goproto.protoc.proto2.Message
}

func init() { xxx_File_proto2_proto2_proto_init() }
func xxx_File_proto2_proto2_proto_init() {
	if File_proto2_proto2_proto != nil {
		return
	}
	File_proto2_proto2_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_proto2_proto2_proto_rawDesc,
		GoTypes:            xxx_File_proto2_proto2_proto_goTypes,
		DependencyIndexes:  xxx_File_proto2_proto2_proto_depIdxs,
		MessageOutputTypes: xxx_File_proto2_proto2_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_proto2_proto2_proto_rawDesc = nil
	xxx_File_proto2_proto2_proto_goTypes = nil
	xxx_File_proto2_proto2_proto_depIdxs = nil
}
