// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/oauthtwo/access_token_registry.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type AccessTokenRegistry struct {
	Tenant string            `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Issued []AccessTokenInfo `protobuf:"bytes,2,rep,name=issued,proto3" json:"issued"`
}

func (m *AccessTokenRegistry) Reset()         { *m = AccessTokenRegistry{} }
func (m *AccessTokenRegistry) String() string { return proto.CompactTextString(m) }
func (*AccessTokenRegistry) ProtoMessage()    {}
func (*AccessTokenRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b56711af2bf99058, []int{0}
}
func (m *AccessTokenRegistry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessTokenRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessTokenRegistry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessTokenRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTokenRegistry.Merge(m, src)
}
func (m *AccessTokenRegistry) XXX_Size() int {
	return m.Size()
}
func (m *AccessTokenRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTokenRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTokenRegistry proto.InternalMessageInfo

func (m *AccessTokenRegistry) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *AccessTokenRegistry) GetIssued() []AccessTokenInfo {
	if m != nil {
		return m.Issued
	}
	return nil
}

type AccessTokenInfo struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	ExpiresAt  int64  `protobuf:"varint,3,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

func (m *AccessTokenInfo) Reset()         { *m = AccessTokenInfo{} }
func (m *AccessTokenInfo) String() string { return proto.CompactTextString(m) }
func (*AccessTokenInfo) ProtoMessage()    {}
func (*AccessTokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b56711af2bf99058, []int{1}
}
func (m *AccessTokenInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessTokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessTokenInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessTokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTokenInfo.Merge(m, src)
}
func (m *AccessTokenInfo) XXX_Size() int {
	return m.Size()
}
func (m *AccessTokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTokenInfo proto.InternalMessageInfo

func (m *AccessTokenInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *AccessTokenInfo) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *AccessTokenInfo) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*AccessTokenRegistry)(nil), "beheroes.doxchain.oauthtwo.AccessTokenRegistry")
	proto.RegisterType((*AccessTokenInfo)(nil), "beheroes.doxchain.oauthtwo.AccessTokenInfo")
}

func init() {
	proto.RegisterFile("doxchain/oauthtwo/access_token_registry.proto", fileDescriptor_b56711af2bf99058)
}

var fileDescriptor_b56711af2bf99058 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0xad, 0x54, 0xba, 0x1e, 0x84, 0x55, 0x24, 0x14, 0x59, 0x4b, 0x4f, 0x05, 0xe9,
	0x06, 0xf4, 0x09, 0xda, 0x93, 0xbd, 0x06, 0x4f, 0x5e, 0x4a, 0x92, 0x4e, 0x93, 0x45, 0xdc, 0x29,
	0xbb, 0x13, 0x4c, 0xdf, 0xc2, 0xc7, 0xea, 0xb1, 0x47, 0x4f, 0x22, 0xc9, 0x8b, 0x48, 0xd3, 0xc4,
	0x16, 0xc1, 0xdb, 0xce, 0xec, 0xff, 0x7f, 0x3f, 0xf3, 0xf3, 0xc9, 0x12, 0x8b, 0x24, 0x8b, 0xb4,
	0x09, 0x30, 0xca, 0x29, 0xa3, 0x77, 0x0c, 0xa2, 0x24, 0x01, 0xe7, 0x16, 0x84, 0xaf, 0x60, 0x16,
	0x16, 0x52, 0xed, 0xc8, 0x6e, 0xd4, 0xda, 0x22, 0xa1, 0x18, 0xc4, 0x90, 0x81, 0x45, 0x70, 0xaa,
	0xf5, 0xa9, 0xd6, 0x37, 0xb8, 0x4e, 0x31, 0xc5, 0x5a, 0x16, 0xec, 0x5f, 0x07, 0xc7, 0xa8, 0xe0,
	0x57, 0xd3, 0x1a, 0xf8, 0xbc, 0xe7, 0x85, 0x0d, 0x4e, 0xdc, 0xf0, 0x1e, 0x81, 0x89, 0x0c, 0xf9,
	0x6c, 0xc8, 0xc6, 0xfd, 0xb0, 0x99, 0xc4, 0x9c, 0xf7, 0xb4, 0x73, 0x39, 0x2c, 0xfd, 0xce, 0xb0,
	0x3b, 0xbe, 0x78, 0xb8, 0x57, 0xff, 0x27, 0xaa, 0x13, 0xf0, 0xdc, 0xac, 0x70, 0x76, 0xb6, 0xfd,
	0xba, 0xf3, 0xc2, 0x06, 0x30, 0xd2, 0xfc, 0xf2, 0x8f, 0x40, 0xf8, 0xfc, 0x3c, 0xb1, 0x10, 0x11,
	0xda, 0x26, 0xb6, 0x1d, 0x85, 0xe4, 0x5c, 0x2f, 0xc1, 0x90, 0x5e, 0x69, 0xb0, 0x7e, 0xa7, 0xfe,
	0x3c, 0xd9, 0x88, 0x5b, 0xde, 0x87, 0x62, 0xad, 0x2d, 0xb8, 0x29, 0xf9, 0xdd, 0x21, 0x1b, 0x77,
	0xc3, 0xe3, 0x62, 0xf6, 0xb4, 0x2d, 0x25, 0xdb, 0x95, 0x92, 0x7d, 0x97, 0x92, 0x7d, 0x54, 0xd2,
	0xdb, 0x55, 0xd2, 0xfb, 0xac, 0xa4, 0xf7, 0xa2, 0x52, 0x4d, 0x59, 0x1e, 0xab, 0x04, 0xdf, 0x82,
	0x18, 0x26, 0x87, 0x53, 0x82, 0xdf, 0xd2, 0x8b, 0x63, 0xed, 0xb4, 0x59, 0x83, 0x8b, 0x7b, 0x75,
	0x6b, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x5a, 0x59, 0xc6, 0x98, 0x01, 0x00, 0x00,
}

