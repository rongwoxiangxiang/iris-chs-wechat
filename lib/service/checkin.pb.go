// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkin.proto

package service

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

type Checkin struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Wid                  int64    `protobuf:"varint,2,opt,name=Wid,proto3" json:"Wid,omitempty"`
	ActivityId           int64    `protobuf:"varint,3,opt,name=ActivityId,proto3" json:"ActivityId,omitempty"`
	Wuid                 int64    `protobuf:"varint,4,opt,name=Wuid,proto3" json:"Wuid,omitempty"`
	Liner                int64    `protobuf:"varint,5,opt,name=Liner,proto3" json:"Liner,omitempty"`
	Total                int64    `protobuf:"varint,6,opt,name=Total,proto3" json:"Total,omitempty"`
	Lastcheckin          int64    `protobuf:"varint,7,opt,name=Lastcheckin,proto3" json:"Lastcheckin,omitempty"`
	CreatedAt            int64    `protobuf:"varint,10,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,11,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkin) Reset()         { *m = Checkin{} }
func (m *Checkin) String() string { return proto.CompactTextString(m) }
func (*Checkin) ProtoMessage()    {}
func (*Checkin) Descriptor() ([]byte, []int) {
	return fileDescriptor_072e71e6019dc001, []int{0}
}

func (m *Checkin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkin.Unmarshal(m, b)
}
func (m *Checkin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkin.Marshal(b, m, deterministic)
}
func (m *Checkin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkin.Merge(m, src)
}
func (m *Checkin) XXX_Size() int {
	return xxx_messageInfo_Checkin.Size(m)
}
func (m *Checkin) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkin.DiscardUnknown(m)
}

var xxx_messageInfo_Checkin proto.InternalMessageInfo

func (m *Checkin) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Checkin) GetWid() int64 {
	if m != nil {
		return m.Wid
	}
	return 0
}

func (m *Checkin) GetActivityId() int64 {
	if m != nil {
		return m.ActivityId
	}
	return 0
}

func (m *Checkin) GetWuid() int64 {
	if m != nil {
		return m.Wuid
	}
	return 0
}

func (m *Checkin) GetLiner() int64 {
	if m != nil {
		return m.Liner
	}
	return 0
}

func (m *Checkin) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Checkin) GetLastcheckin() int64 {
	if m != nil {
		return m.Lastcheckin
	}
	return 0
}

func (m *Checkin) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Checkin) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type ActivityResponse struct {
	Boolean              bool     `protobuf:"varint,1,opt,name=boolean,proto3" json:"boolean,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivityResponse) Reset()         { *m = ActivityResponse{} }
func (m *ActivityResponse) String() string { return proto.CompactTextString(m) }
func (*ActivityResponse) ProtoMessage()    {}
func (*ActivityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_072e71e6019dc001, []int{1}
}

func (m *ActivityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityResponse.Unmarshal(m, b)
}
func (m *ActivityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityResponse.Marshal(b, m, deterministic)
}
func (m *ActivityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityResponse.Merge(m, src)
}
func (m *ActivityResponse) XXX_Size() int {
	return xxx_messageInfo_ActivityResponse.Size(m)
}
func (m *ActivityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityResponse proto.InternalMessageInfo

func (m *ActivityResponse) GetBoolean() bool {
	if m != nil {
		return m.Boolean
	}
	return false
}

func (m *ActivityResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ActivityQuery struct {
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Wid                  int64    `protobuf:"varint,3,opt,name=wid,proto3" json:"wid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivityQuery) Reset()         { *m = ActivityQuery{} }
func (m *ActivityQuery) String() string { return proto.CompactTextString(m) }
func (*ActivityQuery) ProtoMessage()    {}
func (*ActivityQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_072e71e6019dc001, []int{2}
}

func (m *ActivityQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityQuery.Unmarshal(m, b)
}
func (m *ActivityQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityQuery.Marshal(b, m, deterministic)
}
func (m *ActivityQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityQuery.Merge(m, src)
}
func (m *ActivityQuery) XXX_Size() int {
	return xxx_messageInfo_ActivityQuery.Size(m)
}
func (m *ActivityQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityQuery proto.InternalMessageInfo

func (m *ActivityQuery) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ActivityQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ActivityQuery) GetWid() int64 {
	if m != nil {
		return m.Wid
	}
	return 0
}

type Checkinlist struct {
	Page                 int64      `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int64      `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Checkins             []*Checkin `protobuf:"bytes,3,rep,name=Checkins,proto3" json:"Checkins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Checkinlist) Reset()         { *m = Checkinlist{} }
func (m *Checkinlist) String() string { return proto.CompactTextString(m) }
func (*Checkinlist) ProtoMessage()    {}
func (*Checkinlist) Descriptor() ([]byte, []int) {
	return fileDescriptor_072e71e6019dc001, []int{3}
}

func (m *Checkinlist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkinlist.Unmarshal(m, b)
}
func (m *Checkinlist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkinlist.Marshal(b, m, deterministic)
}
func (m *Checkinlist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkinlist.Merge(m, src)
}
func (m *Checkinlist) XXX_Size() int {
	return xxx_messageInfo_Checkinlist.Size(m)
}
func (m *Checkinlist) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkinlist.DiscardUnknown(m)
}

var xxx_messageInfo_Checkinlist proto.InternalMessageInfo

func (m *Checkinlist) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Checkinlist) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Checkinlist) GetCheckins() []*Checkin {
	if m != nil {
		return m.Checkins
	}
	return nil
}

