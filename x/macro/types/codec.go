package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgMintStableCoin{}, "macro/MsgMintStableCoin", nil)
	cdc.RegisterConcrete(&MsgWithdrawCollateral{}, "macro/MsgWithdrawCollateral", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "macro/MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgRepay{}, "macro/MsgRepay", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgMintStableCoin{},
		&MsgWithdrawCollateral{},
		&MsgDeposit{},
		&MsgRepay{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}