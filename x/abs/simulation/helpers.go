package simulation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func FindAccount(accs []simtypes.Account, address string) (simtypes.Account, bool) {
	creator, err := sdk.AccAddressFromBech32(address)

	if err != nil {
		panic(err)
	}
	
	return simtypes.FindAccount(accs, creator)
}
