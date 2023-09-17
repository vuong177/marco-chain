// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: macro/macro/asset.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// UserData defines the macro module's UserData state.
type BorrowerData struct {
	Address              string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CollateralAsset      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=collateral_asset,json=collateralAsset,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"collateral_asset" yaml:"collateral_asset"`
	Borrowed             github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,3,opt,name=borrowed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"borrowed" yaml:"minted_stable_coin"`
	IsRedemptionProvider bool                                     `protobuf:"varint,4,opt,name=is_redemption_provider,json=isRedemptionProvider,proto3" json:"is_redemption_provider,omitempty"`
}

func (m *BorrowerData) Reset()         { *m = BorrowerData{} }
func (m *BorrowerData) String() string { return proto.CompactTextString(m) }
func (*BorrowerData) ProtoMessage()    {}
func (*BorrowerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e4d3a1203883174, []int{0}
}
func (m *BorrowerData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BorrowerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BorrowerData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BorrowerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BorrowerData.Merge(m, src)
}
func (m *BorrowerData) XXX_Size() int {
	return m.Size()
}
func (m *BorrowerData) XXX_DiscardUnknown() {
	xxx_messageInfo_BorrowerData.DiscardUnknown(m)
}

var xxx_messageInfo_BorrowerData proto.InternalMessageInfo

func (m *BorrowerData) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BorrowerData) GetCollateralAsset() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CollateralAsset
	}
	return nil
}

func (m *BorrowerData) GetIsRedemptionProvider() bool {
	if m != nil {
		return m.IsRedemptionProvider
	}
	return false
}

func init() {
	proto.RegisterType((*BorrowerData)(nil), "macro.macro.BorrowerData")
}

func init() { proto.RegisterFile("macro/macro/asset.proto", fileDescriptor_6e4d3a1203883174) }

