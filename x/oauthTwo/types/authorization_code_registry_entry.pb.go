// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/oauthtwo/v1beta1/authorization_code_registry_entry.proto

package types

import (
	fmt "fmt"
	types "github.com/be-heroes/doxchain/x/did/types"
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

type AuthorizationCodeRegistryEntry struct {
	Owner             types.Did `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner"`
	AuthorizationCode string    `protobuf:"bytes,2,opt,name=authorizationCode,proto3" json:"authorizationCode,omitempty"`
	ExpiresAt         int64     `protobuf:"varint,3,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

func (m *AuthorizationCodeRegistryEntry) Reset()         { *m = AuthorizationCodeRegistryEntry{} }
func (m *AuthorizationCodeRegistryEntry) String() string { return proto.CompactTextString(m) }
func (*AuthorizationCodeRegistryEntry) ProtoMessage()    {}
func (*AuthorizationCodeRegistryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_1763526c68c589aa, []int{0}
}
func (m *AuthorizationCodeRegistryEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthorizationCodeRegistryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthorizationCodeRegistryEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthorizationCodeRegistryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationCodeRegistryEntry.Merge(m, src)
}
func (m *AuthorizationCodeRegistryEntry) XXX_Size() int {
	return m.Size()
}
func (m *AuthorizationCodeRegistryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationCodeRegistryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationCodeRegistryEntry proto.InternalMessageInfo

func (m *AuthorizationCodeRegistryEntry) GetOwner() types.Did {
	if m != nil {
		return m.Owner
	}
	return types.Did{}
}

func (m *AuthorizationCodeRegistryEntry) GetAuthorizationCode() string {
	if m != nil {
		return m.AuthorizationCode
	}
	return ""
}

func (m *AuthorizationCodeRegistryEntry) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthorizationCodeRegistryEntry)(nil), "doxchain.oauthtwo.v1beta1.AuthorizationCodeRegistryEntry")
}

func init() {
	proto.RegisterFile("doxchain/oauthtwo/v1beta1/authorization_code_registry_entry.proto", fileDescriptor_1763526c68c589aa)
}

var fileDescriptor_1763526c68c589aa = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x63, 0x0a, 0x48, 0x0d, 0x13, 0x11, 0x43, 0x5b, 0x21, 0x13, 0x31, 0x65, 0x00, 0x5b,
	0x05, 0x71, 0x80, 0x14, 0x90, 0x98, 0x33, 0xb2, 0x44, 0x4e, 0xfc, 0x94, 0x78, 0x20, 0x2f, 0x72,
	0x5c, 0x9a, 0x70, 0x0a, 0xee, 0xc0, 0x65, 0x3a, 0x76, 0x64, 0x42, 0x28, 0xb9, 0x08, 0x4a, 0x93,
	0xa6, 0x08, 0x36, 0xdb, 0xef, 0xf7, 0xf7, 0xf9, 0xb7, 0xed, 0x4b, 0x2c, 0xe3, 0x54, 0xa8, 0x8c,
	0xa3, 0x58, 0x9a, 0xd4, 0xac, 0x90, 0xbf, 0xce, 0x23, 0x30, 0x62, 0xce, 0xdb, 0x3d, 0x6a, 0xf5,
	0x26, 0x8c, 0xc2, 0x2c, 0x8c, 0x51, 0x42, 0xa8, 0x21, 0x51, 0x85, 0xd1, 0x55, 0x08, 0x99, 0xd1,
	0x15, 0xcb, 0x35, 0x1a, 0x74, 0xa6, 0x3b, 0x04, 0xdb, 0x21, 0x58, 0x8f, 0x98, 0x9d, 0x25, 0x98,
	0xe0, 0x36, 0xc5, 0xdb, 0x55, 0x77, 0x61, 0x46, 0x07, 0xa7, 0x54, 0x72, 0xd0, 0x49, 0x25, 0xbb,
	0xf9, 0xe5, 0x07, 0xb1, 0xa9, 0xff, 0x5b, 0x7e, 0x8f, 0x12, 0x82, 0x5e, 0xfd, 0xd8, 0x9a, 0x9d,
	0x3b, 0xfb, 0x08, 0x57, 0x19, 0xe8, 0x09, 0x71, 0x89, 0x77, 0x72, 0x33, 0x65, 0xc3, 0x1b, 0x5a,
	0x4c, 0x8f, 0x64, 0x0f, 0x4a, 0x2e, 0x0e, 0xd7, 0x5f, 0x17, 0x56, 0xd0, 0xa5, 0x9d, 0x2b, 0xfb,
	0x54, 0xfc, 0x05, 0x4f, 0x0e, 0x5c, 0xe2, 0x8d, 0x83, 0xff, 0x03, 0xe7, 0xdc, 0x1e, 0x43, 0x99,
	0x2b, 0x0d, 0x85, 0x6f, 0x26, 0x23, 0x97, 0x78, 0xa3, 0x60, 0x7f, 0xb0, 0x78, 0x5a, 0xd7, 0x94,
	0x6c, 0x6a, 0x4a, 0xbe, 0x6b, 0x4a, 0xde, 0x1b, 0x6a, 0x6d, 0x1a, 0x6a, 0x7d, 0x36, 0xd4, 0x7a,
	0x66, 0x89, 0x32, 0xe9, 0x32, 0x62, 0x31, 0xbe, 0xf0, 0x08, 0xae, 0x53, 0xd0, 0x08, 0x05, 0x1f,
	0x4a, 0x97, 0xfb, 0xaf, 0x36, 0x55, 0x0e, 0x45, 0x74, 0xbc, 0xad, 0x7d, 0xfb, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0x0b, 0xcf, 0xf5, 0x8c, 0x01, 0x00, 0x00,
}

func (m *AuthorizationCodeRegistryEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthorizationCodeRegistryEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthorizationCodeRegistryEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiresAt != 0 {
		i = encodeVarintAuthorizationCodeRegistryEntry(dAtA, i, uint64(m.ExpiresAt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AuthorizationCode) > 0 {
		i -= len(m.AuthorizationCode)
		copy(dAtA[i:], m.AuthorizationCode)
		i = encodeVarintAuthorizationCodeRegistryEntry(dAtA, i, uint64(len(m.AuthorizationCode)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Owner.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuthorizationCodeRegistryEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAuthorizationCodeRegistryEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthorizationCodeRegistryEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuthorizationCodeRegistryEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Owner.Size()
	n += 1 + l + sovAuthorizationCodeRegistryEntry(uint64(l))
	l = len(m.AuthorizationCode)
	if l > 0 {
		n += 1 + l + sovAuthorizationCodeRegistryEntry(uint64(l))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovAuthorizationCodeRegistryEntry(uint64(m.ExpiresAt))
	}
	return n
}

func sovAuthorizationCodeRegistryEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthorizationCodeRegistryEntry(x uint64) (n int) {
	return sovAuthorizationCodeRegistryEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthorizationCodeRegistryEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthorizationCodeRegistryEntry
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
			return fmt.Errorf("proto: AuthorizationCodeRegistryEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthorizationCodeRegistryEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorizationCodeRegistryEntry
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
				return ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizationCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorizationCodeRegistryEntry
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
				return ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizationCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorizationCodeRegistryEntry
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
			skippy, err := skipAuthorizationCodeRegistryEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthorizationCodeRegistryEntry
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
func skipAuthorizationCodeRegistryEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthorizationCodeRegistryEntry
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
					return 0, ErrIntOverflowAuthorizationCodeRegistryEntry
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
					return 0, ErrIntOverflowAuthorizationCodeRegistryEntry
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
				return 0, ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthorizationCodeRegistryEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthorizationCodeRegistryEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthorizationCodeRegistryEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthorizationCodeRegistryEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthorizationCodeRegistryEntry = fmt.Errorf("proto: unexpected end of group")
)
