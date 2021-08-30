package blogfirma

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fly33499/blogfirma/x/blogfirma/keeper"
	"github.com/fly33499/blogfirma/x/blogfirma/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the maplo
	for _, elem := range genState.MaploList {
		k.SetMaplo(ctx, *elem)
	}

	// Set all the hello
	for _, elem := range genState.HelloList {
		k.SetHello(ctx, *elem)
	}

	// Set hello count
	k.SetHelloCount(ctx, genState.HelloCount)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all maplo
	maploList := k.GetAllMaplo(ctx)
	for _, elem := range maploList {
		elem := elem
		genesis.MaploList = append(genesis.MaploList, &elem)
	}

	// Get all hello
	helloList := k.GetAllHello(ctx)
	for _, elem := range helloList {
		elem := elem
		genesis.HelloList = append(genesis.HelloList, &elem)
	}

	// Set the current count
	genesis.HelloCount = k.GetHelloCount(ctx)

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
