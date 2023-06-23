// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/kyc/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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
	return fileDescriptor_f52b05ef9968ac13, []int{0}
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
	return fileDescriptor_f52b05ef9968ac13, []int{1}
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

type QueryGetKYCRegistrationRequest struct {
	KycRegistrationW3CIdentifier string `protobuf:"bytes,1,opt,name=kycRegistrationW3CIdentifier,proto3" json:"kycRegistrationW3CIdentifier,omitempty"`
}

func (m *QueryGetKYCRegistrationRequest) Reset()         { *m = QueryGetKYCRegistrationRequest{} }
func (m *QueryGetKYCRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetKYCRegistrationRequest) ProtoMessage()    {}
func (*QueryGetKYCRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f52b05ef9968ac13, []int{2}
}
func (m *QueryGetKYCRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetKYCRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetKYCRegistrationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetKYCRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetKYCRegistrationRequest.Merge(m, src)
}
func (m *QueryGetKYCRegistrationRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetKYCRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetKYCRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetKYCRegistrationRequest proto.InternalMessageInfo

func (m *QueryGetKYCRegistrationRequest) GetKycRegistrationW3CIdentifier() string {
	if m != nil {
		return m.KycRegistrationW3CIdentifier
	}
	return ""
}

type QueryGetKYCRegistrationResponse struct {
	Registration KYCRegistration `protobuf:"bytes,1,opt,name=registration,proto3" json:"registration"`
}

func (m *QueryGetKYCRegistrationResponse) Reset()         { *m = QueryGetKYCRegistrationResponse{} }
func (m *QueryGetKYCRegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetKYCRegistrationResponse) ProtoMessage()    {}
func (*QueryGetKYCRegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f52b05ef9968ac13, []int{3}
}
func (m *QueryGetKYCRegistrationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetKYCRegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetKYCRegistrationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetKYCRegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetKYCRegistrationResponse.Merge(m, src)
}
func (m *QueryGetKYCRegistrationResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetKYCRegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetKYCRegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetKYCRegistrationResponse proto.InternalMessageInfo

func (m *QueryGetKYCRegistrationResponse) GetRegistration() KYCRegistration {
	if m != nil {
		return m.Registration
	}
	return KYCRegistration{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "doxchain.kyc.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "doxchain.kyc.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryGetKYCRegistrationRequest)(nil), "doxchain.kyc.v1beta1.QueryGetKYCRegistrationRequest")
	proto.RegisterType((*QueryGetKYCRegistrationResponse)(nil), "doxchain.kyc.v1beta1.QueryGetKYCRegistrationResponse")
}

func init() { proto.RegisterFile("doxchain/kyc/v1beta1/query.proto", fileDescriptor_f52b05ef9968ac13) }

var fileDescriptor_f52b05ef9968ac13 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x0b, 0xd3, 0x30,
	0x1c, 0xc5, 0xdb, 0xa1, 0x03, 0xa3, 0x20, 0xc4, 0x1d, 0xa4, 0x8c, 0x6c, 0x16, 0x94, 0x4d, 0xb1,
	0x61, 0x9b, 0x5e, 0x3c, 0x6e, 0x87, 0x31, 0x3c, 0xcc, 0xf5, 0x22, 0x7a, 0x91, 0xb4, 0x8b, 0x5d,
	0xa8, 0x6b, 0xba, 0x24, 0x93, 0x15, 0xf1, 0x22, 0x78, 0x17, 0xfc, 0x52, 0xbb, 0x39, 0xf0, 0xe2,
	0x49, 0x74, 0xf3, 0x83, 0xc8, 0xd2, 0x4c, 0xba, 0xd9, 0x0d, 0xbd, 0x95, 0xe6, 0xe5, 0xfd, 0xde,
	0xfb, 0xff, 0x03, 0x9a, 0x53, 0xbe, 0x0a, 0x67, 0x84, 0x25, 0x38, 0xce, 0x42, 0xfc, 0xb6, 0x13,
	0x50, 0x45, 0x3a, 0x78, 0xb1, 0xa4, 0x22, 0xf3, 0x52, 0xc1, 0x15, 0x87, 0xb5, 0x83, 0xc2, 0x8b,
	0xb3, 0xd0, 0x33, 0x0a, 0xa7, 0x16, 0xf1, 0x88, 0x6b, 0x01, 0xde, 0x7f, 0xe5, 0x5a, 0xa7, 0x1e,
	0x71, 0x1e, 0xbd, 0xa1, 0x98, 0xa4, 0x0c, 0x93, 0x24, 0xe1, 0x8a, 0x28, 0xc6, 0x13, 0x69, 0x4e,
	0xef, 0x87, 0x5c, 0xce, 0xb9, 0xc4, 0x01, 0x91, 0x34, 0x47, 0xfc, 0x01, 0xa6, 0x24, 0x62, 0x89,
	0x16, 0x1b, 0xed, 0x9d, 0xd2, 0x5c, 0x29, 0x11, 0x64, 0x7e, 0xb0, 0x7b, 0x50, 0x2a, 0x89, 0xb3,
	0xf0, 0x95, 0xa0, 0x11, 0x93, 0x4a, 0x14, 0xfc, 0xdc, 0x1a, 0x80, 0x93, 0x3d, 0xf1, 0x99, 0x76,
	0xf0, 0xe9, 0x62, 0x49, 0xa5, 0x72, 0x27, 0xe0, 0xd6, 0xd1, 0x5f, 0x99, 0xf2, 0x44, 0x52, 0xf8,
	0x04, 0x54, 0x73, 0xd2, 0x6d, 0xbb, 0x69, 0xb7, 0xae, 0x77, 0xeb, 0x5e, 0xd9, 0x0c, 0xbc, 0xfc,
	0x56, 0xff, 0xca, 0xfa, 0x7b, 0xc3, 0xf2, 0xcd, 0x0d, 0x77, 0x0a, 0x90, 0xb6, 0x1c, 0x52, 0xf5,
	0xf4, 0xc5, 0xc0, 0x2f, 0x24, 0x31, 0x50, 0xd8, 0x07, 0xf5, 0x38, 0x0b, 0x8b, 0x27, 0xcf, 0x7b,
	0x83, 0xd1, 0x94, 0x26, 0x8a, 0xbd, 0x66, 0x54, 0x68, 0xe6, 0x35, 0xff, 0xa2, 0xc6, 0x15, 0xa0,
	0x71, 0x96, 0x62, 0x4a, 0x8c, 0xc1, 0x8d, 0xe2, 0x1c, 0x4c, 0x95, 0xbb, 0xe5, 0x55, 0x4e, 0x4c,
	0x4c, 0xa7, 0x23, 0x83, 0xee, 0xcf, 0x0a, 0xb8, 0xaa, 0xa1, 0xf0, 0xa3, 0x0d, 0xaa, 0x79, 0x79,
	0xd8, 0x2a, 0xf7, 0xfb, 0x7b, 0xd6, 0x4e, 0xfb, 0x1f, 0x94, 0x79, 0x74, 0xf7, 0xde, 0x87, 0xaf,
	0xbf, 0x3e, 0x57, 0x9a, 0x10, 0xe1, 0x80, 0x3e, 0x9c, 0x51, 0xc1, 0xa9, 0xc4, 0x47, 0xcb, 0xce,
	0x67, 0x0d, 0xbf, 0xd8, 0xe0, 0xe6, 0x49, 0x72, 0xf8, 0xe8, 0x02, 0xe6, 0xec, 0x4e, 0x9c, 0xc7,
	0xff, 0x79, 0xcb, 0x04, 0x1d, 0xeb, 0xa0, 0x23, 0x38, 0x3c, 0x17, 0xf4, 0xf4, 0x35, 0xe2, 0x77,
	0x97, 0xd6, 0xfa, 0xbe, 0x3f, 0x58, 0x6f, 0x91, 0xbd, 0xd9, 0x22, 0xfb, 0xc7, 0x16, 0xd9, 0x9f,
	0x76, 0xc8, 0xda, 0xec, 0x90, 0xf5, 0x6d, 0x87, 0xac, 0x97, 0xed, 0x88, 0xa9, 0xd9, 0x32, 0xf0,
	0x42, 0x3e, 0x2f, 0x83, 0xad, 0x34, 0x4e, 0x65, 0x29, 0x95, 0x41, 0x55, 0x3f, 0xf9, 0xde, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0xca, 0xa4, 0x9e, 0xdc, 0x03, 0x00, 0x00,
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
	KYCRegistration(ctx context.Context, in *QueryGetKYCRegistrationRequest, opts ...grpc.CallOption) (*QueryGetKYCRegistrationResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/doxchain.kyc.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) KYCRegistration(ctx context.Context, in *QueryGetKYCRegistrationRequest, opts ...grpc.CallOption) (*QueryGetKYCRegistrationResponse, error) {
	out := new(QueryGetKYCRegistrationResponse)
	err := c.cc.Invoke(ctx, "/doxchain.kyc.v1beta1.Query/KYCRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	KYCRegistration(context.Context, *QueryGetKYCRegistrationRequest) (*QueryGetKYCRegistrationResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) KYCRegistration(ctx context.Context, req *QueryGetKYCRegistrationRequest) (*QueryGetKYCRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KYCRegistration not implemented")
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
		FullMethod: "/doxchain.kyc.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_KYCRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetKYCRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).KYCRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doxchain.kyc.v1beta1.Query/KYCRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).KYCRegistration(ctx, req.(*QueryGetKYCRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "doxchain.kyc.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "KYCRegistration",
			Handler:    _Query_KYCRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doxchain/kyc/v1beta1/query.proto",
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

func (m *QueryGetKYCRegistrationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetKYCRegistrationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetKYCRegistrationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KycRegistrationW3CIdentifier) > 0 {
		i -= len(m.KycRegistrationW3CIdentifier)
		copy(dAtA[i:], m.KycRegistrationW3CIdentifier)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.KycRegistrationW3CIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetKYCRegistrationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetKYCRegistrationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetKYCRegistrationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Registration.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryGetKYCRegistrationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KycRegistrationW3CIdentifier)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetKYCRegistrationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Registration.Size()
	n += 1 + l + sovQuery(uint64(l))
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
func (m *QueryGetKYCRegistrationRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetKYCRegistrationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetKYCRegistrationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycRegistrationW3CIdentifier", wireType)
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
			m.KycRegistrationW3CIdentifier = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetKYCRegistrationResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetKYCRegistrationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetKYCRegistrationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registration", wireType)
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
			if err := m.Registration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
