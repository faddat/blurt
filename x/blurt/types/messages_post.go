package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePost{}

func NewMsgCreatePost(creator string, title string, body string, ipfs string, category string) *MsgCreatePost {
  return &MsgCreatePost{
		Creator: creator,
    Title: title,
    Body: body,
    Ipfs: ipfs,
    Category: category,
	}
}

func (msg *MsgCreatePost) Route() string {
  return RouterKey
}

func (msg *MsgCreatePost) Type() string {
  return "CreatePost"
}

func (msg *MsgCreatePost) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePost) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePost) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdatePost{}

func NewMsgUpdatePost(creator string, id string, title string, body string, ipfs string, category string) *MsgUpdatePost {
  return &MsgUpdatePost{
        Id: id,
		Creator: creator,
    Title: title,
    Body: body,
    Ipfs: ipfs,
    Category: category,
	}
}

func (msg *MsgUpdatePost) Route() string {
  return RouterKey
}

func (msg *MsgUpdatePost) Type() string {
  return "UpdatePost"
}

func (msg *MsgUpdatePost) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePost) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePost) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgCreatePost{}

func NewMsgDeletePost(creator string, id string) *MsgDeletePost {
  return &MsgDeletePost{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeletePost) Route() string {
  return RouterKey
}

func (msg *MsgDeletePost) Type() string {
  return "DeletePost"
}

func (msg *MsgDeletePost) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePost) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePost) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}
