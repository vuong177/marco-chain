// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: marco/marco/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgMintStableCoin struct {
	// minter
	Minter string `protobuf:"bytes,1,opt,name=minter,proto3" json:"minter,omitempty" yaml:"minter"`
	// collateral asset
	CollateralAsset github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=collateral_asset,json=collateralAsset,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"collateral_asset" yaml:"staking"`
}

func (m *MsgMintStableCoin) Reset()         { *m = MsgMintStableCoin{} }
func (m *MsgMintStableCoin) String() string { return proto.CompactTextString(m) }
func (*MsgMintStableCoin) ProtoMessage()    {}
func (*MsgMintStableCoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc38357a5b10ce12, []int{0}
}
func (m *MsgMintStableCoin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintStableCoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintStableCoin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintStableCoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintStableCoin.Merge(m, src)
}
func (m *MsgMintStableCoin) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintStableCoin) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintStableCoin.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintStableCoin proto.InternalMessageInfo

func (m *MsgMintStableCoin) GetMinter() string {
	if m != nil {
		return m.Minter
	}
	return ""
}

type MsgMintStableCoinResponse struct {
}

func (m *MsgMintStableCoinResponse) Reset()         { *m = MsgMintStableCoinResponse{} }
func (m *MsgMintStableCoinResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMintStableCoinResponse) ProtoMessage()    {}
func (*MsgMintStableCoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc38357a5b10ce12, []int{1}
}
func (m *MsgMintStableCoinResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintStableCoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintStableCoinResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintStableCoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintStableCoinResponse.Merge(m, src)
}
func (m *MsgMintStableCoinResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintStableCoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintStableCoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintStableCoinResponse proto.InternalMessageInfo

type MsgWithdrawCollateral struct {
	// minter
	Minter string `protobuf:"bytes,1,opt,name=minter,proto3" json:"minter,omitempty" yaml:"minter"`
	// collateral asset
	CollateralAsset github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=collateral_asset,json=collateralAsset,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"collateral_asset" yaml:"staking"`
}

func (m *MsgWithdrawCollateral) Reset()         { *m = MsgWithdrawCollateral{} }
func (m *MsgWithdrawCollateral) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawCollateral) ProtoMessage()    {}
func (*MsgWithdrawCollateral) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc38357a5b10ce12, []int{2}
}
func (m *MsgWithdrawCollateral) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawCollateral) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawCollateral.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawCollateral) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawCollateral.Merge(m, src)
}
func (m *MsgWithdrawCollateral) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawCollateral) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawCollateral.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawCollateral proto.InternalMessageInfo

func (m *MsgWithdrawCollateral) GetMinter() string {
	if m != nil {
		return m.Minter
	}
	return ""
}

type MsgWithdrawCollateralResponse struct {
}

func (m *MsgWithdrawCollateralResponse) Reset()         { *m = MsgWithdrawCollateralResponse{} }
func (m *MsgWithdrawCollateralResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawCollateralResponse) ProtoMessage()    {}
func (*MsgWithdrawCollateralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc38357a5b10ce12, []int{3}
}
func (m *MsgWithdrawCollateralResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawCollateralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawCollateralResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawCollateralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawCollateralResponse.Merge(m, src)
}
func (m *MsgWithdrawCollateralResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawCollateralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawCollateralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawCollateralResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgMintStableCoin)(nil), "marco.marco.MsgMintStableCoin")
	proto.RegisterType((*MsgMintStableCoinResponse)(nil), "marco.marco.MsgMintStableCoinResponse")
	proto.RegisterType((*MsgWithdrawCollateral)(nil), "marco.marco.MsgWithdrawCollateral")
	proto.RegisterType((*MsgWithdrawCollateralResponse)(nil), "marco.marco.MsgWithdrawCollateralResponse")
}

func init() { proto.RegisterFile("marco/marco/tx.proto", fileDescriptor_fc38357a5b10ce12) }

