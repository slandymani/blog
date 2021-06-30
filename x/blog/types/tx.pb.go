// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// this line is used by starport scaffolding # proto/tx/message
type MsgCreateComment struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Body    string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	PostID  string `protobuf:"bytes,3,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (m *MsgCreateComment) Reset()         { *m = MsgCreateComment{} }
func (m *MsgCreateComment) String() string { return proto.CompactTextString(m) }
func (*MsgCreateComment) ProtoMessage()    {}
func (*MsgCreateComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_90382f40c04acbba, []int{0}
}
func (m *MsgCreateComment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateComment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateComment.Merge(m, src)
}
func (m *MsgCreateComment) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateComment) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateComment.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateComment proto.InternalMessageInfo

func (m *MsgCreateComment) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateComment) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MsgCreateComment) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

type MsgCreateCommentResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateCommentResponse) Reset()         { *m = MsgCreateCommentResponse{} }
func (m *MsgCreateCommentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCommentResponse) ProtoMessage()    {}
func (*MsgCreateCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90382f40c04acbba, []int{1}
}
func (m *MsgCreateCommentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateCommentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateCommentResponse.Merge(m, src)
}
func (m *MsgCreateCommentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateCommentResponse proto.InternalMessageInfo

func (m *MsgCreateCommentResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgUpdateComment struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Body    string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	PostID  string `protobuf:"bytes,4,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (m *MsgUpdateComment) Reset()         { *m = MsgUpdateComment{} }
func (m *MsgUpdateComment) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateComment) ProtoMessage()    {}
func (*MsgUpdateComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_90382f40c04acbba, []int{2}
}
func (m *MsgUpdateComment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateComment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateComment.Merge(m, src)
}
func (m *MsgUpdateComment) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateComment) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateComment.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateComment proto.InternalMessageInfo

func (m *MsgUpdateComment) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateComment) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgUpdateComment) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MsgUpdateComment) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

type MsgUpdateCommentResponse struct {
}

func (m *MsgUpdateCommentResponse) Reset()         { *m = MsgUpdateCommentResponse{} }
func (m *MsgUpdateCommentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateCommentResponse) ProtoMessage()    {}
func (*MsgUpdateCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90382f40c04acbba, []int{3}
}
func (m *MsgUpdateCommentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateCommentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateCommentResponse.Merge(m, src)
}
func (m *MsgUpdateCommentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateCommentResponse proto.InternalMessageInfo

type MsgDeleteComment struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgDeleteComment) Reset()         { *m = MsgDeleteComment{} }
func (m *MsgDeleteComment) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteComment) ProtoMessage()    {}
func (*MsgDeleteComment) Descriptor() ([]byte, []int) {
	return fileDescriptor_90382f40c04acbba, []int{4}
}
func (m *MsgDeleteComment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteComment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteComment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteComment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteComment.Merge(m, src)
}
func (m *MsgDeleteComment) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteComment) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteComment.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteComment proto.InternalMessageInfo

func (m *MsgDeleteComment) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteComment) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgDeleteCommentResponse struct {
}

func (m *MsgDeleteCommentResponse) Reset()         { *m = MsgDeleteCommentResponse{} }
func (m *MsgDeleteCommentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteCommentResponse) ProtoMessage()    {}
func (*MsgDeleteCommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90382f40c04acbba, []int{5}
}
func (m *MsgDeleteCommentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteCommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteCommentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteCommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteCommentResponse.Merge(m, src)
}
func (m *MsgDeleteCommentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteCommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteCommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteCommentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateComment)(nil), "slandymani.blog.blog.MsgCreateComment")
	proto.RegisterType((*MsgCreateCommentResponse)(nil), "slandymani.blog.blog.MsgCreateCommentResponse")
	proto.RegisterType((*MsgUpdateComment)(nil), "slandymani.blog.blog.MsgUpdateComment")
	proto.RegisterType((*MsgUpdateCommentResponse)(nil), "slandymani.blog.blog.MsgUpdateCommentResponse")
	proto.RegisterType((*MsgDeleteComment)(nil), "slandymani.blog.blog.MsgDeleteComment")
	proto.RegisterType((*MsgDeleteCommentResponse)(nil), "slandymani.blog.blog.MsgDeleteCommentResponse")
}

func init() { proto.RegisterFile("blog/tx.proto", fileDescriptor_90382f40c04acbba) }

