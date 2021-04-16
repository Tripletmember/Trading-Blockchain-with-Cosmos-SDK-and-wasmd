package keeper

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/confio/tgrade/x/twasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis sets supply information for genesis.
//
// CONTRACT: all types of accounts must have been already initialized/created
func InitGenesis(ctx sdk.Context, keeper *Keeper, data types.GenesisState, stakingKeeper wasmkeeper.ValidatorSetSource, msgHandler sdk.Handler) ([]abci.ValidatorUpdate, error) {
	result, err := wasmkeeper.InitGenesis(ctx, &keeper.Keeper, data.Wasm, stakingKeeper, msgHandler)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "wasm")
	}

	// import callbacks from dumped contract infos
	importedCallbackContracts := make(map[string]struct{})
	for i, m := range data.Wasm.Contracts {
		info := m.ContractInfo
		var d types.TgradeContractDetails
		if err := info.ReadExtension(&d); err != nil {
			return nil, sdkerrors.Wrapf(err, "extension contract: %d, %s", i, m.ContractAddress)
		}
		if len(d.RegisteredCallbacks) == 0 {
			continue // nothing to do
		}

		addr, err := sdk.AccAddressFromBech32(m.ContractAddress)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "contract: %d", i)
		}
		if err := keeper.importPrivileged(ctx, addr, info.CodeID, d); err != nil {
			return nil, sdkerrors.Wrapf(err, "callback registration for contract: %s", m.ContractAddress)
		}
		importedCallbackContracts[m.ContractAddress] = struct{}{}
	}

	// set privileged
	for i, a := range data.PrivilegedContractAddresses {
		if _, ok := importedCallbackContracts[a]; ok {
			delete(importedCallbackContracts, a)
			continue // done via import already
		}
		// handle new genesis message contract
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "privileged contract: %d", i)
		}
		if err := keeper.SetPrivileged(ctx, addr); err != nil {
			return nil, sdkerrors.Wrapf(err, "set privileged flag for contract %s", a)
		}
	}

	// sanity check that we do not have a callback imported without a privileged flag missing
	if len(importedCallbackContracts) != 0 {
		return nil, sdkerrors.Wrapf(wasmtypes.ErrInvalidGenesis, "unprivileged contracts with system callbacks: %#v", importedCallbackContracts)
	}
	return result, nil
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper *Keeper) *types.GenesisState {
	var genState types.GenesisState
	genState.Wasm = *wasmkeeper.ExportGenesis(ctx, &keeper.Keeper)
	keeper.IteratePrivileged(ctx, func(contract sdk.AccAddress) bool {
		genState.PrivilegedContractAddresses = append(genState.PrivilegedContractAddresses, contract.String())
		return false
	})
	return &genState
}
