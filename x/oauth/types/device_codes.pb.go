// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/oauth/device_codes.proto

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

type DeviceCodes struct {
	Tenant  string            `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Entries []DeviceCodeEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries"`
}

func (m *DeviceCodes) Reset()         { *m = DeviceCodes{} }
func (m *DeviceCodes) String() string { return proto.CompactTextString(m) }
func (*DeviceCodes) ProtoMessage()    {}
func (*DeviceCodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d76520aea23f034f, []int{0}
}
func (m *DeviceCodes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceCodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceCodes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceCodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceCodes.Merge(m, src)
}
func (m *DeviceCodes) XXX_Size() int {
	return m.Size()
}
func (m *DeviceCodes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceCodes.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceCodes proto.InternalMessageInfo

func (m *DeviceCodes) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *DeviceCodes) GetEntries() []DeviceCodeEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type DeviceCodeEntry struct {
	UserCode   string `protobuf:"bytes,1,opt,name=userCode,proto3" json:"userCode,omitempty"`
	DeviceCode string `protobuf:"bytes,2,opt,name=deviceCode,proto3" json:"deviceCode,omitempty"`
}

func (m *DeviceCodeEntry) Reset()         { *m = DeviceCodeEntry{} }
func (m *DeviceCodeEntry) String() string { return proto.CompactTextString(m) }
func (*DeviceCodeEntry) ProtoMessage()    {}
func (*DeviceCodeEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d76520aea23f034f, []int{1}
}
func (m *DeviceCodeEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceCodeEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceCodeEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceCodeEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceCodeEntry.Merge(m, src)
}
func (m *DeviceCodeEntry) XXX_Size() int {
	return m.Size()
}
func (m *DeviceCodeEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceCodeEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceCodeEntry proto.InternalMessageInfo

func (m *DeviceCodeEntry) GetUserCode() string {
	if m != nil {
		return m.UserCode
	}
	return ""
}

func (m *DeviceCodeEntry) GetDeviceCode() string {
	if m != nil {
		return m.DeviceCode
	}
	return ""
}

func init() {
	proto.RegisterType((*DeviceCodes)(nil), "beheroes.doxchain.oauth.DeviceCodes")
	proto.RegisterType((*DeviceCodeEntry)(nil), "beheroes.doxchain.oauth.DeviceCodeEntry")
}

func init() { proto.RegisterFile("doxchain/oauth/device_codes.proto", fileDescriptor_d76520aea23f034f) }

var fileDescriptor_d76520aea23f034f = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xc9, 0xaf, 0x48,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x4f, 0x49, 0x2d, 0xcb, 0x4c,
	0x4e, 0x8d, 0x4f, 0xce, 0x4f, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4f,
	0x4a, 0xcd, 0x48, 0x2d, 0xca, 0x4f, 0x2d, 0xd6, 0x83, 0xa9, 0xd5, 0x03, 0xab, 0x95, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd1, 0x07, 0xb1, 0x20, 0xca, 0x95, 0xf2, 0xb9, 0xb8, 0x5d, 0xc0,
	0x86, 0x38, 0x83, 0xcc, 0x10, 0x12, 0xe3, 0x62, 0x2b, 0x49, 0xcd, 0x4b, 0xcc, 0x2b, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x3c, 0xb8, 0xd8, 0x53, 0xf3, 0x4a, 0x8a, 0x32,
	0x53, 0x8b, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0x34, 0xf4, 0x70, 0xd8, 0xa3, 0x87, 0x30,
	0xce, 0x35, 0xaf, 0xa4, 0xa8, 0xd2, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x98, 0x76, 0x25,
	0x5f, 0x2e, 0x7e, 0x34, 0x15, 0x42, 0x52, 0x5c, 0x1c, 0xa5, 0xc5, 0xa9, 0x45, 0x20, 0x01, 0xa8,
	0xb5, 0x70, 0xbe, 0x90, 0x1c, 0x17, 0x57, 0x0a, 0x5c, 0xb9, 0x04, 0x13, 0x58, 0x16, 0x49, 0xc4,
	0xc9, 0xf5, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0,
	0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0x52, 0x75, 0x21, 0x8e, 0xd5, 0x87, 0x07,
	0x60, 0x05, 0x34, 0x08, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xa1, 0x61, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x30, 0xbf, 0x6d, 0xd7, 0x61, 0x01, 0x00, 0x00,
}

func (m *DeviceCodes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceCodes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceCodes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeviceCodes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Tenant) > 0 {
		i -= len(m.Tenant)
		copy(dAtA[i:], m.Tenant)
		i = encodeVarintDeviceCodes(dAtA, i, uint64(len(m.Tenant)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeviceCodeEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceCodeEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceCodeEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeviceCode) > 0 {
		i -= len(m.DeviceCode)
		copy(dAtA[i:], m.DeviceCode)
		i = encodeVarintDeviceCodes(dAtA, i, uint64(len(m.DeviceCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UserCode) > 0 {
		i -= len(m.UserCode)
		copy(dAtA[i:], m.UserCode)
		i = encodeVarintDeviceCodes(dAtA, i, uint64(len(m.UserCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeviceCodes(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeviceCodes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceCodes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tenant)
	if l > 0 {
		n += 1 + l + sovDeviceCodes(uint64(l))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovDeviceCodes(uint64(l))
		}
	}
	return n
}

func (m *DeviceCodeEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserCode)
	if l > 0 {
		n += 1 + l + sovDeviceCodes(uint64(l))
	}
	l = len(m.DeviceCode)
	if l > 0 {
		n += 1 + l + sovDeviceCodes(uint64(l))
	}
	return n
}

func sovDeviceCodes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeviceCodes(x uint64) (n int) {
	return sovDeviceCodes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceCodes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceCodes
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
			return fmt.Errorf("proto: DeviceCodes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceCodes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceCodes
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
				return ErrInvalidLengthDeviceCodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceCodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tenant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceCodes
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
				return ErrInvalidLengthDeviceCodes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceCodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, DeviceCodeEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceCodes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeviceCodes
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
func (m *DeviceCodeEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceCodes
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
			return fmt.Errorf("proto: DeviceCodeEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceCodeEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceCodes
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
				return ErrInvalidLengthDeviceCodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceCodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceCodes
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
				return ErrInvalidLengthDeviceCodes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceCodes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceCodes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeviceCodes
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
func skipDeviceCodes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceCodes
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
					return 0, ErrIntOverflowDeviceCodes
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
					return 0, ErrIntOverflowDeviceCodes
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
				return 0, ErrInvalidLengthDeviceCodes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeviceCodes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeviceCodes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeviceCodes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceCodes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeviceCodes = fmt.Errorf("proto: unexpected end of group")
)
