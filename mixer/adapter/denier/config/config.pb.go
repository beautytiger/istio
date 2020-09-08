// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/denier/config/config.proto

// The `denier` adapter is designed to always return a denial to precondition
// checks. You can specify the exact error to return for these denials.
//
// This adapter supports the [checknothing template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/checknothing/),
// the [listentry template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/listentry/),
// and the [quota template](https://istio.io/docs/reference/config/policy-and-telemetry/templates/quota/).

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	rpc "istio.io/gogo-genproto/googleapis/google/rpc"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Configuration format for the Denier adapter.
type Params struct {
	// The error to return when denying a request.
	Status rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status"`
	// The duration for which the denial is valid.
	ValidDuration time.Duration `protobuf:"bytes,2,opt,name=valid_duration,json=validDuration,proto3,stdduration" json:"valid_duration"`
	// The number of times the denial may be used.
	ValidUseCount int32 `protobuf:"varint,3,opt,name=valid_use_count,json=validUseCount,proto3" json:"valid_use_count,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08be602ddb157bc, []int{0}
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
	proto.RegisterType((*Params)(nil), "adapter.denier.config.Params")
}

func init() {
	proto.RegisterFile("mixer/adapter/denier/config/config.proto", fileDescriptor_c08be602ddb157bc)
}

var fileDescriptor_c08be602ddb157bc = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0x3d, 0x4e, 0xc3, 0x30,
	0x18, 0x86, 0x6d, 0x7e, 0x22, 0x14, 0x04, 0x48, 0x11, 0x88, 0xd2, 0xe1, 0x6b, 0xc5, 0x80, 0x3a,
	0xd9, 0x08, 0x2e, 0x80, 0x02, 0x13, 0x13, 0x2a, 0x62, 0x61, 0xa9, 0xdc, 0xc4, 0x8d, 0x22, 0xb5,
	0x71, 0xe4, 0x38, 0x88, 0x91, 0x23, 0x30, 0x72, 0x04, 0x26, 0xce, 0x91, 0x31, 0x63, 0x27, 0x20,
	0xce, 0xc2, 0xd8, 0x23, 0xa0, 0xd8, 0xce, 0xe4, 0x9f, 0xe7, 0x79, 0x3f, 0xbf, 0xb2, 0x3f, 0x59,
	0xa5, 0xaf, 0x5c, 0x52, 0x16, 0xb3, 0x5c, 0x71, 0x49, 0x63, 0x9e, 0xa5, 0x5c, 0xd2, 0x48, 0x64,
	0x8b, 0x34, 0x71, 0x0b, 0xc9, 0xa5, 0x50, 0x22, 0x38, 0x71, 0x0e, 0xb1, 0x0e, 0xb1, 0x70, 0x78,
	0x9c, 0x88, 0x44, 0x18, 0x83, 0x76, 0x3b, 0x2b, 0x0f, 0x21, 0x11, 0x22, 0x59, 0x72, 0x6a, 0x4e,
	0xf3, 0x72, 0x41, 0xe3, 0x52, 0x32, 0x95, 0x8a, 0xcc, 0xf1, 0x53, 0xc7, 0x65, 0x1e, 0xd1, 0x42,
	0x31, 0x55, 0x16, 0x16, 0x9c, 0x7f, 0x61, 0xdf, 0x7b, 0x60, 0x92, 0xad, 0x8a, 0xe0, 0xd2, 0xf7,
	0x2c, 0x1a, 0xe0, 0x31, 0x9e, 0xec, 0x5f, 0x05, 0xc4, 0x86, 0x88, 0xcc, 0x23, 0xf2, 0x68, 0x48,
	0xb8, 0x53, 0x7d, 0x8f, 0xd0, 0xd4, 0x79, 0xc1, 0xbd, 0x7f, 0xf8, 0xc2, 0x96, 0x69, 0x3c, 0xeb,
	0x5f, 0x1b, 0x6c, 0x99, 0xe4, 0x59, 0x9f, 0xec, 0xeb, 0x90, 0x3b, 0x27, 0x84, 0x7b, 0xdd, 0x80,
	0x8f, 0x9f, 0x11, 0x9e, 0x1e, 0x98, 0x68, 0x0f, 0x82, 0x0b, 0xff, 0xc8, 0xce, 0x2a, 0x0b, 0x3e,
	0x8b, 0x44, 0x99, 0xa9, 0xc1, 0xf6, 0x18, 0x4f, 0x76, 0x9d, 0xf7, 0x54, 0xf0, 0xdb, 0xee, 0x32,
	0xbc, 0xa9, 0x1a, 0x40, 0x75, 0x03, 0x68, 0xdd, 0x00, 0xda, 0x34, 0x80, 0xde, 0x34, 0xe0, 0x4f,
	0x0d, 0xa8, 0xd2, 0x80, 0x6b, 0x0d, 0xf8, 0x57, 0x03, 0xfe, 0xd3, 0x80, 0x36, 0x1a, 0xf0, 0x7b,
	0x0b, 0xa8, 0x6e, 0x01, 0xad, 0x5b, 0x40, 0xcf, 0x9e, 0xfd, 0xc1, 0xb9, 0x67, 0x5a, 0x5d, 0xff,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xc4, 0x85, 0xdf, 0x8b, 0x01, 0x00, 0x00,
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
	if m.ValidUseCount != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.ValidUseCount))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ValidDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintConfig(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
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
	l = m.Status.Size()
	n += 1 + l + sovConfig(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)
	n += 1 + l + sovConfig(uint64(l))
	if m.ValidUseCount != 0 {
		n += 1 + sovConfig(uint64(m.ValidUseCount))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`Status:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Status), "Status", "rpc.Status", 1), `&`, ``, 1) + `,`,
		`ValidDuration:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ValidDuration), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`ValidUseCount:` + fmt.Sprintf("%v", this.ValidUseCount) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.ValidDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUseCount", wireType)
			}
			m.ValidUseCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidUseCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)