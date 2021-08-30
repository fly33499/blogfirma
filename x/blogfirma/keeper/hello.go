package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fly33499/blogfirma/x/blogfirma/types"
	"strconv"
)

// GetHelloCount get the total number of TypeName.LowerCamel
func (k Keeper) GetHelloCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelloCountKey))
	byteKey := types.KeyPrefix(types.HelloCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to uint64
		panic("cannot decode count")
	}

	return count
}

// SetHelloCount set the total number of hello
func (k Keeper) SetHelloCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelloCountKey))
	byteKey := types.KeyPrefix(types.HelloCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendHello appends a hello in the store with a new id and update the count
func (k Keeper) AppendHello(
	ctx sdk.Context,
	hello types.Hello,
) uint64 {
	// Create the hello
	count := k.GetHelloCount(ctx)

	// Set the ID of the appended value
	hello.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelloKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&hello)
	store.Set(GetHelloIDBytes(hello.Id), appendedValue)

	// Update hello count
	k.SetHelloCount(ctx, count+1)

	return count
}

// SetHello set a specific hello in the store
func (k Keeper) SetHello(ctx sdk.Context, hello types.Hello) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelloKey))
	b := k.cdc.MustMarshalBinaryBare(&hello)
	store.Set(GetHelloIDBytes(hello.Id), b)
}

// GetHello returns a hello from its id
func (k Keeper) GetHello(ctx sdk.Context, id uint64) types.Hello {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelloKey))
	var hello types.Hello
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetHelloIDBytes(id)), &hello)
	return hello
}

// HasHello checks if the hello exists in the store
func (k Keeper) HasHello(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelloKey))
	return store.Has(GetHelloIDBytes(id))
}

// GetHelloOwner returns the creator of the
func (k Keeper) GetHelloOwner(ctx sdk.Context, id uint64) string {
	return k.GetHello(ctx, id).Creator
}

// RemoveHello removes a hello from the store
func (k Keeper) RemoveHello(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelloKey))
	store.Delete(GetHelloIDBytes(id))
}

// GetAllHello returns all hello
func (k Keeper) GetAllHello(ctx sdk.Context) (list []types.Hello) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HelloKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Hello
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHelloIDBytes returns the byte representation of the ID
func GetHelloIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetHelloIDFromBytes returns ID in uint64 format from a byte array
func GetHelloIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
