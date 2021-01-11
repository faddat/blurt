package blurt

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/faddat/blurt/x/blurt/types"
	"github.com/faddat/blurt/x/blurt/keeper"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePost) (*sdk.Result, error) {
	k.CreatePost(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdatePost) (*sdk.Result, error) {
	var post = types.Post{
		Creator: msg.Creator,
		Id:      msg.Id,
    	Title: msg.Title,
    	Body: msg.Body,
    	Ipfs: msg.Ipfs,
    	Parent: msg.Parent,
    	Category: msg.Category,
	}

    if msg.Creator != k.GetPostOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error                                                                                             
    }          

	k.UpdatePost(ctx, post)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeletePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeletePost) (*sdk.Result, error) {
    if !k.HasPost(ctx, msg.Id) {                                                                                                                                                                    
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)                                                                                                                                
    }                                                                                                                                                                                                  
    if msg.Creator != k.GetPostOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")                                                                                                                       
    } 

	k.DeletePost(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