func (m *AccessTokenRegistry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessTokenRegistry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessTokenRegistry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Issued) > 0 {
		for iNdEx := len(m.Issued) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Issued[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccessTokenRegistry(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Tenant) > 0 {
		i -= len(m.Tenant)
		copy(dAtA[i:], m.Tenant)
		i = encodeVarintAccessTokenRegistry(dAtA, i, uint64(len(m.Tenant)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccessTokenInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessTokenInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessTokenInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiresAt != 0 {
		i = encodeVarintAccessTokenRegistry(dAtA, i, uint64(m.ExpiresAt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintAccessTokenRegistry(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAccessTokenRegistry(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccessTokenRegistry(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccessTokenRegistry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessTokenRegistry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tenant)
	if l > 0 {
		n += 1 + l + sovAccessTokenRegistry(uint64(l))
	}
	if len(m.Issued) > 0 {
		for _, e := range m.Issued {
			l = e.Size()
			n += 1 + l + sovAccessTokenRegistry(uint64(l))
		}
	}
	return n
}

func (m *AccessTokenInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAccessTokenRegistry(uint64(l))
	}
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovAccessTokenRegistry(uint64(l))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovAccessTokenRegistry(uint64(m.ExpiresAt))
	}
	return n
}

func sovAccessTokenRegistry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccessTokenRegistry(x uint64) (n int) {
	return sovAccessTokenRegistry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessTokenRegistry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccessTokenRegistry
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
			return fmt.Errorf("proto: AccessTokenRegistry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessTokenRegistry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTokenRegistry
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
				return ErrInvalidLengthAccessTokenRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTokenRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tenant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTokenRegistry
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
				return ErrInvalidLengthAccessTokenRegistry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTokenRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issued = append(m.Issued, AccessTokenInfo{})
			if err := m.Issued[len(m.Issued)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccessTokenRegistry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccessTokenRegistry
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
func (m *AccessTokenInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccessTokenRegistry
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
			return fmt.Errorf("proto: AccessTokenInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessTokenInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTokenRegistry
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
				return ErrInvalidLengthAccessTokenRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTokenRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTokenRegistry
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
				return ErrInvalidLengthAccessTokenRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTokenRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTokenRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccessTokenRegistry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccessTokenRegistry
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
func skipAccessTokenRegistry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccessTokenRegistry
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
					return 0, ErrIntOverflowAccessTokenRegistry
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
					return 0, ErrIntOverflowAccessTokenRegistry
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
				return 0, ErrInvalidLengthAccessTokenRegistry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccessTokenRegistry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccessTokenRegistry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccessTokenRegistry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccessTokenRegistry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccessTokenRegistry = fmt.Errorf("proto: unexpected end of group")
)
