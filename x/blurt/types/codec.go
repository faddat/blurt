package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(&MsgCreatePost{}, "blurt/CreatePost", nil)
cdc.RegisterConcrete(&MsgUpdatePost{}, "blurt/UpdatePost", nil)
cdc.RegisterConcrete(&MsgDeletePost{}, "blurt/DeletePost", nil)

} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreatePost{},
	&MsgUpdatePost{},
	&MsgDeletePost{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
