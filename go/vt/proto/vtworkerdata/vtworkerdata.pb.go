// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vtworkerdata.proto

package vtworkerdata

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/golang/protobuf/proto"
	logutil "vitess.io/vitess/go/vt/proto/logutil"
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

// ExecuteVtworkerCommandRequest is the payload for ExecuteVtworkerCommand.
type ExecuteVtworkerCommandRequest struct {
	Args                 []string `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteVtworkerCommandRequest) Reset()         { *m = ExecuteVtworkerCommandRequest{} }
func (m *ExecuteVtworkerCommandRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteVtworkerCommandRequest) ProtoMessage()    {}
func (*ExecuteVtworkerCommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32a791ab99179e8e, []int{0}
}
func (m *ExecuteVtworkerCommandRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecuteVtworkerCommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExecuteVtworkerCommandRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExecuteVtworkerCommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteVtworkerCommandRequest.Merge(m, src)
}
func (m *ExecuteVtworkerCommandRequest) XXX_Size() int {
	return m.Size()
}
func (m *ExecuteVtworkerCommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteVtworkerCommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteVtworkerCommandRequest proto.InternalMessageInfo

func (m *ExecuteVtworkerCommandRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

// ExecuteVtworkerCommandResponse is streamed back by ExecuteVtworkerCommand.
type ExecuteVtworkerCommandResponse struct {
	Event                *logutil.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExecuteVtworkerCommandResponse) Reset()         { *m = ExecuteVtworkerCommandResponse{} }
func (m *ExecuteVtworkerCommandResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteVtworkerCommandResponse) ProtoMessage()    {}
func (*ExecuteVtworkerCommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32a791ab99179e8e, []int{1}
}
func (m *ExecuteVtworkerCommandResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecuteVtworkerCommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExecuteVtworkerCommandResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExecuteVtworkerCommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteVtworkerCommandResponse.Merge(m, src)
}
func (m *ExecuteVtworkerCommandResponse) XXX_Size() int {
	return m.Size()
}
func (m *ExecuteVtworkerCommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteVtworkerCommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteVtworkerCommandResponse proto.InternalMessageInfo

func (m *ExecuteVtworkerCommandResponse) GetEvent() *logutil.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecuteVtworkerCommandRequest)(nil), "vtworkerdata.ExecuteVtworkerCommandRequest")
	proto.RegisterType((*ExecuteVtworkerCommandResponse)(nil), "vtworkerdata.ExecuteVtworkerCommandResponse")
}

func init() { proto.RegisterFile("vtworkerdata.proto", fileDescriptor_32a791ab99179e8e) }

var fileDescriptor_32a791ab99179e8e = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2b, 0x29, 0xcf,
	0x2f, 0xca, 0x4e, 0x2d, 0x4a, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x93, 0xe2, 0xcd, 0xc9, 0x4f, 0x2f, 0x2d, 0xc9, 0xcc, 0x81, 0x48, 0x2a, 0x19, 0x73,
	0xc9, 0xba, 0x56, 0xa4, 0x26, 0x97, 0x96, 0xa4, 0x86, 0x41, 0x55, 0x39, 0xe7, 0xe7, 0xe6, 0x26,
	0xe6, 0xa5, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0x24, 0x16, 0xa5,
	0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x6e, 0x5c, 0x72, 0xb8, 0x34,
	0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xa9, 0x70, 0xb1, 0xa6, 0x96, 0xa5, 0xe6, 0x95, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xf1, 0xe9, 0xc1, 0x6c, 0x75, 0x05, 0x89, 0x06, 0x41, 0x24,
	0x9d, 0xac, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19,
	0x8f, 0xe5, 0x18, 0xa2, 0x34, 0xcb, 0x32, 0x4b, 0x52, 0x8b, 0x8b, 0xf5, 0x32, 0xf3, 0xf5, 0x21,
	0x2c, 0xfd, 0xf4, 0x7c, 0xfd, 0xb2, 0x12, 0x7d, 0xb0, 0x63, 0xf5, 0x91, 0x3d, 0x92, 0xc4, 0x06,
	0x16, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x38, 0xbf, 0x62, 0x49, 0xf3, 0x00, 0x00, 0x00,
}

func (m *ExecuteVtworkerCommandRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecuteVtworkerCommandRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExecuteVtworkerCommandRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Args) > 0 {
		for iNdEx := len(m.Args) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Args[iNdEx])
			copy(dAtA[i:], m.Args[iNdEx])
			i = encodeVarintVtworkerdata(dAtA, i, uint64(len(m.Args[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ExecuteVtworkerCommandResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecuteVtworkerCommandResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExecuteVtworkerCommandResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Event != nil {
		{
			size, err := m.Event.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVtworkerdata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVtworkerdata(dAtA []byte, offset int, v uint64) int {
	offset -= sovVtworkerdata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExecuteVtworkerCommandRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovVtworkerdata(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExecuteVtworkerCommandResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovVtworkerdata(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVtworkerdata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVtworkerdata(x uint64) (n int) {
	return sovVtworkerdata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExecuteVtworkerCommandRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVtworkerdata
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
			return fmt.Errorf("proto: ExecuteVtworkerCommandRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecuteVtworkerCommandRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVtworkerdata
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
				return ErrInvalidLengthVtworkerdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVtworkerdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVtworkerdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVtworkerdata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVtworkerdata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExecuteVtworkerCommandResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVtworkerdata
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
			return fmt.Errorf("proto: ExecuteVtworkerCommandResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecuteVtworkerCommandResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVtworkerdata
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
				return ErrInvalidLengthVtworkerdata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVtworkerdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &logutil.Event{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVtworkerdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVtworkerdata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVtworkerdata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVtworkerdata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVtworkerdata
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
					return 0, ErrIntOverflowVtworkerdata
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
					return 0, ErrIntOverflowVtworkerdata
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
				return 0, ErrInvalidLengthVtworkerdata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVtworkerdata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVtworkerdata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVtworkerdata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVtworkerdata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVtworkerdata = fmt.Errorf("proto: unexpected end of group")
)