var fileDescriptor_90382f40c04acbba = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xca, 0xc9, 0x4f,
	0xd7, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x29, 0xce, 0x49, 0xcc, 0x4b,
	0xa9, 0xcc, 0x4d, 0xcc, 0xcb, 0xd4, 0x03, 0xc9, 0x80, 0x09, 0xa5, 0x08, 0x2e, 0x01, 0xdf, 0xe2,
	0x74, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0xe7, 0xfc, 0xdc, 0xdc, 0xd4, 0xbc, 0x12, 0x21, 0x09,
	0x2e, 0xf6, 0x64, 0x90, 0x40, 0x7e, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b,
	0x24, 0xc4, 0xc5, 0x92, 0x94, 0x9f, 0x52, 0x29, 0xc1, 0x04, 0x16, 0x06, 0xb3, 0x85, 0xc4, 0xb8,
	0xd8, 0x0a, 0xf2, 0x8b, 0x4b, 0x3c, 0x5d, 0x24, 0x98, 0xc1, 0xa2, 0x50, 0x9e, 0x92, 0x16, 0x97,
	0x04, 0xba, 0xc9, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x99,
	0x29, 0x60, 0xc3, 0x59, 0x82, 0x98, 0x32, 0x53, 0x94, 0x32, 0xc0, 0xae, 0x08, 0x2d, 0x48, 0x21,
	0xca, 0x15, 0x10, 0xdd, 0x4c, 0x30, 0xdd, 0x70, 0x57, 0x31, 0x63, 0x75, 0x15, 0x0b, 0x8a, 0xab,
	0xa4, 0xc0, 0xae, 0x42, 0xb1, 0x09, 0xe6, 0x2a, 0x25, 0x1b, 0xb0, 0x2b, 0x5c, 0x52, 0x73, 0x52,
	0xc9, 0x70, 0x05, 0xd4, 0x64, 0x14, 0xdd, 0x30, 0x93, 0x8d, 0x0e, 0x30, 0x71, 0x31, 0xfb, 0x16,
	0xa7, 0x0b, 0xa5, 0x73, 0xf1, 0xa2, 0x06, 0xb5, 0x9a, 0x1e, 0xb6, 0x58, 0xd1, 0x43, 0x0f, 0x38,
	0x29, 0x3d, 0xe2, 0xd4, 0xc1, 0x03, 0x38, 0x9d, 0x8b, 0x17, 0x35, 0x34, 0x71, 0x5b, 0x84, 0xa2,
	0x0e, 0x8f, 0x45, 0x58, 0xc3, 0x0c, 0x64, 0x11, 0x6a, 0x80, 0xe1, 0xb6, 0x08, 0x45, 0x1d, 0x1e,
	0x8b, 0xb0, 0x06, 0xa1, 0x93, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0xa9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x23, 0xcc, 0xd4, 0x07,
	0xa7, 0xfe, 0x0a, 0x08, 0x55, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xce, 0x08, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0xb1, 0xa8, 0xb1, 0x19, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateComment(ctx context.Context, in *MsgCreateComment, opts ...grpc.CallOption) (*MsgCreateCommentResponse, error)
	UpdateComment(ctx context.Context, in *MsgUpdateComment, opts ...grpc.CallOption) (*MsgUpdateCommentResponse, error)
	DeleteComment(ctx context.Context, in *MsgDeleteComment, opts ...grpc.CallOption) (*MsgDeleteCommentResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateComment(ctx context.Context, in *MsgCreateComment, opts ...grpc.CallOption) (*MsgCreateCommentResponse, error) {
	out := new(MsgCreateCommentResponse)
	err := c.cc.Invoke(ctx, "/slandymani.blog.blog.Msg/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateComment(ctx context.Context, in *MsgUpdateComment, opts ...grpc.CallOption) (*MsgUpdateCommentResponse, error) {
	out := new(MsgUpdateCommentResponse)
	err := c.cc.Invoke(ctx, "/slandymani.blog.blog.Msg/UpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteComment(ctx context.Context, in *MsgDeleteComment, opts ...grpc.CallOption) (*MsgDeleteCommentResponse, error) {
	out := new(MsgDeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/slandymani.blog.blog.Msg/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateComment(context.Context, *MsgCreateComment) (*MsgCreateCommentResponse, error)
	UpdateComment(context.Context, *MsgUpdateComment) (*MsgUpdateCommentResponse, error)
	DeleteComment(context.Context, *MsgDeleteComment) (*MsgDeleteCommentResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateComment(ctx context.Context, req *MsgCreateComment) (*MsgCreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (*UnimplementedMsgServer) UpdateComment(ctx context.Context, req *MsgUpdateComment) (*MsgUpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (*UnimplementedMsgServer) DeleteComment(ctx context.Context, req *MsgDeleteComment) (*MsgDeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slandymani.blog.blog.Msg/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateComment(ctx, req.(*MsgCreateComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slandymani.blog.blog.Msg/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateComment(ctx, req.(*MsgUpdateComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slandymani.blog.blog.Msg/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteComment(ctx, req.(*MsgDeleteComment))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slandymani.blog.blog.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _Msg_CreateComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _Msg_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Msg_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog/tx.proto",
}

func (m *MsgCreateComment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateComment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateComment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateCommentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateCommentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateCommentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateComment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateComment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateComment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateCommentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateCommentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateCommentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteComment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteComment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteComment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteCommentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteCommentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteCommentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateComment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateCommentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgUpdateComment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateCommentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteComment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgDeleteCommentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateComment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateComment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateComment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateCommentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateCommentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateCommentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateComment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateComment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateComment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateCommentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateCommentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateCommentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteComment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteComment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteComment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteCommentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteCommentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteCommentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