func init() {
	proto.RegisterType((*Checkin)(nil), "service.Checkin")
	proto.RegisterType((*ActivityResponse)(nil), "service.ActivityResponse")
	proto.RegisterType((*ActivityQuery)(nil), "service.ActivityQuery")
	proto.RegisterType((*Checkinlist)(nil), "service.Checkinlist")
}

func init() { proto.RegisterFile("checkin.proto", fileDescriptor_072e71e6019dc001) }

var fileDescriptor_072e71e6019dc001 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x9b, 0xa4, 0x6d, 0xea, 0x84, 0x96, 0xb2, 0x14, 0x59, 0x8a, 0x48, 0xc9, 0xa9, 0x07,
	0x89, 0x50, 0xc1, 0x93, 0x97, 0xda, 0x83, 0x04, 0xea, 0xc1, 0xa8, 0xf4, 0x9c, 0x66, 0x07, 0x5d,
	0x8c, 0x49, 0xc8, 0x6e, 0xab, 0xf5, 0x8d, 0x7c, 0x34, 0xdf, 0x42, 0x76, 0xb3, 0x89, 0xd5, 0x5e,
	0xbc, 0xcd, 0xff, 0xfd, 0x3b, 0x9b, 0x99, 0x7f, 0x03, 0xfd, 0xe4, 0x19, 0x93, 0x17, 0x9e, 0x05,
	0x45, 0x99, 0xcb, 0x9c, 0xb8, 0x02, 0xcb, 0x2d, 0x4f, 0xd0, 0xff, 0xb2, 0xc0, 0x5d, 0x54, 0x16,
	0x19, 0x80, 0x1d, 0x32, 0x6a, 0x4d, 0xac, 0xa9, 0x13, 0xd9, 0x21, 0x23, 0x43, 0x70, 0x56, 0x9c,
	0x51, 0x5b, 0x03, 0x55, 0x92, 0x53, 0x80, 0x79, 0x22, 0xf9, 0x96, 0xcb, 0x5d, 0xc8, 0xa8, 0xa3,
	0x8d, 0x3d, 0x42, 0x08, 0xb4, 0x57, 0x1b, 0xce, 0x68, 0x5b, 0x3b, 0xba, 0x26, 0x23, 0xe8, 0x2c,
	0x79, 0x86, 0x25, 0xed, 0x68, 0x58, 0x09, 0x45, 0x1f, 0x72, 0x19, 0xa7, 0xb4, 0x5b, 0x51, 0x2d,
	0xc8, 0x04, 0xbc, 0x65, 0x2c, 0xa4, 0x99, 0x95, 0xba, 0xda, 0xdb, 0x47, 0xe4, 0x04, 0x8e, 0x16,
	0x25, 0xc6, 0x12, 0xd9, 0x5c, 0x52, 0xd0, 0xfe, 0x0f, 0x50, 0xee, 0x63, 0xc1, 0x8c, 0xeb, 0x55,
	0x6e, 0x03, 0xfc, 0x2b, 0x18, 0xd6, 0xb3, 0x46, 0x28, 0x8a, 0x3c, 0x13, 0x48, 0x28, 0xb8, 0xeb,
	0x3c, 0x4f, 0x31, 0xce, 0xf4, 0xe2, 0xbd, 0xa8, 0x96, 0x2a, 0x8d, 0x66, 0x79, 0x9b, 0x33, 0xff,
	0x16, 0xfa, 0x75, 0xf7, 0xdd, 0x06, 0xcb, 0x9d, 0x5a, 0x81, 0x67, 0x0c, 0xdf, 0x4d, 0x62, 0x95,
	0x50, 0x34, 0xe5, 0xaf, 0x5c, 0x9a, 0xce, 0x4a, 0xa8, 0x28, 0xdf, 0x78, 0x9d, 0x98, 0x2a, 0xfd,
	0x04, 0x3c, 0x93, 0x7b, 0xca, 0x85, 0x54, 0xc9, 0x15, 0xf1, 0x13, 0x9a, 0xbb, 0x74, 0xad, 0x98,
	0xe0, 0x1f, 0x68, 0x6e, 0xd2, 0x35, 0x39, 0x83, 0x9e, 0x69, 0x13, 0xd4, 0x99, 0x38, 0x53, 0x6f,
	0x36, 0x0c, 0xcc, 0x5b, 0x06, 0xc6, 0x88, 0x9a, 0x13, 0xb3, 0x4f, 0x0b, 0x06, 0x46, 0xdc, 0x57,
	0x87, 0xc8, 0x39, 0xb8, 0x37, 0x28, 0xaf, 0xd5, 0x6b, 0x1d, 0x74, 0x8e, 0x0f, 0x88, 0xdf, 0x22,
	0x97, 0xd0, 0x5e, 0xaa, 0x09, 0x8f, 0x1b, 0xef, 0x57, 0x0c, 0xe3, 0xd1, 0xdf, 0x1e, 0xb5, 0x8f,
	0xdf, 0x22, 0x01, 0x74, 0xc3, 0x4c, 0x60, 0x29, 0xff, 0xf7, 0x9d, 0x75, 0x57, 0xff, 0x99, 0x17,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x89, 0xd1, 0x83, 0xaa, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CheckinServiceClient is the client API for CheckinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckinServiceClient interface {
	GetById(ctx context.Context, in *Checkin, opts ...grpc.CallOption) (*Checkin, error)
	List(ctx context.Context, in *ActivityQuery, opts ...grpc.CallOption) (*Checkinlist, error)
	Insert(ctx context.Context, in *Checkin, opts ...grpc.CallOption) (*Checkin, error)
}

