package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/fly33499/blogfirma/x/blogfirma/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MaploAll(c context.Context, req *types.QueryAllMaploRequest) (*types.QueryAllMaploResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var maplos []*types.Maplo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	maploStore := prefix.NewStore(store, types.KeyPrefix(types.MaploKey))

	pageRes, err := query.Paginate(maploStore, req.Pagination, func(key []byte, value []byte) error {
		var maplo types.Maplo
		if err := k.cdc.UnmarshalBinaryBare(value, &maplo); err != nil {
			return err
		}

		maplos = append(maplos, &maplo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMaploResponse{Maplo: maplos, Pagination: pageRes}, nil
}

func (k Keeper) Maplo(c context.Context, req *types.QueryGetMaploRequest) (*types.QueryGetMaploResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMaplo(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetMaploResponse{Maplo: &val}, nil
}
