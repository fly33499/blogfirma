package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fly33499/blogfirma/x/blogfirma/types"
)

func (k msgServer) CreateHello(goCtx context.Context, msg *types.MsgCreateHello) (*types.MsgCreateHelloResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hello = types.Hello{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	id := k.AppendHello(
		ctx,
		hello,
	)

	return &types.MsgCreateHelloResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateHello(goCtx context.Context, msg *types.MsgUpdateHello) (*types.MsgUpdateHelloResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var hello = types.Hello{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	// Checks that the element exists
	if !k.HasHello(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetHelloOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetHello(ctx, hello)

	return &types.MsgUpdateHelloResponse{}, nil
}

func (k msgServer) DeleteHello(goCtx context.Context, msg *types.MsgDeleteHello) (*types.MsgDeleteHelloResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasHello(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetHelloOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveHello(ctx, msg.Id)

	return &types.MsgDeleteHelloResponse{}, nil
}
