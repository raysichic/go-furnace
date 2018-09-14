// Code generated by protoc-gen-go. DO NOT EDIT.
// source: furnace.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Stack struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stack) Reset()         { *m = Stack{} }
func (m *Stack) String() string { return proto.CompactTextString(m) }
func (*Stack) ProtoMessage()    {}
func (*Stack) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fe409196414abd, []int{0}
}
func (m *Stack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stack.Unmarshal(m, b)
}
func (m *Stack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stack.Marshal(b, m, deterministic)
}
func (m *Stack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stack.Merge(m, src)
}
func (m *Stack) XXX_Size() int {
	return xxx_messageInfo_Stack.Size(m)
}
func (m *Stack) XXX_DiscardUnknown() {
	xxx_messageInfo_Stack.DiscardUnknown(m)
}

var xxx_messageInfo_Stack proto.InternalMessageInfo

func (m *Stack) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Proceed struct {
	Failed               bool     `protobuf:"varint,1,opt,name=failed,proto3" json:"failed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proceed) Reset()         { *m = Proceed{} }
func (m *Proceed) String() string { return proto.CompactTextString(m) }
func (*Proceed) ProtoMessage()    {}
func (*Proceed) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fe409196414abd, []int{1}
}
func (m *Proceed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proceed.Unmarshal(m, b)
}
func (m *Proceed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proceed.Marshal(b, m, deterministic)
}
func (m *Proceed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proceed.Merge(m, src)
}
func (m *Proceed) XXX_Size() int {
	return xxx_messageInfo_Proceed.Size(m)
}
func (m *Proceed) XXX_DiscardUnknown() {
	xxx_messageInfo_Proceed.DiscardUnknown(m)
}

var xxx_messageInfo_Proceed proto.InternalMessageInfo

func (m *Proceed) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1fe409196414abd, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Stack)(nil), "proto.Stack")
	proto.RegisterType((*Proceed)(nil), "proto.Proceed")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PreBuildClient is the client API for PreBuild service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PreBuildClient interface {
	Execute(ctx context.Context, in *Stack, opts ...grpc.CallOption) (*Proceed, error)
}

type preBuildClient struct {
	cc *grpc.ClientConn
}

func NewPreBuildClient(cc *grpc.ClientConn) PreBuildClient {
	return &preBuildClient{cc}
}

func (c *preBuildClient) Execute(ctx context.Context, in *Stack, opts ...grpc.CallOption) (*Proceed, error) {
	out := new(Proceed)
	err := c.cc.Invoke(ctx, "/proto.PreBuild/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PreBuildServer is the server API for PreBuild service.
type PreBuildServer interface {
	Execute(context.Context, *Stack) (*Proceed, error)
}

func RegisterPreBuildServer(s *grpc.Server, srv PreBuildServer) {
	s.RegisterService(&_PreBuild_serviceDesc, srv)
}

func _PreBuild_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stack)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreBuildServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PreBuild/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreBuildServer).Execute(ctx, req.(*Stack))
	}
	return interceptor(ctx, in, info, handler)
}

var _PreBuild_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PreBuild",
	HandlerType: (*PreBuildServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _PreBuild_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "furnace.proto",
}

// PostBuildClient is the client API for PostBuild service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostBuildClient interface {
	Execute(ctx context.Context, in *Stack, opts ...grpc.CallOption) (*Empty, error)
}

type postBuildClient struct {
	cc *grpc.ClientConn
}

func NewPostBuildClient(cc *grpc.ClientConn) PostBuildClient {
	return &postBuildClient{cc}
}

func (c *postBuildClient) Execute(ctx context.Context, in *Stack, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.PostBuild/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostBuildServer is the server API for PostBuild service.
type PostBuildServer interface {
	Execute(context.Context, *Stack) (*Empty, error)
}

func RegisterPostBuildServer(s *grpc.Server, srv PostBuildServer) {
	s.RegisterService(&_PostBuild_serviceDesc, srv)
}

func _PostBuild_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stack)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostBuildServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostBuild/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostBuildServer).Execute(ctx, req.(*Stack))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostBuild_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PostBuild",
	HandlerType: (*PostBuildServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _PostBuild_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "furnace.proto",
}

func init() { proto.RegisterFile("furnace.proto", fileDescriptor_d1fe409196414abd) }

var fileDescriptor_d1fe409196414abd = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2b, 0x2d, 0xca,
	0x4b, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xd2, 0x5c,
	0xac, 0xc1, 0x25, 0x89, 0xc9, 0xd9, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x22, 0x17, 0x7b, 0x40, 0x51, 0x7e, 0x72, 0x6a,
	0x6a, 0x8a, 0x90, 0x18, 0x17, 0x5b, 0x5a, 0x62, 0x66, 0x4e, 0x6a, 0x0a, 0x58, 0x01, 0x47, 0x10,
	0x94, 0xa7, 0xc4, 0xce, 0xc5, 0xea, 0x9a, 0x5b, 0x50, 0x52, 0x69, 0x64, 0xcc, 0xc5, 0x11, 0x50,
	0x94, 0xea, 0x54, 0x9a, 0x99, 0x93, 0x22, 0xa4, 0xce, 0xc5, 0xee, 0x5a, 0x91, 0x9a, 0x5c, 0x5a,
	0x92, 0x2a, 0xc4, 0x03, 0xb1, 0x4e, 0x0f, 0x6c, 0x89, 0x14, 0x1f, 0x94, 0x07, 0x35, 0xd5, 0xc8,
	0x88, 0x8b, 0x33, 0x20, 0xbf, 0xb8, 0x04, 0xa2, 0x4b, 0x15, 0x97, 0x2e, 0x18, 0x0f, 0x6c, 0x51,
	0x12, 0x1b, 0x98, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x13, 0xc5, 0x9e, 0x15, 0xd0, 0x00,
	0x00, 0x00,
}