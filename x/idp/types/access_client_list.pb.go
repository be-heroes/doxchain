// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/idp/access_client_list.proto

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

type AccessClientList struct {
	Entries []AccessClientListEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries"`
}

func (m *AccessClientList) Reset()         { *m = AccessClientList{} }
func (m *AccessClientList) String() string { return proto.CompactTextString(m) }
func (*AccessClientList) ProtoMessage()    {}
func (*AccessClientList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6c374807e1e115, []int{0}
}
func (m *AccessClientList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessClientList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessClientList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessClientList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessClientList.Merge(m, src)
}
func (m *AccessClientList) XXX_Size() int {
	return m.Size()
}
func (m *AccessClientList) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessClientList.DiscardUnknown(m)
}

var xxx_messageInfo_AccessClientList proto.InternalMessageInfo

func (m *AccessClientList) GetEntries() []AccessClientListEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type AccessClientListEntry struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	CanRead    bool   `protobuf:"varint,2,opt,name=canRead,proto3" json:"canRead,omitempty"`
	CanWrite   bool   `protobuf:"varint,3,opt,name=canWrite,proto3" json:"canWrite,omitempty"`
	CanExecute bool   `protobuf:"varint,4,opt,name=canExecute,proto3" json:"canExecute,omitempty"`
}

func (m *AccessClientListEntry) Reset()         { *m = AccessClientListEntry{} }
func (m *AccessClientListEntry) String() string { return proto.CompactTextString(m) }
func (*AccessClientListEntry) ProtoMessage()    {}
func (*AccessClientListEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6c374807e1e115, []int{1}
}
func (m *AccessClientListEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessClientListEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessClientListEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessClientListEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessClientListEntry.Merge(m, src)
}
func (m *AccessClientListEntry) XXX_Size() int {
	return m.Size()
}
func (m *AccessClientListEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessClientListEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AccessClientListEntry proto.InternalMessageInfo

func (m *AccessClientListEntry) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *AccessClientListEntry) GetCanRead() bool {
	if m != nil {
		return m.CanRead
	}
	return false
}

func (m *AccessClientListEntry) GetCanWrite() bool {
	if m != nil {
		return m.CanWrite
	}
	return false
}

func (m *AccessClientListEntry) GetCanExecute() bool {
	if m != nil {
		return m.CanExecute
	}
	return false
}

func init() {
	proto.RegisterType((*AccessClientList)(nil), "beheroes.doxchain.idp.AccessClientList")
	proto.RegisterType((*AccessClientListEntry)(nil), "beheroes.doxchain.idp.AccessClientListEntry")
}

func init() {
	proto.RegisterFile("doxchain/idp/access_client_list.proto", fileDescriptor_ae6c374807e1e115)
}

