// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rmtaskmgt_service.proto

package gRPC

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

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

type WorkRequest struct {
	MessageId            string   `protobuf:"bytes,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	UserID               int64    `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	StartFrequency       int64    `protobuf:"varint,3,opt,name=startFrequency,proto3" json:"startFrequency,omitempty"`
	PackageSize          int64    `protobuf:"varint,4,opt,name=packageSize,proto3" json:"packageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkRequest) Reset()         { *m = WorkRequest{} }
func (m *WorkRequest) String() string { return proto.CompactTextString(m) }
func (*WorkRequest) ProtoMessage()    {}
func (*WorkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08494d42a998037, []int{0}
}

func (m *WorkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkRequest.Unmarshal(m, b)
}
func (m *WorkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkRequest.Marshal(b, m, deterministic)
}
func (m *WorkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkRequest.Merge(m, src)
}
func (m *WorkRequest) XXX_Size() int {
	return xxx_messageInfo_WorkRequest.Size(m)
}
func (m *WorkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkRequest proto.InternalMessageInfo

func (m *WorkRequest) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *WorkRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *WorkRequest) GetStartFrequency() int64 {
	if m != nil {
		return m.StartFrequency
	}
	return 0
}

func (m *WorkRequest) GetPackageSize() int64 {
	if m != nil {
		return m.PackageSize
	}
	return 0
}

type WorkReply struct {
	WorkID               string   `protobuf:"bytes,1,opt,name=workID,proto3" json:"workID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkReply) Reset()         { *m = WorkReply{} }
func (m *WorkReply) String() string { return proto.CompactTextString(m) }
func (*WorkReply) ProtoMessage()    {}
func (*WorkReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08494d42a998037, []int{1}
}

func (m *WorkReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkReply.Unmarshal(m, b)
}
func (m *WorkReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkReply.Marshal(b, m, deterministic)
}
func (m *WorkReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkReply.Merge(m, src)
}
func (m *WorkReply) XXX_Size() int {
	return xxx_messageInfo_WorkReply.Size(m)
}
func (m *WorkReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkReply.DiscardUnknown(m)
}

var xxx_messageInfo_WorkReply proto.InternalMessageInfo

func (m *WorkReply) GetWorkID() string {
	if m != nil {
		return m.WorkID
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkRequest)(nil), "gRPC.WorkRequest")
	proto.RegisterType((*WorkReply)(nil), "gRPC.WorkReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RadioManageClient is the client API for RadioManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RadioManageClient interface {
	Tasking(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkReply, error)
}

type radioManageClient struct {
	cc *grpc.ClientConn
}

func NewRadioManageClient(cc *grpc.ClientConn) RadioManageClient {
	return &radioManageClient{cc}
}

func (c *radioManageClient) Tasking(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkReply, error) {
	out := new(WorkReply)
	err := c.cc.Invoke(ctx, "/gRPC.RadioManage/Tasking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RadioManageServer is the server API for RadioManage service.
type RadioManageServer interface {
	Tasking(context.Context, *WorkRequest) (*WorkReply, error)
}

func RegisterRadioManageServer(s *grpc.Server, srv RadioManageServer) {
	s.RegisterService(&_RadioManage_serviceDesc, srv)
}

func _RadioManage_Tasking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadioManageServer).Tasking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC.RadioManage/Tasking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadioManageServer).Tasking(ctx, req.(*WorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RadioManage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gRPC.RadioManage",
	HandlerType: (*RadioManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tasking",
			Handler:    _RadioManage_Tasking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rmtaskmgt_service.proto",
}

func init() { proto.RegisterFile("rmtaskmgt_service.proto", fileDescriptor_c08494d42a998037) }

var fileDescriptor_c08494d42a998037 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xca, 0x2d, 0x49,
	0x2c, 0xce, 0xce, 0x4d, 0x2f, 0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x0f, 0x0a, 0x70, 0x96, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf,
	0x2b, 0x86, 0xa8, 0x51, 0xea, 0x65, 0xe4, 0xe2, 0x0e, 0xcf, 0x2f, 0xca, 0x0e, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe1, 0xe2, 0xcc, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xf5, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0x89, 0x71, 0xb1, 0x95, 0x16, 0xa7,
	0x16, 0x79, 0xba, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x07, 0x41, 0x79, 0x42, 0x6a, 0x5c, 0x7c,
	0xc5, 0x25, 0x89, 0x45, 0x25, 0x6e, 0x45, 0x20, 0x63, 0xf2, 0x92, 0x2b, 0x25, 0x98, 0xc1, 0xf2,
	0x68, 0xa2, 0x42, 0x0a, 0x5c, 0xdc, 0x05, 0x89, 0xc9, 0xd9, 0x89, 0xe9, 0xa9, 0xc1, 0x99, 0x55,
	0xa9, 0x12, 0x2c, 0x60, 0x45, 0xc8, 0x42, 0x4a, 0xca, 0x5c, 0x9c, 0x10, 0xe7, 0x14, 0xe4, 0x54,
	0x82, 0xac, 0x2b, 0xcf, 0x2f, 0xca, 0xf6, 0x74, 0x81, 0xba, 0x04, 0xca, 0x33, 0xca, 0xe6, 0xe2,
	0x0e, 0x4a, 0x4c, 0xc9, 0xcc, 0xf7, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x15, 0x8a, 0xe1, 0x62, 0x0f,
	0x49, 0x2c, 0xce, 0xce, 0xcc, 0x4b, 0x17, 0x12, 0xd4, 0x03, 0xf9, 0x59, 0x0f, 0xc9, 0x47, 0x52,
	0xfc, 0xc8, 0x42, 0x05, 0x39, 0x95, 0x4a, 0xba, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x52, 0x57, 0xe2,
	0xd7, 0x2f, 0x33, 0xd4, 0xaf, 0x86, 0x7b, 0xae, 0xd6, 0x8a, 0x51, 0x2b, 0x4a, 0x50, 0x08, 0x5d,
	0x34, 0x89, 0x0d, 0x1c, 0x50, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xee, 0x3c, 0xb9,
	0x67, 0x01, 0x00, 0x00,
}