type checkinServiceClient struct {
	cc *grpc.ClientConn
}

func NewCheckinServiceClient(cc *grpc.ClientConn) CheckinServiceClient {
	return &checkinServiceClient{cc}
}

func (c *checkinServiceClient) GetById(ctx context.Context, in *Checkin, opts ...grpc.CallOption) (*Checkin, error) {
	out := new(Checkin)
	err := c.cc.Invoke(ctx, "/service.CheckinService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkinServiceClient) List(ctx context.Context, in *ActivityQuery, opts ...grpc.CallOption) (*Checkinlist, error) {
	out := new(Checkinlist)
	err := c.cc.Invoke(ctx, "/service.CheckinService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkinServiceClient) Insert(ctx context.Context, in *Checkin, opts ...grpc.CallOption) (*Checkin, error) {
	out := new(Checkin)
	err := c.cc.Invoke(ctx, "/service.CheckinService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckinServiceServer is the server API for CheckinService service.
type CheckinServiceServer interface {
	GetById(context.Context, *Checkin) (*Checkin, error)
	List(context.Context, *ActivityQuery) (*Checkinlist, error)
	Insert(context.Context, *Checkin) (*Checkin, error)
}

// UnimplementedCheckinServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCheckinServiceServer struct {
}

func (*UnimplementedCheckinServiceServer) GetById(ctx context.Context, req *Checkin) (*Checkin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (*UnimplementedCheckinServiceServer) List(ctx context.Context, req *ActivityQuery) (*Checkinlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCheckinServiceServer) Insert(ctx context.Context, req *Checkin) (*Checkin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}

func RegisterCheckinServiceServer(s *grpc.Server, srv CheckinServiceServer) {
	s.RegisterService(&_CheckinService_serviceDesc, srv)
}

func _CheckinService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Checkin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckinServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CheckinService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckinServiceServer).GetById(ctx, req.(*Checkin))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckinService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckinServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CheckinService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckinServiceServer).List(ctx, req.(*ActivityQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckinService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Checkin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckinServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CheckinService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckinServiceServer).Insert(ctx, req.(*Checkin))
	}
	return interceptor(ctx, in, info, handler)
}

var _CheckinService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.CheckinService",
	HandlerType: (*CheckinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _CheckinService_GetById_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CheckinService_List_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _CheckinService_Insert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkin.proto",
}
