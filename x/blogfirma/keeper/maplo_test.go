package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/fly33499/blogfirma/x/blogfirma/types"
)

func createNMaplo(keeper *Keeper, ctx sdk.Context, n int) []types.Maplo {
	items := make([]types.Maplo, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetMaplo(ctx, items[i])
	}
	return items
}

func TestMaploGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMaplo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMaplo(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestMaploRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMaplo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMaplo(ctx, item.Index)
		_, found := keeper.GetMaplo(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestMaploGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMaplo(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllMaplo(ctx))
}
