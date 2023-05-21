// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/did/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8152301e8091f4d6, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8152301e8091f4d6, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetDidRequest struct {
	DidW3CIdentifier string `protobuf:"bytes,1,opt,name=didW3CIdentifier,proto3" json:"didW3CIdentifier,omitempty"`
}

func (m *QueryGetDidRequest) Reset()         { *m = QueryGetDidRequest{} }
func (m *QueryGetDidRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetDidRequest) ProtoMessage()    {}
func (*QueryGetDidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8152301e8091f4d6, []int{2}
}
func (m *QueryGetDidRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDidRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDidRequest.Merge(m, src)
}
func (m *QueryGetDidRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDidRequest proto.InternalMessageInfo

func (m *QueryGetDidRequest) GetDidW3CIdentifier() string {
	if m != nil {
		return m.DidW3CIdentifier
	}
	return ""
}

type QueryGetDidResponse struct {
	Did Did `protobuf:"bytes,1,opt,name=Did,proto3" json:"Did"`
}

func (m *QueryGetDidResponse) Reset()         { *m = QueryGetDidResponse{} }
func (m *QueryGetDidResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetDidResponse) ProtoMessage()    {}
func (*QueryGetDidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8152301e8091f4d6, []int{3}
}
func (m *QueryGetDidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDidResponse.Merge(m, src)
}
func (m *QueryGetDidResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDidResponse proto.InternalMessageInfo

func (m *QueryGetDidResponse) GetDid() Did {
	if m != nil {
		return m.Did
	}
	return Did{}
}

type QueryAllDidRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllDidRequest) Reset()         { *m = QueryAllDidRequest{} }
func (m *QueryAllDidRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllDidRequest) ProtoMessage()    {}
func (*QueryAllDidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8152301e8091f4d6, []int{4}
}
func (m *QueryAllDidRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllDidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDidRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllDidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDidRequest.Merge(m, src)
}
func (m *QueryAllDidRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllDidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllDidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllDidRequest proto.InternalMessageInfo

func (m *QueryAllDidRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllDidResponse struct {
	DidList    []Did               `protobuf:"bytes,1,rep,name=didList,proto3" json:"didList"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllDidResponse) Reset()         { *m = QueryAllDidResponse{} }
func (m *QueryAllDidResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllDidResponse) ProtoMessage()    {}
func (*QueryAllDidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8152301e8091f4d6, []int{5}
}
func (m *QueryAllDidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllDidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllDidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDidResponse.Merge(m, src)
}
func (m *QueryAllDidResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllDidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllDidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllDidResponse proto.InternalMessageInfo

func (m *QueryAllDidResponse) GetDidList() []Did {
	if m != nil {
		return m.DidList
	}
	return nil
}

func (m *QueryAllDidResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "beheroes.doxchain.did.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "beheroes.doxchain.did.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryGetDidRequest)(nil), "beheroes.doxchain.did.v1beta1.QueryGetDidRequest")
	proto.RegisterType((*QueryGetDidResponse)(nil), "beheroes.doxchain.did.v1beta1.QueryGetDidResponse")
	proto.RegisterType((*QueryAllDidRequest)(nil), "beheroes.doxchain.did.v1beta1.QueryAllDidRequest")
	proto.RegisterType((*QueryAllDidResponse)(nil), "beheroes.doxchain.did.v1beta1.QueryAllDidResponse")
}

func init() { proto.RegisterFile("doxchain/did/v1beta1/query.proto", fileDescriptor_8152301e8091f4d6) }

var fileDescriptor_8152301e8091f4d6 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x8b, 0x13, 0x31,
	0x1c, 0xc5, 0x9b, 0xad, 0x56, 0x8c, 0x17, 0xcd, 0x16, 0x59, 0xcb, 0x3a, 0xd6, 0xa0, 0xab, 0x2e,
	0x9a, 0xd0, 0xee, 0xcd, 0x93, 0xdb, 0x16, 0x17, 0xc1, 0x83, 0xdb, 0x8b, 0xb0, 0x78, 0xc9, 0x34,
	0x71, 0x1a, 0x98, 0x4e, 0x66, 0x27, 0xa9, 0xec, 0x22, 0x5e, 0x3c, 0x0a, 0x82, 0xe0, 0xcd, 0xa3,
	0xdf, 0xc4, 0xdb, 0x1e, 0x17, 0xbc, 0x78, 0x12, 0x69, 0xfd, 0x20, 0x32, 0x49, 0xc6, 0x76, 0xec,
	0x62, 0xbb, 0x87, 0x42, 0x49, 0xfe, 0xef, 0xbd, 0x5f, 0xfa, 0x7f, 0x14, 0x36, 0xb9, 0x3a, 0x1a,
	0x0c, 0x99, 0x4c, 0x28, 0x97, 0x9c, 0xbe, 0x69, 0x85, 0xc2, 0xb0, 0x16, 0x3d, 0x1c, 0x8b, 0xec,
	0x98, 0xa4, 0x99, 0x32, 0x0a, 0xdd, 0x0c, 0xc5, 0x50, 0x64, 0x4a, 0x68, 0x52, 0x8c, 0x12, 0x2e,
	0x39, 0xf1, 0xa3, 0x8d, 0x7a, 0xa4, 0x22, 0x65, 0x27, 0x69, 0xfe, 0xcd, 0x89, 0x1a, 0x9b, 0x91,
	0x52, 0x51, 0x2c, 0x28, 0x4b, 0x25, 0x65, 0x49, 0xa2, 0x0c, 0x33, 0x52, 0x25, 0xda, 0xdf, 0x6e,
	0x0f, 0x94, 0x1e, 0x29, 0x4d, 0x43, 0xa6, 0x85, 0xcb, 0xfa, 0x9b, 0x9c, 0xb2, 0x48, 0x26, 0x76,
	0xd8, 0xcf, 0xde, 0x3e, 0x13, 0x30, 0x65, 0x19, 0x1b, 0x15, 0x76, 0xc1, 0x99, 0x23, 0x39, 0xa4,
	0xbd, 0xc7, 0x75, 0x88, 0xf6, 0xf3, 0x90, 0x17, 0x56, 0xd4, 0x17, 0x87, 0x63, 0xa1, 0x0d, 0x3e,
	0x80, 0xeb, 0xa5, 0x53, 0x9d, 0xaa, 0x44, 0x0b, 0xd4, 0x85, 0x35, 0x67, 0xbe, 0x01, 0x9a, 0xe0,
	0xfe, 0x95, 0xf6, 0x5d, 0xf2, 0xdf, 0xf7, 0x13, 0x27, 0xef, 0x5c, 0x38, 0xf9, 0x79, 0xab, 0xd2,
	0xf7, 0x52, 0xfc, 0xc4, 0x27, 0xee, 0x09, 0xd3, 0x93, 0xdc, 0x27, 0xa2, 0x6d, 0x78, 0x95, 0x4b,
	0xfe, 0x72, 0xa7, 0xfb, 0x8c, 0x8b, 0xc4, 0xc8, 0xd7, 0x52, 0x64, 0x36, 0xe4, 0x72, 0x7f, 0xe1,
	0x1c, 0xef, 0x7b, 0xba, 0xc2, 0xc1, 0xd3, 0x3d, 0x86, 0xd5, 0x9e, 0xe4, 0x1e, 0x0d, 0x2f, 0x41,
	0xeb, 0x49, 0xee, 0xb9, 0x72, 0x11, 0x7e, 0xe5, 0xa1, 0x76, 0xe3, 0x78, 0x0e, 0xea, 0x29, 0x84,
	0xb3, 0xdf, 0xdc, 0x1b, 0x6f, 0x11, 0xb7, 0x20, 0x92, 0x2f, 0x88, 0xb8, 0x32, 0xcc, 0xde, 0x1b,
	0x09, 0xaf, 0xed, 0xcf, 0x29, 0xf1, 0x57, 0xe0, 0x89, 0x0b, 0x7b, 0x4f, 0xdc, 0x81, 0x97, 0xb8,
	0xe4, 0xcf, 0xa5, 0x36, 0x1b, 0xa0, 0x59, 0x3d, 0x17, 0x75, 0x21, 0x44, 0x7b, 0x25, 0xc6, 0x35,
	0xcb, 0x78, 0x6f, 0x29, 0xa3, 0x03, 0x98, 0x87, 0x6c, 0x7f, 0xab, 0xc2, 0x8b, 0x16, 0x12, 0x7d,
	0x04, 0xb0, 0xe6, 0x56, 0x87, 0x5a, 0x4b, 0x80, 0x16, 0xbb, 0xd3, 0x68, 0x9f, 0x47, 0xe2, 0x38,
	0xf0, 0xe6, 0xfb, 0xef, 0xbf, 0x3f, 0xaf, 0x5d, 0x47, 0x75, 0x5a, 0xaa, 0xab, 0x6b, 0x0c, 0xfa,
	0x02, 0xec, 0x66, 0x57, 0x83, 0x29, 0xd5, 0x6a, 0x35, 0x98, 0x72, 0x8f, 0xf0, 0x43, 0x0b, 0xb3,
	0x85, 0xee, 0x94, 0x61, 0xf2, 0xcf, 0xdb, 0x7f, 0xbb, 0xf8, 0x0e, 0x7d, 0x00, 0xb0, 0xd6, 0x93,
	0x7c, 0x37, 0x8e, 0x57, 0xe3, 0x2b, 0x35, 0x6c, 0x35, 0xbe, 0x72, 0x6b, 0xf0, 0x0d, 0xcb, 0xb7,
	0x8e, 0xae, 0x2d, 0xf0, 0x75, 0xba, 0x27, 0x93, 0x00, 0x9c, 0x4e, 0x02, 0xf0, 0x6b, 0x12, 0x80,
	0x4f, 0xd3, 0xa0, 0x72, 0x3a, 0x0d, 0x2a, 0x3f, 0xa6, 0x41, 0xe5, 0xe0, 0x41, 0x24, 0xcd, 0x70,
	0x1c, 0x92, 0x81, 0x1a, 0xd1, 0x50, 0x3c, 0x72, 0x99, 0x33, 0x83, 0x23, 0x6b, 0x61, 0x8e, 0x53,
	0xa1, 0xc3, 0x9a, 0xfd, 0x67, 0xd8, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x14, 0xaa, 0x64, 0x69,
	0xff, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	Did(ctx context.Context, in *QueryGetDidRequest, opts ...grpc.CallOption) (*QueryGetDidResponse, error)
	DidAll(ctx context.Context, in *QueryAllDidRequest, opts ...grpc.CallOption) (*QueryAllDidResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/beheroes.doxchain.did.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Did(ctx context.Context, in *QueryGetDidRequest, opts ...grpc.CallOption) (*QueryGetDidResponse, error) {
	out := new(QueryGetDidResponse)
	err := c.cc.Invoke(ctx, "/beheroes.doxchain.did.v1beta1.Query/Did", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DidAll(ctx context.Context, in *QueryAllDidRequest, opts ...grpc.CallOption) (*QueryAllDidResponse, error) {
	out := new(QueryAllDidResponse)
	err := c.cc.Invoke(ctx, "/beheroes.doxchain.did.v1beta1.Query/DidAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	Did(context.Context, *QueryGetDidRequest) (*QueryGetDidResponse, error)
	DidAll(context.Context, *QueryAllDidRequest) (*QueryAllDidResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Did(ctx context.Context, req *QueryGetDidRequest) (*QueryGetDidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Did not implemented")
}
func (*UnimplementedQueryServer) DidAll(ctx context.Context, req *QueryAllDidRequest) (*QueryAllDidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beheroes.doxchain.did.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Did_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetDidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Did(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beheroes.doxchain.did.v1beta1.Query/Did",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Did(ctx, req.(*QueryGetDidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DidAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllDidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DidAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beheroes.doxchain.did.v1beta1.Query/DidAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DidAll(ctx, req.(*QueryAllDidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "beheroes.doxchain.did.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Did",
			Handler:    _Query_Did_Handler,
		},
		{
			MethodName: "DidAll",
			Handler:    _Query_DidAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doxchain/did/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetDidRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDidRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDidRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DidW3CIdentifier) > 0 {
		i -= len(m.DidW3CIdentifier)
		copy(dAtA[i:], m.DidW3CIdentifier)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.DidW3CIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetDidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Did.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllDidRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDidRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDidRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllDidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DidList) > 0 {
		for iNdEx := len(m.DidList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DidList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetDidRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DidW3CIdentifier)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetDidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Did.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllDidRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllDidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DidList) > 0 {
		for _, e := range m.DidList {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetDidRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetDidRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDidRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidW3CIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidW3CIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetDidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetDidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Did.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllDidRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllDidRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDidRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllDidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllDidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DidList = append(m.DidList, Did{})
			if err := m.DidList[len(m.DidList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
