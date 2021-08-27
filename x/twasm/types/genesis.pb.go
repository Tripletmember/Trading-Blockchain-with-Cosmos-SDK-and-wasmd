// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: confio/twasm/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/CosmWasm/wasmd/x/wasm/types"
	_ "github.com/cosmos/cosmos-sdk/types"
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
	Wasm                        types.GenesisState `protobuf:"bytes,1,opt,name=wasm,proto3" json:"wasm_genesis,omitempty"`
	PrivilegedContractAddresses []string           `protobuf:"bytes,2,rep,name=privileged_contract_addresses,json=privilegedContractAddresses,proto3" json:"privileged_contract_addresses,omitempty"`
	PinnedCodeIDs               []uint64           `protobuf:"varint,3,rep,packed,name=pinned_code_ids,json=pinnedCodeIds,proto3" json:"pinned_code_ids,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_89c4cd47eb0533ed, []int{0}
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

func (m *GenesisState) GetWasm() types.GenesisState {
	if m != nil {
		return m.Wasm
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetPrivilegedContractAddresses() []string {
	if m != nil {
		return m.PrivilegedContractAddresses
	}
	return nil
}

func (m *GenesisState) GetPinnedCodeIDs() []uint64 {
	if m != nil {
		return m.PinnedCodeIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "confio.twasm.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("confio/twasm/v1beta1/genesis.proto", fileDescriptor_89c4cd47eb0533ed)
}

var fileDescriptor_89c4cd47eb0533ed = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd1, 0x41, 0x4a, 0xf3, 0x40,
	0x14, 0x07, 0xf0, 0xa4, 0x2d, 0x1f, 0x7c, 0xd1, 0x22, 0x84, 0x22, 0xb5, 0xe2, 0xa4, 0xb4, 0xa0,
	0x05, 0x65, 0x42, 0xf5, 0x02, 0x9a, 0x0a, 0xe2, 0x4e, 0x14, 0x14, 0xdc, 0x84, 0x24, 0xf3, 0x8c,
	0x03, 0x26, 0x2f, 0x64, 0xc6, 0x6a, 0x6f, 0xe1, 0x51, 0x3c, 0x46, 0x97, 0x5d, 0xba, 0x0a, 0x92,
	0xee, 0x7a, 0x0a, 0xc9, 0x34, 0xb5, 0x41, 0xd0, 0x5d, 0xc8, 0xfb, 0xcd, 0xff, 0x3f, 0xcc, 0x33,
	0x7a, 0x01, 0xc6, 0x0f, 0x1c, 0x6d, 0xf9, 0xe2, 0x89, 0xc8, 0x1e, 0x0f, 0x7d, 0x90, 0xde, 0xd0,
	0x0e, 0x21, 0x06, 0xc1, 0x05, 0x4d, 0x52, 0x94, 0x68, 0xb6, 0x96, 0x86, 0x2a, 0x43, 0x4b, 0xd3,
	0x69, 0x85, 0x18, 0xa2, 0x02, 0x76, 0xf1, 0xb5, 0xb4, 0x1d, 0x12, 0xa0, 0x88, 0x50, 0xd8, 0xbe,
	0x27, 0xe0, 0x3b, 0x2e, 0x40, 0x1e, 0x97, 0xf3, 0x7e, 0x31, 0x57, 0x5d, 0xbf, 0x17, 0xf6, 0xde,
	0x6b, 0xc6, 0xe6, 0xc5, 0xf2, 0xcf, 0x8d, 0xf4, 0x24, 0x98, 0x77, 0x46, 0xa3, 0xe0, 0x6d, 0xbd,
	0xab, 0x0f, 0x36, 0x8e, 0xfb, 0x74, 0x15, 0x42, 0xab, 0x37, 0xa2, 0xd5, 0x23, 0x0e, 0x99, 0x66,
	0x96, 0xb6, 0xc8, 0xac, 0xed, 0x82, 0xb8, 0x65, 0xfe, 0x11, 0x46, 0x5c, 0x42, 0x94, 0xc8, 0xc9,
	0xb5, 0x0a, 0x34, 0xd1, 0xd8, 0x4b, 0x52, 0x3e, 0xe6, 0x4f, 0x10, 0x02, 0x73, 0x03, 0x8c, 0x65,
	0xea, 0x05, 0xd2, 0xf5, 0x18, 0x4b, 0x41, 0x08, 0x10, 0xed, 0x5a, 0xb7, 0x3e, 0xf8, 0xef, 0x1c,
	0x2e, 0x32, 0xeb, 0xe0, 0x4f, 0x58, 0x49, 0xde, 0x5d, 0xc3, 0x51, 0xe9, 0xce, 0x56, 0xcc, 0xbc,
	0x35, 0xb6, 0x12, 0x1e, 0xc7, 0x2a, 0x83, 0x81, 0xcb, 0x99, 0x68, 0xd7, 0xbb, 0xf5, 0x41, 0xc3,
	0xa1, 0x79, 0x66, 0x35, 0xaf, 0xd4, 0x68, 0x84, 0x0c, 0x2e, 0xcf, 0xc5, 0x22, 0xb3, 0x76, 0x7e,
	0xd8, 0x4a, 0x4b, 0x33, 0x59, 0x5b, 0x26, 0x9c, 0xd3, 0x69, 0x4e, 0xf4, 0x59, 0x4e, 0xf4, 0xcf,
	0x9c, 0xe8, 0x6f, 0x73, 0xa2, 0xcd, 0xe6, 0x44, 0xfb, 0x98, 0x13, 0xed, 0x7e, 0x3f, 0xe4, 0xf2,
	0xf1, 0xd9, 0xa7, 0x01, 0x46, 0xf6, 0x6a, 0xd9, 0x61, 0xea, 0x31, 0xb0, 0x5f, 0xcb, 0xad, 0xcb,
	0x49, 0x02, 0xc2, 0xff, 0xa7, 0xde, 0xfe, 0xe4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xee, 0xd5, 0xcb,
	0x4d, 0x12, 0x02, 0x00, 0x00,
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
	if len(m.PinnedCodeIDs) > 0 {
		dAtA2 := make([]byte, len(m.PinnedCodeIDs)*10)
		var j1 int
		for _, num := range m.PinnedCodeIDs {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGenesis(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PrivilegedContractAddresses) > 0 {
		for iNdEx := len(m.PrivilegedContractAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PrivilegedContractAddresses[iNdEx])
			copy(dAtA[i:], m.PrivilegedContractAddresses[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.PrivilegedContractAddresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Wasm.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.Wasm.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.PrivilegedContractAddresses) > 0 {
		for _, s := range m.PrivilegedContractAddresses {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PinnedCodeIDs) > 0 {
		l = 0
		for _, e := range m.PinnedCodeIDs {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
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
				return fmt.Errorf("proto: wrong wireType = %d for field Wasm", wireType)
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
			if err := m.Wasm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivilegedContractAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrivilegedContractAddresses = append(m.PrivilegedContractAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PinnedCodeIDs = append(m.PinnedCodeIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.PinnedCodeIDs) == 0 {
					m.PinnedCodeIDs = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PinnedCodeIDs = append(m.PinnedCodeIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PinnedCodeIDs", wireType)
			}
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
