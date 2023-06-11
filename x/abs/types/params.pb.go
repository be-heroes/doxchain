// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/abs/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/be-heroes/doxchain/x/did/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	BlockExpireOffset       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=block_expire_offset,json=blockExpireOffset,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"block_expire_offset" yaml:"block_expire_offset"`
	BreakFactor             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=break_factor,json=breakFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"break_factor" yaml:"break_factor"`
	ThrottledRollingAverage github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=throttled_rolling_average,json=throttledRollingAverage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"throttled_rolling_average" yaml:"throttled_rolling_average"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5ba7663fa1ce94, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "doxchain.abs.v1beta1.Params")
}

func init() { proto.RegisterFile("doxchain/abs/v1beta1/params.proto", fileDescriptor_2f5ba7663fa1ce94) }

var fileDescriptor_2f5ba7663fa1ce94 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x93, 0x7f, 0xff, 0x14, 0x8c, 0x2e, 0xa6, 0x05, 0x6b, 0x87, 0xa4, 0x66, 0x10, 0x1d,
	0x9a, 0xa3, 0xb8, 0x75, 0xb3, 0x5a, 0x41, 0x10, 0x94, 0x8c, 0x2e, 0xe1, 0x2e, 0x77, 0x4d, 0x42,
	0x93, 0xbe, 0xe1, 0xee, 0x2c, 0x2d, 0xb8, 0xbb, 0x89, 0xa3, 0xa3, 0x1f, 0xa7, 0x63, 0x47, 0x71,
	0x08, 0xd2, 0x7e, 0x83, 0x7e, 0x02, 0xf1, 0x52, 0x43, 0x07, 0x1d, 0x3a, 0xdd, 0xbd, 0xef, 0xf3,
	0xf0, 0x7b, 0xde, 0xe3, 0x3d, 0xe3, 0x88, 0xc2, 0x24, 0x88, 0x70, 0x3c, 0x42, 0x98, 0x08, 0x34,
	0xee, 0x10, 0x26, 0x71, 0x07, 0x65, 0x98, 0xe3, 0x54, 0xb8, 0x19, 0x07, 0x09, 0x66, 0xfd, 0xc7,
	0xe2, 0x62, 0x22, 0xdc, 0xb5, 0xa5, 0x59, 0x0f, 0x21, 0x04, 0x65, 0x40, 0xdf, 0xb7, 0xc2, 0xdb,
	0xb4, 0x4a, 0x1c, 0x8d, 0x69, 0x89, 0xa3, 0x31, 0x2d, 0x74, 0xe7, 0xa9, 0x62, 0x54, 0xef, 0x14,
	0xdc, 0x7c, 0x34, 0x6a, 0x24, 0x81, 0x60, 0xe8, 0xb3, 0x49, 0x16, 0x73, 0xe6, 0xc3, 0x60, 0x20,
	0x98, 0x6c, 0xe8, 0x2d, 0xfd, 0x64, 0xa7, 0x77, 0x33, 0xcb, 0x6d, 0xed, 0x23, 0xb7, 0x8f, 0xc3,
	0x58, 0x46, 0x0f, 0xc4, 0x0d, 0x20, 0x45, 0x01, 0x88, 0x14, 0xc4, 0xfa, 0x68, 0x0b, 0x3a, 0x44,
	0x72, 0x9a, 0x31, 0xe1, 0x5e, 0x8f, 0xe4, 0x2a, 0xb7, 0x9b, 0x53, 0x9c, 0x26, 0x5d, 0xe7, 0x17,
	0xa4, 0xe3, 0xed, 0xab, 0x6e, 0x5f, 0x35, 0x6f, 0x55, 0xcf, 0x8c, 0x8c, 0x3d, 0xc2, 0x19, 0x1e,
	0xfa, 0x03, 0x1c, 0x48, 0xe0, 0x8d, 0x7f, 0x2a, 0xb6, 0xbf, 0x45, 0xec, 0x25, 0x0b, 0x56, 0xb9,
	0x5d, 0x5b, 0xc7, 0x6e, 0xb0, 0x1c, 0x6f, 0x57, 0x95, 0x57, 0xaa, 0x32, 0x9f, 0x75, 0xe3, 0x50,
	0x46, 0x1c, 0xa4, 0x4c, 0x18, 0xf5, 0x39, 0x24, 0x49, 0x3c, 0x0a, 0x7d, 0x3c, 0x66, 0x1c, 0x87,
	0xac, 0x51, 0x51, 0xb9, 0xde, 0xd6, 0xcf, 0x6d, 0x15, 0xb9, 0x7f, 0x82, 0x1d, 0xef, 0xa0, 0xd4,
	0xbc, 0x42, 0x3a, 0x2f, 0x94, 0xee, 0xff, 0xd7, 0x37, 0x5b, 0xeb, 0x5d, 0xcc, 0x16, 0x96, 0x3e,
	0x5f, 0x58, 0xfa, 0xe7, 0xc2, 0xd2, 0x5f, 0x96, 0x96, 0x36, 0x5f, 0x5a, 0xda, 0xfb, 0xd2, 0xd2,
	0xee, 0x4f, 0x37, 0x86, 0x20, 0xac, 0x1d, 0x31, 0x0e, 0x4c, 0xa0, 0x72, 0xb1, 0x13, 0xf5, 0x53,
	0xd4, 0x2c, 0xa4, 0xaa, 0xb6, 0x7a, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0x75, 0xa8, 0xec, 0x53,
	0x46, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ThrottledRollingAverage.Size()
		i -= size
		if _, err := m.ThrottledRollingAverage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.BreakFactor.Size()
		i -= size
		if _, err := m.BreakFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.BlockExpireOffset.Size()
		i -= size
		if _, err := m.BlockExpireOffset.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BlockExpireOffset.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.BreakFactor.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.ThrottledRollingAverage.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockExpireOffset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockExpireOffset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BreakFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BreakFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThrottledRollingAverage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ThrottledRollingAverage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