var fileDescriptor_fc38357a5b10ce12 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4d, 0x2c, 0x4a,
	0xce, 0xd7, 0x87, 0x90, 0x25, 0x15, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xdc, 0x60, 0xbe,
	0x1e, 0x98, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x8b, 0xeb, 0x83, 0x58, 0x10, 0x25, 0x4a,
	0x6b, 0x18, 0xb9, 0x04, 0x7d, 0x8b, 0xd3, 0x7d, 0x33, 0xf3, 0x4a, 0x82, 0x4b, 0x12, 0x93, 0x72,
	0x52, 0x9d, 0xf3, 0x33, 0xf3, 0x84, 0x34, 0xb9, 0xd8, 0x72, 0x33, 0xf3, 0x4a, 0x52, 0x8b, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x04, 0x3f, 0xdd, 0x93, 0xe7, 0xad, 0x4c, 0xcc, 0xcd, 0xb1,
	0x52, 0x82, 0x88, 0x2b, 0x05, 0x41, 0x15, 0x08, 0xe5, 0x70, 0x09, 0x24, 0xe7, 0xe7, 0xe4, 0x24,
	0x96, 0xa4, 0x16, 0x25, 0xe6, 0xc4, 0x27, 0x16, 0x17, 0xa7, 0x96, 0x48, 0x30, 0x81, 0x35, 0x39,
	0x9e, 0xb8, 0x27, 0xcf, 0x70, 0xeb, 0x9e, 0xbc, 0x7a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x7e, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0x31, 0x94, 0xd2, 0x2d, 0x4e, 0xc9, 0xd6,
	0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x03, 0xd9, 0xfa, 0xe9, 0x9e, 0x3c, 0x1f, 0xc4, 0x8e, 0xe2,
	0x92, 0xc4, 0xec, 0xcc, 0xbc, 0x74, 0xa5, 0x20, 0x7e, 0x84, 0xd1, 0x8e, 0x20, 0x93, 0x95, 0xa4,
	0xb9, 0x24, 0x31, 0x5c, 0x1b, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xaa, 0xb4, 0x81, 0x91,
	0x4b, 0xd4, 0xb7, 0x38, 0x3d, 0x3c, 0xb3, 0x24, 0x23, 0xa5, 0x28, 0xb1, 0xdc, 0x19, 0xae, 0x77,
	0xf0, 0xfa, 0x47, 0x9e, 0x4b, 0x16, 0xab, 0x8b, 0x61, 0x7e, 0x32, 0x3a, 0xca, 0xc8, 0xc5, 0xec,
	0x5b, 0x9c, 0x2e, 0x14, 0xc1, 0xc5, 0x87, 0x16, 0x47, 0x72, 0x7a, 0x48, 0xb1, 0xab, 0x87, 0x11,
	0x2a, 0x52, 0x6a, 0xf8, 0xe5, 0x61, 0x36, 0x08, 0xa5, 0x70, 0x09, 0x61, 0x09, 0x31, 0x25, 0x74,
	0xdd, 0x98, 0x6a, 0xa4, 0xb4, 0x08, 0xab, 0x81, 0xd9, 0xe2, 0xa4, 0x7e, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c,
	0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xbc, 0x15, 0xb0, 0x64, 0x0b, 0x0a, 0xb9, 0x24, 0x36, 0x70,
	0xba, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x33, 0xbc, 0x71, 0xd2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	MintStableCoin(ctx context.Context, in *MsgMintStableCoin, opts ...grpc.CallOption) (*MsgMintStableCoinResponse, error)
	WithdrawCollateral(ctx context.Context, in *MsgWithdrawCollateral, opts ...grpc.CallOption) (*MsgWithdrawCollateralResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MintStableCoin(ctx context.Context, in *MsgMintStableCoin, opts ...grpc.CallOption) (*MsgMintStableCoinResponse, error) {
	out := new(MsgMintStableCoinResponse)
	err := c.cc.Invoke(ctx, "/marco.marco.Msg/MintStableCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawCollateral(ctx context.Context, in *MsgWithdrawCollateral, opts ...grpc.CallOption) (*MsgWithdrawCollateralResponse, error) {
	out := new(MsgWithdrawCollateralResponse)
	err := c.cc.Invoke(ctx, "/marco.marco.Msg/WithdrawCollateral", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	MintStableCoin(context.Context, *MsgMintStableCoin) (*MsgMintStableCoinResponse, error)
	WithdrawCollateral(context.Context, *MsgWithdrawCollateral) (*MsgWithdrawCollateralResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) MintStableCoin(ctx context.Context, req *MsgMintStableCoin) (*MsgMintStableCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintStableCoin not implemented")
}
func (*UnimplementedMsgServer) WithdrawCollateral(ctx context.Context, req *MsgWithdrawCollateral) (*MsgWithdrawCollateralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawCollateral not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_MintStableCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintStableCoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintStableCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marco.marco.Msg/MintStableCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintStableCoin(ctx, req.(*MsgMintStableCoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawCollateral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawCollateral)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawCollateral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marco.marco.Msg/WithdrawCollateral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawCollateral(ctx, req.(*MsgWithdrawCollateral))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "marco.marco.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintStableCoin",
			Handler:    _Msg_MintStableCoin_Handler,
		},
		{
			MethodName: "WithdrawCollateral",
			Handler:    _Msg_WithdrawCollateral_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marco/marco/tx.proto",
}

func (m *MsgMintStableCoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintStableCoin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintStableCoin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CollateralAsset.Size()
		i -= size
		if _, err := m.CollateralAsset.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Minter) > 0 {
		i -= len(m.Minter)
		copy(dAtA[i:], m.Minter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Minter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintStableCoinResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintStableCoinResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintStableCoinResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawCollateral) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawCollateral) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawCollateral) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CollateralAsset.Size()
		i -= size
		if _, err := m.CollateralAsset.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Minter) > 0 {
		i -= len(m.Minter)
		copy(dAtA[i:], m.Minter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Minter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawCollateralResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawCollateralResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawCollateralResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgMintStableCoin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Minter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.CollateralAsset.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgMintStableCoinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdrawCollateral) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Minter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.CollateralAsset.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgWithdrawCollateralResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgMintStableCoin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintStableCoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintStableCoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Minter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralAsset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgMintStableCoinResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintStableCoinResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintStableCoinResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawCollateral) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawCollateral: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawCollateral: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Minter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralAsset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawCollateralResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawCollateralResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawCollateralResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
