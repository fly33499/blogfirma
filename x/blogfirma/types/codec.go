package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateMaplo{}, "blogfirma/CreateMaplo", nil)
	cdc.RegisterConcrete(&MsgUpdateMaplo{}, "blogfirma/UpdateMaplo", nil)
	cdc.RegisterConcrete(&MsgDeleteMaplo{}, "blogfirma/DeleteMaplo", nil)

	cdc.RegisterConcrete(&MsgCreateHello{}, "blogfirma/CreateHello", nil)
	cdc.RegisterConcrete(&MsgUpdateHello{}, "blogfirma/UpdateHello", nil)
	cdc.RegisterConcrete(&MsgDeleteHello{}, "blogfirma/DeleteHello", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMaplo{},
		&MsgUpdateMaplo{},
		&MsgDeleteMaplo{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHello{},
		&MsgUpdateHello{},
		&MsgDeleteHello{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