var fileDescriptor_ae6c374807e1e115 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x1b, 0x37, 0xdc, 0x8c, 0x17, 0x29, 0x0e, 0xc2, 0x0e, 0xb1, 0x0c, 0x84, 0x0a, 0x9a,
	0x82, 0x3e, 0x81, 0x1b, 0xbb, 0xed, 0xd4, 0x8b, 0xe0, 0x65, 0xa6, 0xe9, 0x9f, 0x36, 0x30, 0x93,
	0x92, 0x64, 0xd0, 0x3d, 0x81, 0x57, 0x1f, 0x6b, 0xc7, 0x1d, 0x3d, 0x89, 0xb4, 0x2f, 0x22, 0x66,
	0x74, 0xc8, 0xd8, 0x2d, 0x5f, 0xbe, 0x5f, 0x7e, 0x90, 0x0f, 0xdf, 0xe6, 0xba, 0x16, 0x25, 0x97,
	0x2a, 0x91, 0x79, 0x95, 0x70, 0x21, 0xc0, 0xda, 0xa5, 0x58, 0x49, 0x50, 0x6e, 0xb9, 0x92, 0xd6,
	0xb1, 0xca, 0x68, 0xa7, 0xc3, 0x51, 0x06, 0x25, 0x18, 0x0d, 0x96, 0x75, 0x3c, 0x93, 0x79, 0x35,
	0xbe, 0x2e, 0x74, 0xa1, 0x3d, 0x91, 0xfc, 0x9d, 0xf6, 0xf0, 0xe4, 0x0d, 0x5f, 0x3d, 0x7b, 0xd1,
	0xcc, 0x7b, 0x16, 0xd2, 0xba, 0x70, 0x81, 0x07, 0xa0, 0x9c, 0x91, 0x60, 0x09, 0x8a, 0x7a, 0xf1,
	0xe5, 0xe3, 0x3d, 0x3b, 0xa9, 0x64, 0xc7, 0x2f, 0xe7, 0xca, 0x99, 0xcd, 0xb4, 0xbf, 0xfd, 0xbe,
	0x09, 0xd2, 0x4e, 0x31, 0xf9, 0x40, 0x78, 0x74, 0x12, 0x0c, 0x09, 0x1e, 0x08, 0x03, 0xdc, 0x69,
	0x43, 0x50, 0x84, 0xe2, 0x8b, 0xb4, 0x8b, 0xbe, 0xe1, 0x2a, 0x05, 0x9e, 0x93, 0xb3, 0x08, 0xc5,
	0xc3, 0xb4, 0x8b, 0xe1, 0x18, 0x0f, 0x05, 0x57, 0x2f, 0x46, 0x3a, 0x20, 0x3d, 0x5f, 0x1d, 0x72,
	0x48, 0x31, 0x16, 0x5c, 0xcd, 0x6b, 0x10, 0x6b, 0x07, 0xa4, 0xef, 0xdb, 0x7f, 0x37, 0xd3, 0xd9,
	0xb6, 0xa1, 0x68, 0xd7, 0x50, 0xf4, 0xd3, 0x50, 0xf4, 0xd9, 0xd2, 0x60, 0xd7, 0xd2, 0xe0, 0xab,
	0xa5, 0xc1, 0xeb, 0x5d, 0x21, 0x5d, 0xb9, 0xce, 0x98, 0xd0, 0xef, 0x49, 0x06, 0x0f, 0xfb, 0xbf,
	0x26, 0x87, 0xb9, 0x6b, 0x3f, 0xb8, 0xdb, 0x54, 0x60, 0xb3, 0x73, 0xbf, 0xdb, 0xd3, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x72, 0x3d, 0x95, 0x8f, 0x8d, 0x01, 0x00, 0x00,
}

func (m *AccessClientList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessClientList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessClientList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintAccessClientList(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AccessClientListEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessClientListEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessClientListEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CanExecute {
		i--
		if m.CanExecute {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.CanWrite {
		i--
		if m.CanWrite {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.CanRead {
		i--
		if m.CanRead {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAccessClientList(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccessClientList(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccessClientList(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessClientList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovAccessClientList(uint64(l))
		}
	}
	return n
}

func (m *AccessClientListEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAccessClientList(uint64(l))
	}
	if m.CanRead {
		n += 2
	}
	if m.CanWrite {
		n += 2
	}
	if m.CanExecute {
		n += 2
	}
	return n
}

func sovAccessClientList(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccessClientList(x uint64) (n int) {
	return sovAccessClientList(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessClientList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccessClientList
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
			return fmt.Errorf("proto: AccessClientList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessClientList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessClientList
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
				return ErrInvalidLengthAccessClientList
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccessClientList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, AccessClientListEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccessClientList(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccessClientList
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
func (m *AccessClientListEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccessClientList
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
			return fmt.Errorf("proto: AccessClientListEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessClientListEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessClientList
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
				return ErrInvalidLengthAccessClientList
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessClientList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanRead", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessClientList
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
			m.CanRead = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanWrite", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessClientList
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
			m.CanWrite = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanExecute", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessClientList
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
			m.CanExecute = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAccessClientList(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccessClientList
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
func skipAccessClientList(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccessClientList
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
					return 0, ErrIntOverflowAccessClientList
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
					return 0, ErrIntOverflowAccessClientList
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
				return 0, ErrInvalidLengthAccessClientList
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccessClientList
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccessClientList
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccessClientList        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccessClientList          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccessClientList = fmt.Errorf("proto: unexpected end of group")
)
