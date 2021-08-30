package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/fly33499/blogfirma/x/blogfirma/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) HelloAll(c context.Context, req *types.QueryAllHelloRequest) (*types.QueryAllHelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hellos []*types.Hello
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	helloStore := prefix.NewStore(store, types.KeyPrefix(types.HelloKey))

	pageRes, err := query.Paginate(helloStore, req.Pagination, func(key []byte, value []byte) error {
		var hello types.Hello
		if err := k.cdc.UnmarshalBinaryBare(value, &hello); err != nil {
			return err
		}

		hellos = append(hellos, &hello)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHelloResponse{Hello: hellos, Pagination: pageRes}, nil
}

func (k Keeper) Hello(c context.Context, req *types.QueryGetHelloRequest) (*types.QueryGetHelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hello types.Hello
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasHello(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelloKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetHelloIDBytes(req.Id)), &hello)

	return &types.QueryGetHelloResponse{Hello: &hello}, nil
}
