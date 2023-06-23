// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/oauthtwo/v1beta1/genesis.proto

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

type GenesisState struct {
	Params                      Params                      `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	AccessTokenRegistries       []AccessTokenRegistry       `protobuf:"bytes,2,rep,name=accessTokenRegistries,proto3" json:"accessTokenRegistries"`
	AuthorizationCodeRegistries []AuthorizationCodeRegistry `protobuf:"bytes,3,rep,name=authorizationCodeRegistries,proto3" json:"authorizationCodeRegistries"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_440be9fe4c5fbf17, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetAccessTokenRegistries() []AccessTokenRegistry {
	if m != nil {
		return m.AccessTokenRegistries
	}
	return nil
}

func (m *GenesisState) GetAuthorizationCodeRegistries() []AuthorizationCodeRegistry {
	if m != nil {
		return m.AuthorizationCodeRegistries
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "doxchain.oauthtwo.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("doxchain/oauthtwo/v1beta1/genesis.proto", fileDescriptor_440be9fe4c5fbf17)
}

var fileDescriptor_440be9fe4c5fbf17 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbd, 0x4e, 0x02, 0x41,
	0x10, 0xc7, 0xef, 0xc0, 0x50, 0x1c, 0x56, 0x17, 0x4d, 0x10, 0x93, 0x15, 0x2d, 0x94, 0xc6, 0xdd,
	0x80, 0x5a, 0x59, 0x18, 0xb1, 0xd0, 0xd2, 0xa0, 0x95, 0x0d, 0xd9, 0x3b, 0x26, 0x77, 0xab, 0xe1,
	0x86, 0xec, 0x0e, 0x0a, 0xc6, 0x87, 0xf0, 0x75, 0x7c, 0x03, 0x4a, 0x4a, 0x2b, 0x63, 0xe0, 0x45,
	0x0c, 0xf7, 0xe1, 0xf7, 0x5d, 0x37, 0xc5, 0xff, 0xf7, 0xfb, 0xcf, 0x64, 0x9c, 0xbd, 0x3e, 0x8e,
	0xfd, 0x50, 0xaa, 0x48, 0xa0, 0x1c, 0x51, 0x48, 0x0f, 0x28, 0xee, 0x5b, 0x1e, 0x90, 0x6c, 0x89,
	0x00, 0x22, 0x30, 0xca, 0xf0, 0xa1, 0x46, 0x42, 0x77, 0x23, 0x0b, 0xf2, 0x2c, 0xc8, 0xd3, 0x60,
	0x7d, 0x2d, 0xc0, 0x00, 0xe3, 0x94, 0x58, 0x4e, 0x09, 0x50, 0xdf, 0xcd, 0x37, 0x0f, 0xa5, 0x96,
	0x83, 0x54, 0x5c, 0x3f, 0xca, 0xcf, 0x49, 0xdf, 0x07, 0x63, 0x7a, 0x84, 0x77, 0x10, 0xf5, 0x34,
	0x04, 0xca, 0x90, 0x9e, 0xa4, 0xd8, 0x71, 0x01, 0x36, 0xa2, 0x10, 0xb5, 0x7a, 0x94, 0xa4, 0x30,
	0xea, 0xf9, 0xd8, 0x87, 0x5f, 0xf0, 0xce, 0x4b, 0xc9, 0x59, 0x3d, 0x4f, 0xce, 0xbb, 0x22, 0x49,
	0xe0, 0x9e, 0x38, 0x95, 0x64, 0xa9, 0x9a, 0xdd, 0xb0, 0x9b, 0xd5, 0xf6, 0x36, 0xcf, 0x3d, 0x97,
	0x5f, 0xc6, 0xc1, 0xce, 0xca, 0xf4, 0x6d, 0xcb, 0xea, 0xa6, 0x98, 0x7b, 0xeb, 0xac, 0x27, 0xdb,
	0x5e, 0x2f, 0x97, 0xed, 0x26, 0x75, 0x0a, 0x4c, 0xad, 0xd4, 0x28, 0x37, 0xab, 0x6d, 0x5e, 0xe0,
	0x3b, 0xfd, 0xc3, 0x4d, 0x52, 0xf9, 0xff, 0x4a, 0xf7, 0xc9, 0xd9, 0xfc, 0x71, 0xe2, 0x19, 0xf6,
	0xe1, 0x5b, 0x63, 0x39, 0x6e, 0x3c, 0x2c, 0x6a, 0xcc, 0xa1, 0xb3, 0xde, 0x22, 0x7d, 0xe7, 0x62,
	0x3a, 0x67, 0xf6, 0x6c, 0xce, 0xec, 0xf7, 0x39, 0xb3, 0x9f, 0x17, 0xcc, 0x9a, 0x2d, 0x98, 0xf5,
	0xba, 0x60, 0xd6, 0x0d, 0x0f, 0x14, 0x85, 0x23, 0x8f, 0xfb, 0x38, 0x10, 0x1e, 0xec, 0x87, 0xa0,
	0x11, 0x8c, 0xf8, 0xfc, 0xd3, 0xf8, 0xeb, 0x53, 0x34, 0x19, 0x82, 0xf1, 0x2a, 0xf1, 0x33, 0x0e,
	0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x8b, 0x75, 0x5e, 0x84, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthorizationCodeRegistries) > 0 {
		for iNdEx := len(m.AuthorizationCodeRegistries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AuthorizationCodeRegistries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AccessTokenRegistries) > 0 {
		for iNdEx := len(m.AccessTokenRegistries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccessTokenRegistries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.AccessTokenRegistries) > 0 {
		for _, e := range m.AccessTokenRegistries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AuthorizationCodeRegistries) > 0 {
		for _, e := range m.AuthorizationCodeRegistries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessTokenRegistries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessTokenRegistries = append(m.AccessTokenRegistries, AccessTokenRegistry{})
			if err := m.AccessTokenRegistries[len(m.AccessTokenRegistries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizationCodeRegistries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizationCodeRegistries = append(m.AuthorizationCodeRegistries, AuthorizationCodeRegistry{})
			if err := m.AuthorizationCodeRegistries[len(m.AuthorizationCodeRegistries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
