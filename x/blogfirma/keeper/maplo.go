package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fly33499/blogfirma/x/blogfirma/types"
)

// SetMaplo set a specific maplo in the store from its index
func (k Keeper) SetMaplo(ctx sdk.Context, maplo types.Maplo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MaploKey))
	b := k.cdc.MustMarshalBinaryBare(&maplo)
	store.Set(types.KeyPrefix(maplo.Index), b)
}

// GetMaplo returns a maplo from its index
func (k Keeper) GetMaplo(ctx sdk.Context, index string) (val types.Maplo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MaploKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveMaplo removes a maplo from the store
func (k Keeper) RemoveMaplo(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MaploKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllMaplo returns all maplo
func (k Keeper) GetAllMaplo(ctx sdk.Context) (list []types.Maplo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MaploKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Maplo
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
