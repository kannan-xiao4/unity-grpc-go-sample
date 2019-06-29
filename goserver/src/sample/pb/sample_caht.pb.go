// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample_caht.proto

package sample_caht

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cb3f5f5f64d1c4, []int{0}
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

type Name struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cb3f5f5f64d1c4, []int{1}
}

func (m *Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Name.Unmarshal(m, b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Name.Marshal(b, m, deterministic)
}
func (m *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(m, src)
}
func (m *Name) XXX_Size() int {
	return xxx_messageInfo_Name.Size(m)
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Replay struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Replay) Reset()         { *m = Replay{} }
func (m *Replay) String() string { return proto.CompactTextString(m) }
func (*Replay) ProtoMessage()    {}
func (*Replay) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cb3f5f5f64d1c4, []int{2}
}

func (m *Replay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Replay.Unmarshal(m, b)
}
func (m *Replay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Replay.Marshal(b, m, deterministic)
}
func (m *Replay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Replay.Merge(m, src)
}
func (m *Replay) XXX_Size() int {
	return xxx_messageInfo_Replay.Size(m)
}
func (m *Replay) XXX_DiscardUnknown() {
	xxx_messageInfo_Replay.DiscardUnknown(m)
}

var xxx_messageInfo_Replay proto.InternalMessageInfo

