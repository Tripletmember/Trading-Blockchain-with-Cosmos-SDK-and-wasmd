package contract_test

import (
	_ "embed"
	"encoding/json"
	"sort"
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctmtypes "github.com/cosmos/cosmos-sdk/x/ibc/light-clients/07-tendermint/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/confio/tgrade/x/poe/contract"
	"github.com/confio/tgrade/x/poe/types"
)

//go:embed tg4_stake.wasm
var randomContract []byte

func TestValidatorsGovProposal(t *testing.T) {
	// setup contracts and seed some data
	ctx, example, vals := setupPoEContracts(t)
	require.Len(t, vals, 3)

	op1Addr, _ := sdk.AccAddressFromBech32(vals[0].OperatorAddress)
	anyAddress := types.RandomAccAddress()

	contractKeeper := example.TWasmKeeper.GetContractKeeper()

	valVotingAddr, err := example.PoEKeeper.GetPoEContractAddress(ctx, types.PoEContractTypeValidatorVoting)
	require.NoError(t, err)
	distrAddr, err := example.PoEKeeper.GetPoEContractAddress(ctx, types.PoEContractTypeDistribution)
	require.NoError(t, err)
	// ensure members set
	members, err := contract.QueryTG4Members(ctx, example.TWasmKeeper, distrAddr, nil)
	require.NoError(t, err)
	require.Len(t, members, 3)
	for _, m := range members {
		t.Logf("%s : %d\n", m.Addr, m.Weight)
	}
	// upload any contract that is not pinned
	codeID, err := contractKeeper.Create(ctx, anyAddress, randomContract, nil)
	require.NoError(t, err)
	require.False(t, example.TWasmKeeper.IsPinnedCode(ctx, codeID), "pinned")
	any, err := codectypes.NewAnyWithValue(&ibctmtypes.ClientState{})
	require.NoError(t, err)
	specs := map[string]struct {
		src       contract.ValidatorProposal
		assertExp func(t *testing.T, ctx sdk.Context)
	}{
		"pin code": {
			src: contract.ValidatorProposal{
				PinCodes: &contract.CodeIDsWrapper{
					CodeIDs: []uint64{codeID},
				},
			},
			assertExp: func(t *testing.T, ctx sdk.Context) {
				assert.True(t, example.TWasmKeeper.IsPinnedCode(ctx, codeID), "pinned")
			},
		},
		"chain upgrade": {
			src: contract.ValidatorProposal{
				RegisterUpgrade: &contract.ChainUpgrade{
					Name:     "v2",
					Info:     "v2-info",
					Height:   7654321,
					DeleteMe: []byte(`{"type_url": "/ibc.lightclients.tendermint.v1.ClientState", "value": "EgAaACIAKgAyADoA"}`),
				},
			},
			assertExp: func(t *testing.T, ctx sdk.Context) {
				gotPlan, exists := example.UpgradeKeeper.GetUpgradePlan(ctx)
				assert.True(t, exists, "exists")
				exp := upgradetypes.Plan{
					Name:                "v2",
					Info:                "v2-info",
					Height:              7654321,
					UpgradedClientState: any,
				}
				assert.Equal(t, exp, gotPlan)
			},
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
			// when submit proposal
			proposalMsg := contract.ValidatorVotingExecuteMsg{
				Propose: &contract.ValidatorVotingPropose{
					Title:       "My proposal",
					Description: "My description",
					Proposal:    spec.src,
				},
			}
			msgBz, err := json.Marshal(proposalMsg)
			require.NoError(t, err)
			_, err = contractKeeper.Execute(ctx, valVotingAddr, op1Addr, msgBz, nil)
			require.NoError(t, err, "exec: %s", string(msgBz))

			// then it is persisted
			adapter := contract.NewVotingContractAdapter(valVotingAddr, example.TWasmKeeper, nil)
			rsp, err := adapter.LatestProposal(ctx)
			require.NoError(t, err)
			require.Equal(t, contract.ProposalStatusOpen, rsp.Status)
			myProposalID := rsp.ID
			t.Logf("%d %s- voting power: %s\n", 0, op1Addr.String(), vals[0].Tokens)

			// and when all validators vote
			// first val has auto YES due to submission, let another one vote
			for i, val := range vals[1:] {
				t.Logf("%d %s - voting power: %s\n", i+1, val.OperatorAddress, val.Tokens)
				opAddr, _ := sdk.AccAddressFromBech32(val.OperatorAddress)
				require.NoError(t, adapter.VoteProposal(ctx, myProposalID, contract.YES_VOTE, opAddr), "voter: %d", i)
			}
			// then
			rsp, err = adapter.QueryProposal(ctx, myProposalID)
			require.NoError(t, err)
			require.Equal(t, contract.ProposalStatusPassed, rsp.Status)

			// and when execute proposal
			require.NoError(t, adapter.ExecuteProposal(ctx, myProposalID, op1Addr))

			// then
			rsp, err = adapter.QueryProposal(ctx, myProposalID)
			require.NoError(t, err)
			require.Equal(t, contract.ProposalStatusExecuted, rsp.Status)
			// and verify action state
			spec.assertExp(t, ctx)
		})
	}
}

