// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: users.proto

package users

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Money                int64    `protobuf:"varint,3,opt,name=money,proto3" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetMoney() int64 {
	if m != nil {
		return m.Money
	}
	return 0
}

type UserList struct {
	User                 []*User  `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}
func (m *UserList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return m.Size()
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
}
func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Response struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*UserList)(nil), "UserList")
	proto.RegisterType((*GetUserRequest)(nil), "GetUserRequest")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea) }

var fileDescriptor_030765f334c86cea = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x6a, 0xf3, 0x30,
	0x10, 0x85, 0xad, 0x38, 0xc9, 0xef, 0x4c, 0x20, 0x09, 0xc3, 0x4f, 0x70, 0x5d, 0x30, 0x46, 0x25,
	0x90, 0x95, 0x02, 0x69, 0x0f, 0x50, 0x4a, 0x4a, 0x37, 0xed, 0xc6, 0x90, 0x03, 0x38, 0x64, 0x1a,
	0x02, 0xb1, 0xe5, 0x5a, 0xf2, 0x22, 0x37, 0xe9, 0x15, 0x7a, 0x93, 0x2e, 0x7b, 0x84, 0xe2, 0x5e,
	0xa4, 0x48, 0xb2, 0x17, 0x2d, 0xdd, 0xcd, 0x9b, 0x37, 0xd2, 0xfb, 0x1e, 0x8c, 0x6b, 0x45, 0x95,
	0x12, 0x65, 0x25, 0xb5, 0x8c, 0x2e, 0x0f, 0x52, 0x1e, 0x4e, 0xb4, 0xb2, 0x6a, 0x57, 0x3f, 0xaf,
	0x28, 0x2f, 0xf5, 0xd9, 0x99, 0xfc, 0x16, 0xfa, 0x5b, 0x45, 0x15, 0x4e, 0xa0, 0x77, 0xdc, 0x87,
	0x2c, 0x61, 0x4b, 0x3f, 0xed, 0x1d, 0xf7, 0x88, 0xd0, 0x2f, 0xb2, 0x9c, 0xc2, 0x5e, 0xc2, 0x96,
	0xa3, 0xd4, 0xce, 0xf8, 0x1f, 0x06, 0xb9, 0x2c, 0xe8, 0x1c, 0xfa, 0xf6, 0xcc, 0x09, 0xbe, 0x80,
	0xc0, 0xfc, 0xf0, 0x78, 0x54, 0x1a, 0x2f, 0xa0, 0x6f, 0x92, 0x43, 0x96, 0xf8, 0xcb, 0xf1, 0x7a,
	0x20, 0x8c, 0x91, 0xda, 0x15, 0x4f, 0x60, 0xf2, 0x40, 0xda, 0x2e, 0xe8, 0xa5, 0x26, 0xa5, 0x7f,
	0x47, 0xf2, 0x1b, 0x08, 0x52, 0x52, 0xa5, 0x2c, 0x14, 0xe1, 0x1c, 0x86, 0x4a, 0x67, 0xba, 0x56,
	0xd6, 0x0f, 0xd2, 0x56, 0xe1, 0x0c, 0xfc, 0x5c, 0x1d, 0x5a, 0x2a, 0x33, 0xae, 0xdf, 0x98, 0xcb,
	0xdf, 0x64, 0x3a, 0xc3, 0x15, 0x04, 0x6d, 0x88, 0xc2, 0xb9, 0x70, 0xbd, 0x45, 0xd7, 0x5b, 0xdc,
	0x9b, 0xde, 0xd1, 0x48, 0x74, 0xb8, 0xdc, 0xc3, 0x2b, 0xf8, 0xd7, 0x3e, 0xc0, 0xa9, 0xf8, 0xc9,
	0x17, 0x39, 0x7c, 0xee, 0xe1, 0x02, 0xa6, 0xdb, 0x72, 0x9f, 0x69, 0x32, 0xfa, 0xc9, 0x94, 0x46,
	0xe7, 0x45, 0x23, 0xd1, 0x11, 0x73, 0x0f, 0x13, 0x80, 0x0d, 0x9d, 0xc8, 0x9d, 0xfd, 0x75, 0x71,
	0x37, 0x7b, 0x6f, 0x62, 0xf6, 0xd1, 0xc4, 0xec, 0xb3, 0x89, 0xd9, 0xeb, 0x57, 0xec, 0xed, 0x86,
	0x16, 0xee, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x22, 0x09, 0xd2, 0x6e, 0xb1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserDataClient is the client API for UserData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserDataClient interface {
	GetUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserList, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUserMoney(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
}

type userDataClient struct {
	cc *grpc.ClientConn
}

func NewUserDataClient(cc *grpc.ClientConn) UserDataClient {
	return &userDataClient{cc}
}

func (c *userDataClient) GetUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/UserData/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserData/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataClient) UpdateUserMoney(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/UserData/UpdateUserMoney", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/UserData/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDataServer is the server API for UserData service.
type UserDataServer interface {
	GetUsers(context.Context, *emptypb.Empty) (*UserList, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	UpdateUserMoney(context.Context, *User) (*Response, error)
	DeleteUser(context.Context, *User) (*Response, error)
}

// UnimplementedUserDataServer can be embedded to have forward compatible implementations.
type UnimplementedUserDataServer struct {
}

func (*UnimplementedUserDataServer) GetUsers(ctx context.Context, req *emptypb.Empty) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedUserDataServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserDataServer) UpdateUserMoney(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMoney not implemented")
}
func (*UnimplementedUserDataServer) DeleteUser(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterUserDataServer(s *grpc.Server, srv UserDataServer) {
	s.RegisterService(&_UserData_serviceDesc, srv)
}

func _UserData_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserData/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServer).GetUsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserData_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserData/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserData_UpdateUserMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServer).UpdateUserMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserData/UpdateUserMoney",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServer).UpdateUserMoney(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserData_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserData/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserData",
	HandlerType: (*UserDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _UserData_GetUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserData_GetUser_Handler,
		},
		{
			MethodName: "UpdateUserMoney",
			Handler:    _UserData_UpdateUserMoney_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserData_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Money != 0 {
		i = encodeVarintUsers(dAtA, i, uint64(m.Money))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUsers(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintUsers(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.User) > 0 {
		for iNdEx := len(m.User) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.User[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUsers(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetUserRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetUserRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetUserRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintUsers(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintUsers(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status {
		i--
		if m.Status {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUsers(dAtA []byte, offset int, v uint64) int {
	offset -= sovUsers(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUsers(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUsers(uint64(l))
	}
	if m.Money != 0 {
		n += 1 + sovUsers(uint64(m.Money))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UserList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.User) > 0 {
		for _, e := range m.User {
			l = e.Size()
			n += 1 + l + sovUsers(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetUserRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUsers(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status {
		n += 2
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovUsers(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUsers(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUsers(x uint64) (n int) {
	return sovUsers(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUsers
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
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
				return ErrInvalidLengthUsers
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUsers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Money", wireType)
			}
			m.Money = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Money |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUsers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUsers
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUsers
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
			return fmt.Errorf("proto: UserList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUsers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUsers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = append(m.User, &User{})
			if err := m.User[len(m.User)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUsers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUsers
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetUserRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUsers
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
			return fmt.Errorf("proto: GetUserRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetUserRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUsers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUsers
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUsers
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Status = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUsers
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
				return ErrInvalidLengthUsers
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUsers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUsers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUsers
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUsers(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUsers
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
					return 0, ErrIntOverflowUsers
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
					return 0, ErrIntOverflowUsers
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
				return 0, ErrInvalidLengthUsers
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUsers
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUsers
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUsers        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUsers          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUsers = fmt.Errorf("proto: unexpected end of group")
)