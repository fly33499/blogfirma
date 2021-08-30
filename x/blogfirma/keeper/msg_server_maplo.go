package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fly33499/blogfirma/x/blogfirma/types"
)

func (k msgServer) CreateMaplo(goCtx context.Context, msg *types.MsgCreateMaplo) (*types.MsgCreateMaploResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetMaplo(ctx, msg.Index)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("index %v already set", msg.Index))
	}

	var maplo = types.Maplo{
		Index:   msg.Index,
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	k.SetMaplo(
		ctx,
		maplo,
	)
	return &types.MsgCreateMaploResponse{}, nil
}

func (k msgServer) UpdateMaplo(goCtx context.Context, msg *types.MsgUpdateMaplo) (*types.MsgUpdateMaploResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMaplo(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var maplo = types.Maplo{
		Index:   msg.Index,
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	k.SetMaplo(ctx, maplo)

	return &types.MsgUpdateMaploResponse{}, nil
}

func (k msgServer) DeleteMaplo(goCtx context.Context, msg *types.MsgDeleteMaplo) (*types.MsgDeleteMaploResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMaplo(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("index %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMaplo(ctx, msg.Index)

	return &types.MsgDeleteMaploResponse{}, nil
}
