package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		MaploList: []*Maplo{},
		HelloList: []*Hello{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in maplo
	maploIndexMap := make(map[string]bool)

	for _, elem := range gs.MaploList {
		if _, ok := maploIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for maplo")
		}
		maploIndexMap[elem.Index] = true
	}
	// Check for duplicated ID in hello
	helloIdMap := make(map[uint64]bool)

	for _, elem := range gs.HelloList {
		if _, ok := helloIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for hello")
		}
		helloIdMap[elem.Id] = true
	}

	return nil
}
