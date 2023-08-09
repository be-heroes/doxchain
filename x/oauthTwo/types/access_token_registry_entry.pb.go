// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/oauthtwo/v1beta1/access_token_registry_entry.proto

package types

import (
	fmt "fmt"
	types "github.com/be-heroes/doxchain/x/did/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type AccessTokenRegistryEntry struct {
	Owner     types.Did `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner"`
	Jti       string    `protobuf:"bytes,2,opt,name=jti,proto3" json:"jti,omitempty"`
	ExpiresAt int64     `protobuf:"varint,3,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

func (m *AccessTokenRegistryEntry) Reset()         { *m = AccessTokenRegistryEntry{} }
func (m *AccessTokenRegistryEntry) String() string { return proto.CompactTextString(m) }
func (*AccessTokenRegistryEntry) ProtoMessage()    {}
func (*AccessTokenRegistryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e668325ff912be35, []int{0}
}
func (m *AccessTokenRegistryEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessTokenRegistryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessTokenRegistryEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessTokenRegistryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTokenRegistryEntry.Merge(m, src)
}
func (m *AccessTokenRegistryEntry) XXX_Size() int {
	return m.Size()
}
func (m *AccessTokenRegistryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTokenRegistryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTokenRegistryEntry proto.InternalMessageInfo

func (m *AccessTokenRegistryEntry) GetOwner() types.Did {
	if m != nil {
		return m.Owner
	}
	return types.Did{}
}

func (m *AccessTokenRegistryEntry) GetJti() string {
	if m != nil {
		return m.Jti
	}
	return ""
}

func (m *AccessTokenRegistryEntry) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*AccessTokenRegistryEntry)(nil), "doxchain.oauthtwo.v1beta1.AccessTokenRegistryEntry")
}

func init() {
	proto.RegisterFile("doxchain/oauthtwo/v1beta1/access_token_registry_entry.proto", fileDescriptor_e668325ff912be35)
}

var fileDescriptor_e668325ff912be35 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x63, 0x0a, 0x48, 0x0d, 0x0b, 0x8a, 0x18, 0xd2, 0x0a, 0x99, 0x88, 0x29, 0x0b, 0xb6,
	0x0a, 0x62, 0x62, 0x6a, 0x05, 0x12, 0x73, 0xc4, 0xc4, 0x12, 0x25, 0xf1, 0x29, 0x31, 0x88, 0x38,
	0xb2, 0xaf, 0x34, 0x19, 0x79, 0x03, 0x1e, 0xab, 0x63, 0x47, 0x26, 0x84, 0x92, 0x17, 0x41, 0xf9,
	0xd3, 0x74, 0x3b, 0xd9, 0xf7, 0xfb, 0x7e, 0xfa, 0xce, 0x7e, 0x10, 0xaa, 0x4c, 0xb2, 0x48, 0xe6,
	0x5c, 0x45, 0x6b, 0xcc, 0x70, 0xa3, 0xf8, 0xe7, 0x22, 0x06, 0x8c, 0x16, 0x3c, 0x4a, 0x12, 0x30,
	0x26, 0x44, 0xf5, 0x0e, 0x79, 0xa8, 0x21, 0x95, 0x06, 0x75, 0x15, 0x42, 0x8e, 0xba, 0x62, 0x85,
	0x56, 0xa8, 0x9c, 0xd9, 0x1e, 0x66, 0x7b, 0x98, 0x0d, 0xf0, 0xfc, 0x22, 0x55, 0xa9, 0xea, 0xb6,
	0x78, 0x3b, 0xf5, 0xc0, 0x9c, 0x8e, 0x36, 0x21, 0xc5, 0x28, 0x12, 0x52, 0xf4, 0xff, 0xd7, 0x5f,
	0xc4, 0x76, 0x97, 0x9d, 0xf6, 0xa5, 0xb5, 0x06, 0x83, 0xf4, 0xa9, 0x75, 0x3a, 0xf7, 0xf6, 0x89,
	0xda, 0xe4, 0xa0, 0x5d, 0xe2, 0x11, 0xff, 0xec, 0x76, 0xc6, 0x46, 0x7b, 0x1b, 0x30, 0x84, 0xb1,
	0x47, 0x29, 0x56, 0xc7, 0xdb, 0xdf, 0x2b, 0x2b, 0xe8, 0xb7, 0x9d, 0x73, 0x7b, 0xf2, 0x86, 0xd2,
	0x3d, 0xf2, 0x88, 0x3f, 0x0d, 0xda, 0xd1, 0xb9, 0xb4, 0xa7, 0x50, 0x16, 0x52, 0x83, 0x59, 0xa2,
	0x3b, 0xf1, 0x88, 0x3f, 0x09, 0x0e, 0x0f, 0xab, 0xe7, 0x6d, 0x4d, 0xc9, 0xae, 0xa6, 0xe4, 0xaf,
	0xa6, 0xe4, 0xbb, 0xa1, 0xd6, 0xae, 0xa1, 0xd6, 0x4f, 0x43, 0xad, 0x57, 0x96, 0x4a, 0xcc, 0xd6,
	0x31, 0x4b, 0xd4, 0x07, 0x8f, 0xe1, 0x26, 0x03, 0xad, 0xc0, 0xf0, 0xb1, 0x52, 0x79, 0x38, 0x21,
	0x56, 0x05, 0x98, 0xf8, 0xb4, 0x2b, 0x75, 0xf7, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x12, 0xc1, 0xf0,
	0x7c, 0x64, 0x01, 0x00, 0x00,
}

func (m *AccessTokenRegistryEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessTokenRegistryEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessTokenRegistryEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiresAt != 0 {
		i = encodeVarintAccessTokenRegistryEntry(dAtA, i, uint64(m.ExpiresAt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Jti) > 0 {
		i -= len(m.Jti)
		copy(dAtA[i:], m.Jti)
		i = encodeVarintAccessTokenRegistryEntry(dAtA, i, uint64(len(m.Jti)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Owner.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAccessTokenRegistryEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAccessTokenRegistryEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccessTokenRegistryEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessTokenRegistryEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Owner.Size()
	n += 1 + l + sovAccessTokenRegistryEntry(uint64(l))
	l = len(m.Jti)
	if l > 0 {
		n += 1 + l + sovAccessTokenRegistryEntry(uint64(l))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovAccessTokenRegistryEntry(uint64(m.ExpiresAt))
	}
	return n
}

func sovAccessTokenRegistryEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccessTokenRegistryEntry(x uint64) (n int) {
	return sovAccessTokenRegistryEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessTokenRegistryEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccessTokenRegistryEntry
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
			return fmt.Errorf("proto: AccessTokenRegistryEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessTokenRegistryEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTokenRegistryEntry
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
				return ErrInvalidLengthAccessTokenRegistryEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTokenRegistryEntry
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
				return fmt.Errorf("proto: wrong wireType = %d for field Jti", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTokenRegistryEntry
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
				return ErrInvalidLengthAccessTokenRegistryEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTokenRegistryEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Jti = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTokenRegistryEntry
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
			skippy, err := skipAccessTokenRegistryEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccessTokenRegistryEntry
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
func skipAccessTokenRegistryEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccessTokenRegistryEntry
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
					return 0, ErrIntOverflowAccessTokenRegistryEntry
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
					return 0, ErrIntOverflowAccessTokenRegistryEntry
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
				return 0, ErrInvalidLengthAccessTokenRegistryEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccessTokenRegistryEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccessTokenRegistryEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccessTokenRegistryEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccessTokenRegistryEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccessTokenRegistryEntry = fmt.Errorf("proto: unexpected end of group")
)