func (m *Replay) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Summary struct {
	Replaies             []*Replay `protobuf:"bytes,1,rep,name=replaies,proto3" json:"replaies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Summary) Reset()         { *m = Summary{} }
func (m *Summary) String() string { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()    {}
func (*Summary) Descriptor() ([]byte, []int) {
	return fileDescriptor_83cb3f5f5f64d1c4, []int{3}
}

func (m *Summary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary.Unmarshal(m, b)
}
func (m *Summary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary.Marshal(b, m, deterministic)
}
func (m *Summary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary.Merge(m, src)
}
func (m *Summary) XXX_Size() int {
	return xxx_messageInfo_Summary.Size(m)
}
func (m *Summary) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary.DiscardUnknown(m)
}

var xxx_messageInfo_Summary proto.InternalMessageInfo

func (m *Summary) GetReplaies() []*Replay {
	if m != nil {
		return m.Replaies
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "sample.chat.Empty")
	proto.RegisterType((*Name)(nil), "sample.chat.Name")
	proto.RegisterType((*Replay)(nil), "sample.chat.Replay")
	proto.RegisterType((*Summary)(nil), "sample.chat.Summary")
}

func init() { proto.RegisterFile("sample_caht.proto", fileDescriptor_83cb3f5f5f64d1c4) }

var fileDescriptor_83cb3f5f5f64d1c4 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x6b, 0xd1, 0x36, 0xed, 0x45, 0x0c, 0x3d, 0x18, 0xa2, 0x4e, 0x95, 0xa7, 0x4c, 0xa1,
	0x94, 0xa5, 0xea, 0x08, 0x42, 0x65, 0x40, 0x0c, 0x66, 0x63, 0x41, 0x26, 0x3d, 0x35, 0x91, 0xe2,
	0x3a, 0xb2, 0x2f, 0x43, 0x7e, 0x35, 0x7f, 0x01, 0x35, 0xa1, 0xa8, 0x91, 0x82, 0xd4, 0xcd, 0x3e,
	0x7d, 0xef, 0xee, 0x3d, 0x3d, 0x98, 0x79, 0x6d, 0xca, 0x82, 0x3e, 0x53, 0x9d, 0x71, 0x52, 0x3a,
	0xcb, 0x16, 0xc3, 0x76, 0x94, 0xa4, 0x99, 0x66, 0x19, 0xc0, 0xe8, 0xd9, 0x94, 0x5c, 0xcb, 0x39,
	0x0c, 0xdf, 0xb4, 0x21, 0x44, 0x18, 0x1e, 0xb4, 0xa1, 0x48, 0x2c, 0x44, 0x3c, 0x55, 0xcd, 0x5b,
	0x4a, 0x18, 0x2b, 0x2a, 0x0b, 0x5d, 0x63, 0x04, 0x81, 0x21, 0xef, 0xf5, 0xfe, 0x04, 0x9c, 0xbe,
	0x72, 0x03, 0xc1, 0x7b, 0x65, 0x8c, 0x76, 0x35, 0xde, 0xc1, 0xc4, 0x1d, 0xf1, 0x9c, 0x7c, 0x24,
	0x16, 0x57, 0x71, 0xb8, 0xba, 0x49, 0xce, 0x6e, 0x26, 0xed, 0x2e, 0xf5, 0x07, 0xad, 0xbe, 0x05,
	0x80, 0xb2, 0x15, 0xd3, 0xb6, 0xca, 0x77, 0x84, 0xf7, 0x30, 0x7a, 0xa1, 0xa2, 0xb0, 0x38, 0xeb,
	0xc8, 0x8e, 0xf6, 0xe6, 0x7d, 0x9b, 0xe4, 0x00, 0xd7, 0x30, 0x7d, 0xcd, 0x3d, 0x6f, 0x1d, 0x11,
	0x23, 0x76, 0x98, 0x26, 0xde, 0x3f, 0xba, 0xa5, 0xc0, 0x0d, 0x84, 0x8a, 0x52, 0xeb, 0x76, 0xad,
	0xb6, 0xe7, 0xe4, 0x6d, 0x67, 0xf4, 0x1b, 0x52, 0x0e, 0x62, 0x81, 0x6b, 0x98, 0x3c, 0x65, 0x9a,
	0x39, 0x3f, 0xec, 0x2f, 0xf7, 0x1a, 0x8b, 0xa5, 0x78, 0xbc, 0xfe, 0x08, 0xcf, 0x8a, 0xf9, 0x1a,
	0x37, 0xcd, 0x3c, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x17, 0x4a, 0x0a, 0xf7, 0xae, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteGuideClient interface {
	Hello(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Replay, error)
	ListGreet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (RouteGuide_ListGreetClient, error)
	RecordGreet(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordGreetClient, error)
	Chatting(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_ChattingClient, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) Hello(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Replay, error) {
	out := new(Replay)
	err := c.cc.Invoke(ctx, "/sample.chat.RouteGuide/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) ListGreet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (RouteGuide_ListGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[0], "/sample.chat.RouteGuide/ListGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideListGreetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListGreetClient interface {
	Recv() (*Replay, error)
	grpc.ClientStream
}

type routeGuideListGreetClient struct {
	grpc.ClientStream
}

func (x *routeGuideListGreetClient) Recv() (*Replay, error) {
	m := new(Replay)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RecordGreet(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[1], "/sample.chat.RouteGuide/RecordGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordGreetClient{stream}
	return x, nil
}

type RouteGuide_RecordGreetClient interface {
	Send(*Name) error
	CloseAndRecv() (*Summary, error)
	grpc.ClientStream
}

type routeGuideRecordGreetClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordGreetClient) Send(m *Name) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordGreetClient) CloseAndRecv() (*Summary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Summary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) Chatting(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_ChattingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteGuide_serviceDesc.Streams[2], "/sample.chat.RouteGuide/Chatting", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideChattingClient{stream}
	return x, nil
}

type RouteGuide_ChattingClient interface {
	Send(*Name) error
	Recv() (*Replay, error)
	grpc.ClientStream
}

type routeGuideChattingClient struct {
	grpc.ClientStream
}

func (x *routeGuideChattingClient) Send(m *Name) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideChattingClient) Recv() (*Replay, error) {
	m := new(Replay)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideServer is the server API for RouteGuide service.
type RouteGuideServer interface {
	Hello(context.Context, *Name) (*Replay, error)
	ListGreet(*Empty, RouteGuide_ListGreetServer) error
	RecordGreet(RouteGuide_RecordGreetServer) error
	Chatting(RouteGuide_ChattingServer) error
}

// UnimplementedRouteGuideServer can be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (*UnimplementedRouteGuideServer) Hello(ctx context.Context, req *Name) (*Replay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedRouteGuideServer) ListGreet(req *Empty, srv RouteGuide_ListGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method ListGreet not implemented")
}
func (*UnimplementedRouteGuideServer) RecordGreet(srv RouteGuide_RecordGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordGreet not implemented")
}
func (*UnimplementedRouteGuideServer) Chatting(srv RouteGuide_ChattingServer) error {
	return status.Errorf(codes.Unimplemented, "method Chatting not implemented")
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.chat.RouteGuide/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).Hello(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_ListGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).ListGreet(m, &routeGuideListGreetServer{stream})
}

type RouteGuide_ListGreetServer interface {
	Send(*Replay) error
	grpc.ServerStream
}

type routeGuideListGreetServer struct {
	grpc.ServerStream
}

func (x *routeGuideListGreetServer) Send(m *Replay) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_RecordGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RecordGreet(&routeGuideRecordGreetServer{stream})
}

type RouteGuide_RecordGreetServer interface {
	SendAndClose(*Summary) error
	Recv() (*Name, error)
	grpc.ServerStream
}

type routeGuideRecordGreetServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordGreetServer) SendAndClose(m *Summary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordGreetServer) Recv() (*Name, error) {
	m := new(Name)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_Chatting_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).Chatting(&routeGuideChattingServer{stream})
}

type RouteGuide_ChattingServer interface {
	Send(*Replay) error
	Recv() (*Name, error)
	grpc.ServerStream
}

type routeGuideChattingServer struct {
	grpc.ServerStream
}

func (x *routeGuideChattingServer) Send(m *Replay) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideChattingServer) Recv() (*Name, error) {
	m := new(Name)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sample.chat.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _RouteGuide_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListGreet",
			Handler:       _RouteGuide_ListGreet_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordGreet",
			Handler:       _RouteGuide_RecordGreet_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Chatting",
			Handler:       _RouteGuide_Chatting_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sample_caht.proto",
}
