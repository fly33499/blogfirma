package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fly33499/blogfirma/x/blogfirma/types"
	"github.com/stretchr/testify/assert"
)

func createNHello(keeper *Keeper, ctx sdk.Context, n int) []types.Hello {
	items := make([]types.Hello, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendHello(ctx, items[i])
	}
	return items
}

func TestHelloGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNHello(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetHello(ctx, item.Id))
	}
}

func TestHelloExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNHello(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasHello(ctx, item.Id))
	}
}

func TestHelloRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNHello(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHello(ctx, item.Id)
		assert.False(t, keeper.HasHello(ctx, item.Id))
	}
}

func TestHelloGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNHello(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllHello(ctx))
}

func TestHelloCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNHello(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetHelloCount(ctx))
}