var fileDescriptor_6e4d3a1203883174 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0x4e, 0xe3, 0x40,
	0x10, 0xb5, 0x93, 0xd3, 0x5d, 0xce, 0x39, 0xe9, 0x4e, 0x56, 0x74, 0x71, 0x52, 0xd8, 0x91, 0x0b,
	0x70, 0x13, 0x5b, 0x09, 0x48, 0x91, 0xe8, 0x62, 0xd2, 0xd1, 0x20, 0xd3, 0xd1, 0x58, 0x6b, 0x7b,
	0x65, 0x56, 0xd8, 0xde, 0x68, 0x77, 0x13, 0xc8, 0x5f, 0xc0, 0x6f, 0x50, 0xf3, 0x11, 0x29, 0x23,
	0x2a, 0x44, 0x61, 0x50, 0xf2, 0x07, 0x91, 0xe8, 0x91, 0x77, 0x97, 0x80, 0xa8, 0x68, 0x66, 0xec,
	0x79, 0xf3, 0xde, 0xbe, 0x99, 0x5d, 0xad, 0x9d, 0x83, 0x98, 0x60, 0x4f, 0x44, 0x40, 0x29, 0x64,
	0xee, 0x94, 0x60, 0x86, 0xf5, 0x26, 0x2f, 0xb9, 0x3c, 0x76, 0x5b, 0x29, 0x4e, 0x31, 0xaf, 0x7b,
	0xd5, 0x97, 0x68, 0xe9, 0x9a, 0x31, 0xa6, 0x39, 0xa6, 0x5e, 0x04, 0x28, 0xf4, 0xe6, 0x83, 0x08,
	0x32, 0x30, 0xf0, 0x62, 0x8c, 0x0a, 0x89, 0x77, 0x04, 0x1e, 0x0a, 0xa2, 0xf8, 0x11, 0x90, 0xfd,
	0x5a, 0xd3, 0xfe, 0xf8, 0x98, 0x10, 0x7c, 0x05, 0xc9, 0x04, 0x30, 0xa0, 0x0f, 0xb5, 0x5f, 0x20,
	0x49, 0x08, 0xa4, 0xd4, 0x50, 0x7b, 0xaa, 0xf3, 0xdb, 0x37, 0x1e, 0xee, 0xfb, 0x2d, 0xc9, 0x19,
	0x0b, 0xe4, 0x8c, 0x11, 0x54, 0xa4, 0xc1, 0x7b, 0xa3, 0x7e, 0xab, 0x6a, 0xff, 0x62, 0x9c, 0x65,
	0x80, 0x41, 0x02, 0xb2, 0x90, 0xbb, 0x37, 0x6a, 0xbd, 0xba, 0xd3, 0x1c, 0x76, 0x5c, 0x49, 0xad,
	0xbc, 0xb9, 0xd2, 0x9b, 0x7b, 0x8c, 0x51, 0xe1, 0x9f, 0x2c, 0x4b, 0x4b, 0xd9, 0x96, 0x56, 0x7b,
	0x01, 0xf2, 0xec, 0xc8, 0xfe, 0x2a, 0x60, 0xdf, 0x3d, 0x5b, 0x4e, 0x8a, 0xd8, 0xc5, 0x2c, 0x72,
	0x63, 0x9c, 0x4b, 0xdb, 0x32, 0xf5, 0x69, 0x72, 0xe9, 0xb1, 0xc5, 0x14, 0x52, 0xae, 0x45, 0x83,
	0xbf, 0x1f, 0xf4, 0x71, 0xc5, 0xd6, 0x53, 0xad, 0x11, 0x89, 0xb9, 0x12, 0xa3, 0xce, 0x07, 0xe1,
	0xe7, 0x3d, 0x95, 0xd6, 0xde, 0x37, 0x44, 0x27, 0x30, 0xde, 0x96, 0x56, 0x47, 0x38, 0xcb, 0x51,
	0xc1, 0x60, 0x12, 0x52, 0x06, 0xa2, 0x0c, 0x86, 0xd5, 0x5e, 0xed, 0x60, 0x27, 0xae, 0x1f, 0x6a,
	0xff, 0x11, 0x0d, 0x09, 0x4c, 0x60, 0x3e, 0x65, 0x08, 0x17, 0xd5, 0x96, 0xe7, 0x28, 0x81, 0xc4,
	0xf8, 0xd1, 0x53, 0x9d, 0x46, 0xd0, 0x42, 0x34, 0xd8, 0x81, 0xa7, 0x12, 0xf3, 0xc7, 0xcb, 0xb5,
	0xa9, 0xae, 0xd6, 0xa6, 0xfa, 0xb2, 0x36, 0xd5, 0x9b, 0x8d, 0xa9, 0xac, 0x36, 0xa6, 0xf2, 0xb8,
	0x31, 0x95, 0xf3, 0xfd, 0x4f, 0xf6, 0xe6, 0x33, 0x5c, 0xa4, 0x83, 0xd1, 0x48, 0x3e, 0x8b, 0x6b,
	0x99, 0xb9, 0xc7, 0xe8, 0x27, 0xbf, 0xc1, 0x83, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xe3,
	0x02, 0x80, 0x3a, 0x02, 0x00, 0x00,
}

func (m *BorrowerData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BorrowerData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BorrowerData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsRedemptionProvider {
		i--
		if m.IsRedemptionProvider {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.Borrowed.Size()
		i -= size
		if _, err := m.Borrowed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAsset(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.CollateralAsset) > 0 {
		for iNdEx := len(m.CollateralAsset) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollateralAsset[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAsset(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAsset(dAtA []byte, offset int, v uint64) int {
	offset -= sovAsset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BorrowerData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if len(m.CollateralAsset) > 0 {
		for _, e := range m.CollateralAsset {
			l = e.Size()
			n += 1 + l + sovAsset(uint64(l))
		}
	}
	l = m.Borrowed.Size()
	n += 1 + l + sovAsset(uint64(l))
	if m.IsRedemptionProvider {
		n += 2
	}
	return n
}

func sovAsset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAsset(x uint64) (n int) {
	return sovAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BorrowerData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsset
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
			return fmt.Errorf("proto: BorrowerData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BorrowerData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAsset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralAsset = append(m.CollateralAsset, types.Coin{})
			if err := m.CollateralAsset[len(m.CollateralAsset)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Borrowed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Borrowed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRedemptionProvider", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
			m.IsRedemptionProvider = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAsset
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
func skipAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
				return 0, ErrInvalidLengthAsset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAsset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAsset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAsset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAsset = fmt.Errorf("proto: unexpected end of group")
)
