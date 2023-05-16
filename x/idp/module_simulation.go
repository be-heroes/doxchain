package idp

import (
	"math/rand"

	"github.com/be-heroes/doxchain/testutil/sample"
	idpsimulation "github.com/be-heroes/doxchain/x/idp/simulation"
	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = idpsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAuthenticationRequest = "op_weight_msg_basic_authentication_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAuthenticationRequest int = 100

	opWeightMsgCreateClientRegistry = "op_weight_msg_client_registrations"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateClientRegistry int = 100

	opWeightMsgUpdateClientRegistry = "op_weight_msg_client_registrations"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateClientRegistry int = 100

	opWeightMsgDeleteClientRegistry = "op_weight_msg_client_registrations"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteClientRegistry int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	idpGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ClientRegistryList: []types.ClientRegistry{
			{
				Creator: sample.AccAddress(),
			},
			{
				Creator: sample.AccAddress(),
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&idpGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAuthenticationRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAuthenticationRequest, &weightMsgAuthenticationRequest, nil,
		func(_ *rand.Rand) {
			weightMsgAuthenticationRequest = defaultWeightMsgAuthenticationRequest
		},
	)

	var weightMsgCreateClientRegistry int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateClientRegistry, &weightMsgCreateClientRegistry, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClientRegistry = defaultWeightMsgCreateClientRegistry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateClientRegistry,
		idpsimulation.SimulateMsgCreateClientRegistry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateClientRegistry int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateClientRegistry, &weightMsgUpdateClientRegistry, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClientRegistry = defaultWeightMsgUpdateClientRegistry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateClientRegistry,
		idpsimulation.SimulateMsgUpdateClientRegistry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteClientRegistry int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteClientRegistry, &weightMsgDeleteClientRegistry, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteClientRegistry = defaultWeightMsgDeleteClientRegistry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteClientRegistry,
		idpsimulation.SimulateMsgDeleteClientRegistry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