func TestQueryTG4Members(t *testing.T) {
	// setup contracts and seed some data
	ctx, example, vals := setupPoEContracts(t)
	require.Len(t, vals, 3)

	contractAddr, err := example.PoEKeeper.GetPoEContractAddress(ctx, types.PoEContractTypeDistribution)
	require.NoError(t, err)
	// ensure members set
	// FIXME: Get members from test setup
	expMembers, err := contract.QueryTG4Members(ctx, example.TWasmKeeper, contractAddr, nil)
	require.NoError(t, err)
	require.Len(t, expMembers, 3)

	specs := map[string]struct {
		pagination *types.Paginator
		expVal     []contract.TG4Member
		expEmpty   bool
		expError   bool
	}{
		"query no pagination": {
			pagination: nil,
			expVal:     expMembers,
		},
		"query offset 0, limit 2": {
			pagination: &types.Paginator{Limit: 2},
			expVal:     expMembers[:2],
		},
		"query offset 2, limit 2": {
			pagination: &types.Paginator{StartAfter: []byte(expMembers[1].Addr), Limit: 2},
			expVal:     expMembers[2:],
		},
		"query offset 3, limit 2": {
			pagination: &types.Paginator{StartAfter: []byte(expMembers[2].Addr), Limit: 2},
			expEmpty:   true,
		},
		"query offset invalid addr, limit 2": {
			pagination: &types.Paginator{StartAfter: []byte("invalid"), Limit: 2},
			expError:   true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			// when
			gotMembers, err := contract.QueryTG4Members(ctx, example.TWasmKeeper, contractAddr, spec.pagination)

			// then
			if spec.expError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if spec.expEmpty {
					assert.Equal(t, 0, len(gotMembers))
				} else {
					assert.Equal(t, spec.expVal, gotMembers)
				}
			}
		})
	}
}

func buildStartAfter(t *testing.T, m *contract.TG4Member) []byte {
	startAfter, err := json.Marshal(m)
	require.NoError(t, err)
	return startAfter
}

func TestQueryTG4MembersByWeight(t *testing.T) {
	// setup contracts and seed some data
	ctx, example, vals := setupPoEContracts(t)
	require.Len(t, vals, 3)

	contractAddr, err := example.PoEKeeper.GetPoEContractAddress(ctx, types.PoEContractTypeDistribution)
	require.NoError(t, err)
	// ensure members set
	// FIXME: Get members from test setup
	expMembers, err := contract.QueryTG4Members(ctx, example.TWasmKeeper, contractAddr, nil)
	require.NoError(t, err)
	require.Len(t, expMembers, 3)

	sort.Slice(expMembers, func(i, j int) bool {
		return expMembers[i].Weight > expMembers[j].Weight
	})

	specs := map[string]struct {
		pagination *types.Paginator
		expVal     []contract.TG4Member
		expEmpty   bool
		expError   bool
	}{
		"query no pagination": {
			pagination: nil,
			expVal:     expMembers,
		},
		"query offset 0, limit 2": {
			pagination: &types.Paginator{Limit: 2},
			expVal:     expMembers[:2],
		},
		"query offset 2, limit 2": {
			pagination: &types.Paginator{StartAfter: buildStartAfter(t, &expMembers[1]), Limit: 2},
			expVal:     expMembers[2:],
		},
		"query offset 3, limit 2": {
			pagination: &types.Paginator{StartAfter: buildStartAfter(t, &expMembers[2]), Limit: 2},
			expEmpty:   true,
		},
		"query offset invalid addr, limit 2": {
			pagination: &types.Paginator{StartAfter: []byte("invalid"), Limit: 2},
			expError:   true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			// when
			gotMembers, err := contract.QueryTG4MembersByWeight(ctx, example.TWasmKeeper, contractAddr, spec.pagination)

			// then
			if spec.expError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if spec.expEmpty {
					assert.Equal(t, 0, len(gotMembers))
				} else {
					assert.Equal(t, spec.expVal, gotMembers)
				}
			}
		})
	}
}
