package types_test

import (
	"testing"

	"github.com/be-heroes/doxchain/x/oauth2/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated DeviceCodeRegistry",
			genState: &types.GenesisState{
				DeviceCodeRegistryList: []types.DeviceCodeRegistry{
					{
						Tenant: "0",
					},
					{
						Tenant: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated AccessTokenRegistry",
			genState: &types.GenesisState{
				AccessTokenRegistryList: []types.AccessTokenRegistry{
					{
						Tenant: "0",
					},
					{
						Tenant: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}