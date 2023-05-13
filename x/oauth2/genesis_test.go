package oauth_test

import (
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/oauth2"
	"github.com/be-heroes/doxchain/x/oauth2/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DeviceCodeRegistryList: []types.DeviceCodeRegistry{
			{
				Tenant: "0",
			},
			{
				Tenant: "1",
			},
		},
		AccessTokenRegistryList: []types.AccessTokenRegistry{
			{
				Tenant: "0",
			},
			{
				Tenant: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OauthKeeper(t)
	oauth.InitGenesis(ctx, *k, genesisState)
	got := oauth.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DeviceCodeRegistryList, got.DeviceCodeRegistryList)
	require.ElementsMatch(t, genesisState.AccessTokenRegistryList, got.AccessTokenRegistryList)
	// this line is used by starport scaffolding # genesis/test/assert
}